package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/zeiss/pkg/channels"
	"github.com/zeiss/pkg/slices"
	"github.com/zeiss/pkg/utilx"
)

// FilterFunc is the filter for events.
func FilterFunc(types ...string) func(cloudevents.Event) bool {
	return func(e cloudevents.Event) bool {
		return slices.In(e.Type(), types...)
	}
}

// Filter is the filter for events.
func Filter(input <-chan cloudevents.Event, types ...string) <-chan cloudevents.Event {
	return channels.Filter(input, FilterFunc(types...))
}

// EventHandler is the handler for events.
type EventHandler struct {
	events chan cloudevents.Event
}

// Opt is the option for event handler.
type Opt func(*EventHandler)

// Events returns the events channel.
func (h *EventHandler) Events() <-chan cloudevents.Event {
	return h.events
}

// WithBufferSize sets the buffer size for the events channel.
func WithBufferSize(size int) Opt {
	return func(h *EventHandler) {
		h.events = make(chan cloudevents.Event, size)
	}
}

// WithEvents sets the events channel.
func WithEvents(events chan cloudevents.Event) Opt {
	return func(h *EventHandler) {
		h.events = events
	}
}

// Close closes the events channel.
func (h *EventHandler) Close() {
	if h.events != nil {
		close(h.events)
	}
}

// ServeHTTP is the handler for events.
func (h *EventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var events []cloudevents.Event
	err := dec.Decode(&events)
	if utilx.NotEmpty(err) {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			http.Error(w, "Request body contains badly-formed JSON", http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the payload.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			http.Error(w, fmt.Sprintf("Request body contains unknown field %s", fieldName), http.StatusBadRequest)
		case errors.Is(err, io.EOF):
			http.Error(w, "Request body must not be empty", http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			http.Error(w, "Request body must not be larger than 1MB", http.StatusRequestEntityTooLarge)
		default:
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		http.Error(w, "Request body must only contain a single JSON object", http.StatusBadRequest)
		return
	}

	for _, e := range events {
		h.events <- e
	}

	w.WriteHeader(http.StatusAccepted)
}

// NewEventHandler creates a new event handler.
func NewEventHandler(opts ...Opt) *EventHandler {
	e := &EventHandler{}
	e.events = make(chan cloudevents.Event, 1)

	for _, opt := range opts {
		opt(e)
	}

	return e
}
