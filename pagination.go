package opentok

// Pagination ...
type Pagination struct {
	Page            int    `json:"page"`
	NumPages        int    `json:"num_pages"`
	PageSize        int    `json:"page_size"`
	Total           int    `json:"total"`
	Start           int    `json:"start"`
	End             int    `json:"end"`
	URI             string `json:"uri"`
	FirstPageURI    string `json:"first_page_uri"`
	PreviousPageURI string `json:"previous_page_uri"`
	NextPageURI     string `json:"next_page_uri"`
	LastPageURI     string `json:"last_page_uri"`
}
