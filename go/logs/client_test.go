package logs_test

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apex/logs/go/logs"
	"github.com/tj/assert"
)

// Test GetProjects.
func TestService_GetProjects(t *testing.T) {
	c := logs.Client{
		URL: "http://localhost:5001",
	}

	res, err := c.GetProjects()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Projects, "empty projects")

	p := res.Projects[0]
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Description)
	assert.NotEmpty(t, p.Name)
}

// Benchmark the client.
func Benchmark(b *testing.B) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(ioutil.Discard, r.Body)
	}))
	defer s.Close()

	c := logs.Client{URL: s.URL}
	count := 100

	var events []logs.Event
	for i := 0; i < count; i++ {
		events = append(events, logs.Event{
			Level:   "info",
			Message: "Some message here",
			Fields: map[string]interface{}{
				"region":  "us-west-2",
				"project": "api",
				"version": "v1.0.0",
			},
		})
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.SetBytes(int64(count))

	for i := 0; i < b.N; i++ {
		err := c.AddEvents(logs.AddEventsInput{
			ProjectID: "testing",
			Events:    events,
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}
