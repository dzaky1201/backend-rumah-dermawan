package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type WebResponseList struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Count  int         `json:"count"`
	Offset int         `json:"offset"`
	Data   interface{} `json:"data"`
}
