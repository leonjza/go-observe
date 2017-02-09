package observatory

type ScanObject struct {
	// If an error occured, the Error key will be set. This is how
	// we will know if shit went south.
	Error string `json:"error"`

	// Values from:
	//	https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#scan
	EndTime             string            `json:"end_time"`
	Grade               string            `json:"grade"`
	Hidden              bool              `json:"hidden"`
	LikelihoodIndicator string            `json:"likelihood_indicator"`
	ResponseHeaders     map[string]string `json:"response_headers"`
	ScanID              int               `json:"scan_id"`
	Score               int               `json:"score"`
	StartTime           string            `json:"start_time"`
	State               string            `json:"state"`
	TestsFailed         int               `json:"tests_failed"`
	TestsPassed         int               `json:"tests_passed"`
	TestsQuantity       int               `json:"tests_quantity"`
}

type ScanDetail struct {
	// Values from:
	//  https://github.com/mozilla/http-observatory/blob/master/httpobs/docs/api.md#tests
	Expectation      string `json:"expectation"`
	Name             string `json:"name"`
	Pass             bool   `json:"pass"`
	Result           string `json:"result"`
	ScoreDescription string `json:"score_description"`
	ScoreModifier    int    `json:"score_modifier"`
}
