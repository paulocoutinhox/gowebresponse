package gowebresponse

import (
	"testing"
	"fmt"
)

func TestNewList(t *testing.T) {
	wr := NewGoWebResponse()

	if wr == nil {
		t.Error("WebResponse is not created")
	}
}

func TestToString(t *testing.T) {
	wr := NewGoWebResponse()

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestAdd(t *testing.T) {
	wr := NewGoWebResponse()
	wr.AddData("key", "value")

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{"key":"value"}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestClearData(t *testing.T) {
	wr := NewGoWebResponse()
	wr.AddData("key1", "value1")
	wr.AddData("key2", "value2")
	wr.ClearData()

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestAddDataError(t *testing.T) {
	wr := NewGoWebResponse()
	wr.AddDataError("key", "value")

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{"errors":{"key":"value"}}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestAddDataAndAddDataError(t *testing.T) {
	wr := NewGoWebResponse()
	wr.AddData("key", "value")
	wr.AddDataError("key", "value")

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{"errors":{"key":"value"},"key":"value"}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestClearDataError(t *testing.T) {
	wr := NewGoWebResponse()
	wr.AddDataError("key", "value")
	wr.ClearDataErrors()

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{"errors":{}}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestSuccessTrue(t *testing.T) {
	wr := NewGoWebResponse()
	wr.Success = true

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":true,"message":"","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestSuccessFalse(t *testing.T) {
	wr := NewGoWebResponse()
	wr.Success = false

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestMessageEmpty(t *testing.T) {
	wr := NewGoWebResponse()
	wr.Message = ""

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func TestMessageFilled(t *testing.T) {
	wr := NewGoWebResponse()
	wr.Message = "test"

	jsonData, err := wr.ToString()

	if err != nil {
		t.Errorf("Error: %s", err)
		return
	}

	var jsonResult = jsonData
	var jsonExpected = `{"success":false,"message":"test","data":{}}`

	if jsonExpected != jsonResult {
		t.Errorf("Invalid json result: %s", jsonResult)
	}
}

func ExampleHowToUse() {
	// initializing
	wr := NewGoWebResponse()

	// adding data
	wr.AddData("name", "Paulo")
	wr.AddData("age", "30")

	// clear data
	wr.ClearData()

	// adding data on error and validate message
	wr.Message = "validate"
	wr.AddDataError("name", "Name field cannot be empty")
	wr.AddDataError("email", "Email is invalid")

	// clear error data
	wr.ClearDataErrors()

	// clear message
	wr.Message = ""

	// get json as string
	jsonData, _ := wr.ToString()
	fmt.Print(jsonData)

	// Output: {"success":false,"message":"","data":{"errors":{}}}
}