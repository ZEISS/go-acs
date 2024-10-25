package events_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/stretchr/testify/require"
	"github.com/zeiss/go-acs/events"
)

func TestEventHandler_ServeHTTP(t *testing.T) {
	event := cloudevents.NewEvent()
	event.SetID("test")
	event.SetType("test")
	event.SetSource("test")

	bb := new(bytes.Buffer)
	enc := json.NewEncoder(bb).Encode([]cloudevents.Event{event})
	require.NoError(t, enc)

	req, err := http.NewRequest(http.MethodPost, "/events", bb)
	require.NoError(t, err)

	in := make(chan cloudevents.Event, 1)

	hh := events.NewEventHandler(events.WithEvents(in))
	rr := httptest.NewRecorder()

	hh.ServeHTTP(rr, req)
	require.Equal(t, http.StatusAccepted, rr.Code)

	e := <-in
	require.Equal(t, event, e)
}

func TestFilterFunc(t *testing.T) {
	event := cloudevents.NewEvent()
	event.SetID("test")
	event.SetType("test")
	event.SetSource("test")

	fn := events.FilterFunc("test")
	require.True(t, fn(event))

	bb := new(bytes.Buffer)
	enc := json.NewEncoder(bb).Encode([]cloudevents.Event{event})
	require.NoError(t, enc)

	req, err := http.NewRequest(http.MethodPost, "/events", bb)
	require.NoError(t, err)

	in := make(chan cloudevents.Event, 1)

	hh := events.NewEventHandler(events.WithEvents(in))
	rr := httptest.NewRecorder()

	hh.ServeHTTP(rr, req)
	require.Equal(t, http.StatusAccepted, rr.Code)

	e := <-events.Filter(in, "test")
	require.Equal(t, event, e)
}
