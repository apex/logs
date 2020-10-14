package logs_test

import (
	"os"
	"strings"
	"testing"

	"github.com/apex/logs"
	"github.com/tj/assert"
)

// Test GetProjects.
func TestService_GetProjects(t *testing.T) {
	c := logs.Client{
		URL:       os.Getenv("URL"),
		AuthToken: os.Getenv("AUTH_TOKEN"),
	}

	res, err := c.GetProjects()
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Projects, "empty projects")

	p := res.Projects[0]
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Description)
	assert.NotEmpty(t, p.Name)
}

// Test erroring on HTTP.
func TestHTTP(t *testing.T) {
	c := logs.Client{
		URL:       strings.Replace(os.Getenv("URL"), "https://", "http://", 1),
		AuthToken: os.Getenv("AUTH_TOKEN"),
	}

	_, err := c.GetProjects()
	assert.EqualError(t, err, "Client.URL must be HTTPS, not HTTP")
}
