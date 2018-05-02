package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("cant open file: %v", err)
		return
	}
	defer file.Close()

	type entry struct {
		raw  string
		id   string
		date time.Time
	}

	files := make(map[string]entry)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		u, err := url.Parse(line)
		if err != nil {
			log.Printf("cant parse url: %v", err)
			return
		}

		base := path.Base(u.Path)

		parts := strings.Split(base, "_")
		if len(parts) != 4 {
			log.Printf("invalid line ( unexpected number of parts): %v", line)
			return
		}

		t, err := time.Parse("20060102", parts[0])
		if err != nil {
			log.Printf("cant parse date: %v", err)
			return
		}

		e := entry{
			raw:  line,
			id:   parts[2],
			date: t,
		}

		if cur, ok := files[e.id]; ok {
			if cur.date.After(e.date) {
				continue
			}
		}

		files[e.id] = e
	}

	for _, e := range files {
		fmt.Println(e.raw)
	}
}
