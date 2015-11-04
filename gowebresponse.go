package gowebresponse

import (
	"fmt"
	"encoding/json"
)

type WebResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewGoWebResponse() *WebResponse {
	return &WebResponse{
		Success: false,
		Message: "",
		Data:    make(map[string]interface{}),
	}
}

func (s *WebResponse) AddData(key string, value interface{}) {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}

	s.Data[key] = value
}

func (s *WebResponse) ClearData() {
	s.Data = make(map[string]interface{})
}

func (s *WebResponse) ClearDataErrors() {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}

	v := make(map[string]string)
	s.Data["errors"] = v
}

func (s *WebResponse) AddDataError(key string, message string) error {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}

	errMap, ok := s.Data["errors"]

	if !ok {
		v := make(map[string]string)
		v[key] = message
		s.Data["errors"] = v
		return nil
	}

	switch v := errMap.(type) {
	default:
		return fmt.Errorf("Unexpected type: %T", v)
	case map[string]string:
		v[key] = message
	}

	return nil
}

func (s *WebResponse) ToString() (string, error) {
	jsonData, err := json.Marshal(s)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}
