package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	casterAddr := flag.String("caster", "127.0.0.1:2101", "caster host:port")
	outPath := flag.String("out", "mountpoints.json", "path to the output stats file")
	pollInterval := flag.Duration("poll", 60*time.Second, "sourcetable poll interval")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	stats := NewStatsFile(*outPath)
	run(ctx, *casterAddr, *pollInterval, stats)
}

func run(ctx context.Context, casterAddr string, pollInterval time.Duration, stats *StatsFile) {
	active := make(map[string]context.CancelFunc)
	defer func() {
		for _, cancel := range active {
			cancel()
		}
	}()

	refresh := func() {
		mounts, err := fetchSourcetable(casterAddr)
		if err != nil {
			log.Printf("sourcetable refresh failed: %v", err)
			return
		}
		seen := make(map[string]bool, len(mounts))
		for _, m := range mounts {
			seen[m] = true
			if _, ok := active[m]; ok {
				continue
			}
			mctx, cancel := context.WithCancel(ctx)
			active[m] = cancel
			log.Printf("subscribing to new mountpoint %q", m)
			go runMountpoint(mctx, casterAddr, m, stats)
		}
		for m, cancel := range active {
			if !seen[m] {
				log.Printf("mountpoint %q no longer in sourcetable, stopping", m)
				cancel()
				delete(active, m)
			}
		}
	}

	refresh()
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			refresh()
		case <-ctx.Done():
			log.Println("shutting down")
			return
		}
	}
}
