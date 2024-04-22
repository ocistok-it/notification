package custerr

func ErrServiceUnsupported() *CustErr {
	return New("SERVICE_UNSUPPORTED", "service unsupported. currently service supported dingtalk and mail only")
}

func ErrInvalidConnection() *CustErr {
	return New("INVALID_CONN", "invalid connection")
}
