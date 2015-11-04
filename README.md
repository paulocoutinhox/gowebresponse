# gowebresponse

Golang WebRespone class for web and service

With this library you can always send the same json structure to your web response, it turn development very fast and easy because to other part always know what it will receive

# Installing

```bash
go get github.com/prsolucoes/gowebresponse
```

# Importing into your project

```golang
import "github.com/prsolucoes/gowebresponse"
```

# How to use

```golang
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
```

Obs: There is a test method to use all methods

# Test

```bash
make test
```

or

```bash
go test -v
```

Thanks.
