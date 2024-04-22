package custerr

import "fmt"

type (
	CustErr struct {
		Message    string
		Code       string
		Stacktrace error
	}

	ErrCode string
)

func New(code string, message string) *CustErr {
	return &CustErr{Code: code, Message: message}
}

func (c *CustErr) WithStacktrace(err error) *CustErr {
	c.Stacktrace = err
	return c
}

func (c *CustErr) Error() string {
	return fmt.Sprintf("[%s] %s", c.Code, c.Message)
}
