package opentok

import (
	"net/url"
	"strings"

	"github.com/aljiwala/opentok/config"
)

// Endpoint should construct API endpoint, which would return an *url.URL.
func Endpoint(parts ...string) *url.URL {
	return endpoint(parts...)
}

// CreateSessionEndpoint constuct create session API endpoint, which should
// return an *url.URL.
func CreateSessionEndpoint() *url.URL {
	return createSessionEndpoint()
}

func endpoint(parts ...string) *url.URL {
	up := []string{config.APIHost}
	up = append(up, parts...)
	u, parseErr := url.Parse(strings.Join(up, "/"))
	if parseErr != nil {
		panic(parseErr)
	}
	return u
}

func sessionEndpoint(end string) *url.URL {
	return endpoint("session", end)
}

func createSessionEndpoint() *url.URL {
	return sessionEndpoint("create")
}
