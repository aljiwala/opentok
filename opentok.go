package opentok

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// OpenTok holds all the necessary information to create sessions and interact
// with the OpenTok platform.
type OpenTok struct {
	// APIKey that you get after creating a project at the `OpenTok Dashboard`.
	APIKey int

	// APISecret that you get after creating a project with the `OpenTok Dashboard`.
	APISecret string

	// PartnerAuth ...
	PartnerAuth string

	// UserAgent ...
	UserAgent string

	Client httpClient
}

// Claims ...
// Ref: https://tokbox.com/developer/rest/#authentication
type Claims struct {
	IST string `json:"ist,omitempty"`
	jwt.StandardClaims
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// New creates a new OpenTok object.
func (openT *OpenTok) New(apiKey int, apiSecret string, httpClient *http.Client) *OpenTok {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &OpenTok{
		APIKey:      apiKey,
		APISecret:   apiSecret,
		PartnerAuth: fmt.Sprintf("%d:%s", apiKey, apiSecret),
		Client:      httpClient,
	}
}

// AuthenticationKey ...
func (openT *OpenTok) AuthenticationKey() (string, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(2 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		"project",
		jwt.StandardClaims{
			Issuer:    strconv.Itoa(openT.APIKey),
			IssuedAt:  issuedAt.UTC().Unix(),
			ExpiresAt: expiresAt.UTC().Unix(),
		},
	})
	return token.SignedString([]byte(openT.APISecret))
}

// NewRequest should make a request to particular API endpoint with provided method.
func (openT *OpenTok) NewRequest(
	method string, u *url.URL, body io.Reader, withAuth bool,
) (*http.Request, error) {
	req, _ := http.NewRequest(method, u.String(), body)
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Add("User-Agent", openT.UserAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	if withAuth { // Add auth header(s).
		openTAuth, authErr := openT.AuthenticationKey()
		if authErr != nil {
			return nil, authErr
		}

		req.Header.Add("X-OPENTOK-AUTH", openTAuth)
	}

	return req, nil
}
