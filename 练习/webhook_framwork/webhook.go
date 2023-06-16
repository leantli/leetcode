// Package tgitwebhook provides a framework for handling a TGit webhook
// requests.
//
// See https://git.woa.com/help/menu/manual/webhooks.html.
//
// NB. The types defined in this package are (most likely) noninterchangeable
// with the ones defined in tgit package thanks to the TGit team.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const maxResponseBodyBytes = 1 << 20 // 1MiB

const (
	headerEventType = "X-Event"
	headerTraceID   = "X-Trace-Id"
)

const (
	eventTypeMergeRequest = "Merge Request Hook"
	eventTypeNote         = "Note Hook"
	eventTypePush         = "Push Hook"
	eventTypeReview       = "Review Hook"
)

// Error denotes an webhook handler error.
type Error struct {
	code int
	msg  string
}

// Error implements error.
func (e *Error) Error() string {
	return e.msg
}

// Errorf returns an webhook error with a given http status code.
func Errorf(statusCode int, format string, a ...interface{}) error {
	return &Error{
		statusCode,
		fmt.Sprintf(format, a...),
	}
}

// Handler defines the interface for handling webhook requests.
// Note that, TGit uses tight deadline for webhook requests, hence don't do anything
// heavy within the webhook context.
type Handler interface {
	HandleMergeRequestEvent(ctx context.Context, evt *MergeRequestEvent) error
	HandlePushEvent(ctx context.Context, evt *PushEvent) error
	HandleNoteEvent(ctx context.Context, evt *NoteEvent) error
	HandleReviewEvent(ctx context.Context, evt *ReviewEvent) error
}

var errUnimplemented = Errorf(http.StatusNotImplemented, "unimplemented")

// UnimplementedHandler is an unimplemented Handler, which can be
// embedded in webhook implementations.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// HandleMergeRequestEvent implements Handler.
func (UnimplementedHandler) HandleMergeRequestEvent(context.Context, *MergeRequestEvent) error {
	return errUnimplemented
}

// HandlePushEvent implements Handler.
func (UnimplementedHandler) HandlePushEvent(context.Context, *PushEvent) error {
	return errUnimplemented
}

// HandleNoteEvent implements Handler.
func (UnimplementedHandler) HandleNoteEvent(context.Context, *NoteEvent) error {
	return errUnimplemented
}

// HandleReviewEvent implements Handler.
func (UnimplementedHandler) HandleReviewEvent(context.Context, *ReviewEvent) error {
	return errUnimplemented
}

// Server returns a http.Handler that serves TGit webhook requests.
func Server(h Handler) http.Handler {
	return server{h}
}

// server handles webhook requests from TGit service.
type server struct {
	Handler
}

// ServeHTTP implements http.Handler
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tid := r.Header.Get(headerTraceID)
	if err := s.serveHTTP(w, r); err != nil {
		code := http.StatusInternalServerError
		var we *Error
		if errors.As(err, &we) {
			code = we.code
		}
		log.Printf("Failed to handle request %s: %v", tid, err)
		http.Error(w, err.Error(), code)
	}
}

func (s server) serveHTTP(w http.ResponseWriter, r *http.Request) error {
	ty := r.Header.Get(headerEventType)
	switch ty {
	default:
		return Errorf(http.StatusNotImplemented, "unsupported event %q", ty)
	case eventTypeMergeRequest:
		var evt MergeRequestEvent
		if err := parseRequest(r, &evt); err != nil {
			return err
		}
		return s.HandleMergeRequestEvent(r.Context(), &evt)
	case eventTypePush:
		var evt PushEvent
		if err := parseRequest(r, &evt); err != nil {
			return err
		}
		return s.HandlePushEvent(r.Context(), &evt)
	case eventTypeNote:
		var evt NoteEvent
		if err := parseRequest(r, &evt); err != nil {
			return err
		}
		return s.HandleNoteEvent(r.Context(), &evt)
	case eventTypeReview:
		var evt ReviewEvent
		if err := parseRequest(r, &evt); err != nil {
			return err
		}
		return s.HandleReviewEvent(r.Context(), &evt)
	}
}

func parseRequest(hreq *http.Request, req interface{}) (err error) {
	b, err := io.ReadAll(io.LimitReader(hreq.Body, maxResponseBodyBytes))
	if err != nil {
		return Errorf(http.StatusBadRequest, "failed to read request: %v", err)
	}
	if err := json.Unmarshal(b, req); err != nil {
		return Errorf(http.StatusBadRequest, "failed to parse request: %v", err)
	}
	return nil
}
