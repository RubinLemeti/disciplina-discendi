package helper


type Records struct {
	First  int `json:"first"`  // position number (index) of the first record on the current page relative to the total list
	Last   int `json:"last"`   // position number (index) of the last record on the current page.
	Total  int `json:"total"`  // the total number of resources based on the current filters
	Limit  int `json:"limit"`  // the limit being used to get this collection
	Offset int `json:"offset"` // the number of records being skipped to get this collection
}

type Pagination struct {
	First    int     `json:"first"`    // the number of the first page
	Previous int     `json:"previous"` // the number of the previous page
	Current  int     `json:"current"`  // the number of the current page
	Next     int     `json:"next"`     // the number of the next page
	Last     int     `json:"last"`     // the number of the last page
	Records  Records `json:"records"`  // numbers of the resource
}

type Metadata struct {
	Pagination Pagination `json:"pagination"`
}

type ItemResponseModel[T any] struct {
	Data       T      `json:"data"`
	Path       string `json:"path"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
}
type ListResponseModel[T any] struct {
	Data       T        `json:"data"`
	Meta       Metadata `json:"meta"`
	Path       string   `json:"path"`
	Success    bool     `json:"success"` // true or false
	StatusCode int      `json:"statusCode"`
	Timestamp  string   `json:"timestamp"`
}

type SuccessfulResponseModel struct {
	Path       string `json:"path"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	Timestamp  string `json:"timestamp"`
}

type ErrorSubModel struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

type FailureResponseModel struct {
	Error      ErrorSubModel `json:"error"`
	Path       string        `json:"path"`
	Success    bool          `json:"success"`
	StatusCode int           `json:"statusCode"`
	Timestamp  string        `json:"timestamp"`
}
