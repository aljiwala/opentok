package opentok

// Wraps http.Response. So we can add more functionalities later.
import "net/http"

// Response ...
type Response struct {
	*http.Response
	Pagination
}

// NewResponse ...
func NewResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}
