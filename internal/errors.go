package internal

import "encoding/json"

type APIError struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APIError) Error() string {
	return e.Msg
}

type APISuccessData struct {
	Status int         `json:"-"`
	Data   interface{} `json:"-"`
}

type APISuccessResponse struct {
	Status int    `json:"-"`
	Msg    string `json:"message"`
}

func (e APISuccessData) Error() string {
	jsonBytes, _ := json.Marshal(e.Data)
	return string(jsonBytes)
}

func (e APISuccessResponse) Error() string {
	return e.Msg
}
