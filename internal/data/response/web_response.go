package response

type WebResponse struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   interface{} `json:"Data,omitempty"`
}
