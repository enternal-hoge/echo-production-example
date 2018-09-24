package error

import "fmt"

func Handle(err error) {
	if err != nil {
		switch e := err.(type) {
		case *DbConnectError:
			fmt.Println("DbConnectError", "message:", e.Message)
		case *DbExecuteError:
			fmt.Println("DbConnectError", "message:", e.Message)
		case *JsonMarshalError:
			fmt.Println("JsonMarshalError", "message:", e.Message)
		case *JsonUnmarshalError:
			fmt.Println("JsonUnmarshalError", "message:", e.Message)
		case *ReadRequestBodyError:
			fmt.Println("ReadRequestBodyError", "message:", e.Message)
		case *TimeParseError:
			fmt.Println("TimeParseError", "message:", e.Message)
		case *configParseError:
			fmt.Println("configParseError", "message:", e.Message)
		default:
			fmt.Println("Unknown Error", err)
		}
	} else {
		fmt.Println("no error")
	}
}

type DbConnectError struct {
	Message string
}

func (e *DbConnectError) Error() string {
	return e.Message
}

type DbExecuteError struct {
	Message string
}

func (e *DbExecuteError) Error() string {
	return e.Message
}

type JsonMarshalError struct {
	Message string
}

func (e *JsonMarshalError) Error() string {
	return e.Message
}

type JsonUnmarshalError struct {
	Message string
}

func (e *JsonUnmarshalError) Error() string {
	return e.Message
}

type ReadRequestBodyError struct {
	Message string
}

func (e *ReadRequestBodyError) Error() string {
	return e.Message
}

type TimeParseError struct {
	Message string
}

func (e *TimeParseError) Error() string {
	return e.Message
}

type configParseError struct {
	Message string
}

func (e *configParseError) Error() string {
	return e.Message
}
