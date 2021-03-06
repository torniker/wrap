package response

import (
	"encoding/json"
	"os"

	"github.com/torniker/wrap/logger"
)

// NewCLI create an instance of cli responder
func NewCLI() *CLI {
	return &CLI{
		headers: make(map[string]string),
	}
}

// CLI is type for cli responses
type CLI struct {
	// writer    http.ResponseWriter
	headers   map[string]string
	status    int
	committed bool
	output    interface{}
}

// SetStatus sets response status code
func (c *CLI) SetStatus(status int) {
	c.status = status
}

// SetHeader sets header for responder
func (c *CLI) SetHeader(key, val string) {
	c.headers[key] = val
}

// Status returns response status code
func (c *CLI) Status() int {
	return c.status
}

// Commited returns response status
func (c *CLI) Commited() bool {
	return c.committed
}

// Write commits and writes data into the response body
func (c *CLI) Write(body interface{}) error {
	if c.status == 0 {
		c.SetStatus(200)
	}
	c.committed = true
	c.output = body
	logger.Infof("<--- status: %v, body: %#v", c.status, body)
	return json.NewEncoder(os.Stdout).Encode(body)
}

// SuccessWithNoContent commits and sets http status to 204 No Content
func (c *CLI) SuccessWithNoContent() {
	c.SetStatus(204)
	c.committed = true
}

// Output returns response output
func (c *CLI) Output() interface{} {
	return c.output
}

// EnableCORS sets corresponsing headers to enable CORS
func (c *CLI) EnableCORS(origin, methods, headers string) {
	c.SetHeader("Access-Control-Allow-Origin", origin)
	c.SetHeader("Access-Control-Allow-Methods", methods)
	c.SetHeader("Access-Control-Allow-Headers", headers)
}
