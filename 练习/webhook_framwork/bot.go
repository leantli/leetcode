package main

import (
	"context"
	"errors"
	"regexp"
	"sync"
)

// workFunc is an async work to complete.
type workFunc func(context.Context)

// bot automates tencent2 daily maintenance.
type bot struct {
	UnimplementedHandler

	ctx     context.Context
	cancel  context.CancelFunc
	eventCh chan workFunc
	wg      sync.WaitGroup

	pollingCh chan workFunc
}

// Start starts workers for bot work.
func (b *bot) Start(ctx context.Context) {
	go b.startEventWorkers(ctx)
}

// TODO: add gracefully shutdown support.
func (b *bot) startEventWorkers(ctx context.Context) {
	b.ctx, b.cancel = context.WithCancel(ctx)
	const workCnt = 100
	b.eventCh = make(chan workFunc, workCnt)
	for i := 0; i < cap(b.eventCh); i++ {
		b.wg.Add(1)
		go func() {
			for w := range b.eventCh {
				w(b.ctx)
			}
			b.wg.Done()
		}()
	}
	b.wg.Wait()
}

func (b *bot) submitEventWorks(ctx context.Context, ws ...workFunc) error {
	return b.submit(ctx, b.eventCh, ws...)
}

// submit async works to the work queue. The ctx not tied to the
// work but to allow early cancellation in case the queue is full.
func (b *bot) submit(ctx context.Context, ch chan workFunc, ws ...workFunc) error {
	for _, w := range ws {
		select {
		case ch <- w:
			continue
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

func (b *bot) HandleMergeRequestEvent(ctx context.Context, evt *MergeRequestEvent) error {
	switch evt.MergeRequest.Action {
	case MergeRequestActionClose, MergeRequestActionMerge:
		return nil
	}
	if evt.MergeRequest.TargetBranch != "b.cfg.Mainline" {
		return nil
	}
	return b.submitEventWorks(ctx,
		func(ctx context.Context) {
			// 省略具体操作
		},
	)
}

func (b *bot) HandleNoteEvent(ctx context.Context, evt *NoteEvent) error {
	if evt.Note.System {
		return nil
	}
	switch {
	case evt.MergeRequest != "":
		return b.handleMergeRequestNoteEvent(ctx, evt)
	default:
		return errors.New("unimplemented")
	}
}

func (b *bot) HandleReviewEvent(ctx context.Context, evt *ReviewEvent) error {
	if evt.Event != "invite" {
		return nil
	}
	return b.submitEventWorks(ctx,
		func(ctx context.Context) {
			// 省略具体操作
		},
	)
}

var commandRE = regexp.MustCompile(`(?m)^/\w+`)

func (b *bot) handleMergeRequestNoteEvent(ctx context.Context, evt *NoteEvent) error {
	if evt.Note.CommitID != "" {
		return nil
	}
	for _, cmd := range commandRE.FindAllString(evt.Note.Note, -1) {
		switch cmd {
		default:
			if err := b.submitEventWorks(ctx, func(ctx context.Context) {
				// 省略具体操作
			}); err != nil {
				return err
			}
		}
	}
	return nil
}
