package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/encero/ruian_parser/ent"
	"github.com/encero/ruian_parser/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	SourceDataURL string
}

func main() {
	cfg := Config{
		SourceDataURL: "",
	}

	err := run(cfg)
	if err != nil {
		log.Println(err)
	}
}

func run(cfg Config) error {
	entc, err := connectDB()
	if err != nil {
		return err
	}
	defer entc.Close()

	httpClient := &http.Client{
		Timeout: 5 * time.Minute,
	}

	api := RuianAPI{
		Doer: httpClient,
	}

	list, err := loadLinkList(api)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mapped := buildPipeline(ctx, httpClient, list)
	storage := configureStorage(entc)

	storeMappedItems(ctx, mapped, storage)

	log.Println("process finished")

	return nil
}

func loadLinkList(api RuianAPI) ([]string, error) {
	log.Println("loading full data link list")

	list, err := api.FullDataLinkList(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed getting full data link list: %w", err)
	}

	log.Println("loaded")

	list, err = FilterNewestFromList(list)
	if err != nil {
		return nil, fmt.Errorf("failed filtering newest from list: %w", err)
	}

	return list, nil
}

func connectDB() (*ent.Client, error) {
	entc, err := ent.Open("sqlite3", "file:db.lite?cache=shared&_fk=1&_journal_mode=wal")
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %w", err)
	}
	defer entc.Close()

	// Run the auto migration tool.
	if err := entc.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %w", err)
	}

	return entc, nil
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

func buildPipeline(ctx context.Context, httpClient *http.Client, list []string) <-chan interface{} {
	ctx, cancel := context.WithCancel(ctx)

	downloaded, downloader := NewDownloader(httpClient, list)
	runStep(ctx, cancel, downloader)

	cached, cacher := NewFileCacher(downloaded)
	runStep(ctx, cancel, cacher)

	decompressed, decompressor := NewDecompressor(cached)
	runStep(ctx, cancel, decompressor)

	mapped, mapper := NewMapper(decompressed)
	runStep(ctx, cancel, mapper)

	return mapped
}

func configureStorage(entc *ent.Client) *Storage {
	storage := NewStorage(entc)

	storage.AddHandler(StoreAddressPlace)
	storage.AddHandler(StoreCity)
	storage.AddHandler(StoreStreet)

	return storage
}

func storeMappedItems(ctx context.Context, mapped <-chan interface{}, storage *Storage) {
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
}
