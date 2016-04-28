package gowebresponse

import (
	"encoding/json"
	"fmt"
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

	m := make([]struct {
		Key   string `json:"0"`
		Value string `json:"1"`
	}, 0)

	s.Data["errors"] = m
}

func (s *WebResponse) AddDataError(key string, message string) error {
	if s.Data == nil {
		s.Data = make(map[string]interface{})
	}

	errMap, ok := s.Data["errors"]

	if !ok {
		v := [2]string{key, message}

		m := make([][2]string, 0)
		m = append(m, v)
		s.Data["errors"] = m

		return nil
	}

	switch m := errMap.(type) {
	default:
		return fmt.Errorf("Unexpected type: %T", m)
	case [][2]string:
		v := [2]string{key, message}

		m = append(m, v)
		s.Data["errors"] = m
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
