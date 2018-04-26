package config

const (
	urlScheme = "https"

	// BaseAPI ...
	BaseAPI = "api.opentok.com"

	// APIHost ...
	APIHost = urlScheme + "://" + BaseAPI

	// APIVersion represents version of API which we're going to consume.
	APIVersion = "v2"

	// APIFormat represents the format in which the response would be returned.
	APIFormat = "json"
)
