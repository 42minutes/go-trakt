package trakt

type pageable struct {
	NextPage  *Hyperlink
	LastPage  *Hyperlink
	FirstPage *Hyperlink
	PrevPage  *Hyperlink
}

// Result wraps the original response for the server and the possible error.
type Result struct {
	Response *Response
	Err      error
	pageable
}

func (r *Result) HasError() bool {
	return r.Err != nil
}

func (r *Result) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}

	return ""
}

func newResult(resp *Response, err error) *Result {
	pageable := pageable{}

	return &Result{Response: resp, pageable: pageable, Err: err}
}
