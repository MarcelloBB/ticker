package worker

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/MarcelloBB/ticker/internal/service"
)

type UptimeWorker struct {
	service     *service.UptimeService
	interval    time.Duration
	concurrency int
}

func NewUptimeWorker(service *service.UptimeService, interval time.Duration, concurrency int) *UptimeWorker {
	if interval <= 0 {
		interval = 30 * time.Second
	}
	if concurrency <= 0 {
		concurrency = 5
	}

	return &UptimeWorker{
		service:     service,
		interval:    interval,
		concurrency: concurrency,
	}
}

func (w *UptimeWorker) Start(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(w.interval)
		defer ticker.Stop()

		w.runCycle(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				w.runCycle(ctx)
			}
		}
	}()
}

func (w *UptimeWorker) runCycle(ctx context.Context) {
	targets, err := w.service.ListTargets(ctx)
	if err != nil {
		fmt.Println("background uptime worker list error:", err)
		return
	}

	if len(targets) == 0 {
		return
	}

	sem := make(chan struct{}, w.concurrency)
	var wg sync.WaitGroup

	for _, target := range targets {
		targetID := target.ID
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			if _, probeErr := w.service.CheckTarget(ctx, targetID); probeErr != nil {
				if errors.Is(probeErr, service.ErrTargetNotFound) {
					return
				}
				fmt.Printf("background uptime probe failed for target %d: %v\n", targetID, probeErr)
			}
		}()
	}

	wg.Wait()
}
