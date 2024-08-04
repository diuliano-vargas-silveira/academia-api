package response

import "time"

type Response struct {
	Code        int         `json:"code"`
	Status      string      `json:"status"`
	Data        interface{} `json:"data"`
	Error       interface{} `json:"error"`
	RequestedAt time.Time   `json:"requestedAt"`
}
