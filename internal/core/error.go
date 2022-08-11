package core

type errorResp struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Timestamp        int64  `json:"timestamp"`
	Duration         int64  `json:"duration"`
	Exception        string `json:"exception"`
}
