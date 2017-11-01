package exchange

import "fmt"

type ReadBodyError struct{}

func (e *ReadBodyError) Error() string {
	return fmt.Sprint("Cannot read response body")
}

type UnmarshalError struct{}

func (e *UnmarshalError) Error() string {
	return fmt.Sprint("Cannot unmarshal data")
}

type ParseTimeError struct{}

func (e *ParseTimeError) Error() string {
	return fmt.Sprint("Cannot parse time format")
}

type CreateReqError struct{}

func (e *CreateReqError) Error() string {
	return fmt.Sprint("Cannnot create new request")
}

type GetRespError struct{}

func (e *GetRespError) Error() string {
	return fmt.Sprint("Cannnot get response")
}

type ReadBytesError struct{}

func (e *ReadBytesError) Error() string {
	return fmt.Sprint("Cannot read bytes from reader")
}
