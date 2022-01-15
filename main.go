package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/encero/ruian_parser/ent"
	"github.com/encero/ruian_parser/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	entc, err := ent.Open("sqlite3", "file:db.lite?cache=shared&_fk=1&_journal_mode=wal")
	if err != nil {
		log.Printf("failed opening connection to sqlite: %v", err)
		return
	}
	defer entc.Close()

	// Run the auto migration tool.
	if err := entc.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Printf("failed creating schema resources: %v", err)
		return
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Minute,
	}

	api := RuianAPI{
		Doer: httpClient,
	}

	log.Println("loading full data link list")

	list, err := api.FullDataLinkList(context.Background())
	if err != nil {
		log.Printf("failed getting full data link list: %v", err)
		return
	}

	log.Println("loaded")

	list, err = FilterNewestFromList(list)
	if err != nil {
		log.Printf("failed filtering newest from list: %v", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	downloaded, downloader := NewDownloader(httpClient, list)
	runStep(ctx, cancel, downloader)

	cached, cacher := NewFileCacher(downloaded)
	runStep(ctx, cancel, cacher)

	decompressed, decompressor := NewDecompressor(cached)
	runStep(ctx, cancel, decompressor)

	mapped, mapper := NewMapper(decompressed)
	runStep(ctx, cancel, mapper)

	storage := NewStorage(entc)

	storage.AddHandler(StoreAddressPlace)
	storage.AddHandler(StoreCity)
	storage.AddHandler(StoreStreet)

	loaded := 0
	lastCheckpoint := 0
	start := time.Now()

	for item := range mapped {
		err := storage.Store(ctx, item)
		if err != nil {
			log.Printf("failed storing item: %v", err)
			return
		}

		loaded++
		lastCheckpoint++

		if since := time.Since(start).Seconds(); since > 5 {
			log.Printf("processed: %d %f,0/s", loaded, float64(lastCheckpoint)/since)

			start = time.Now()
			lastCheckpoint = 0
		}
	}

	log.Println("process finished")
}

func runStep(ctx context.Context, cancel context.CancelFunc, step func(context.Context) error) {
	go func() {
		err := step(ctx)
		if err != nil {
			log.Printf("failed running step: %v", err)
			cancel()
		}
	}()
}
