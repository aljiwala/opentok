package tokbox

// Tokbox holds the APIKey and APISecret value.
type Tokbox struct {
	APIKey, APISecret string
}

// NewClient should return a new tokbox client by considering given API key and
// API secret.
func NewClient(apiKey, apiSecret string) *Tokbox {
	return &Tokbox{APIKey: apiKey, APISecret: apiSecret}
}
