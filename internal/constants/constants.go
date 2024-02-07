package constants

type HTTPVersion uint8

const (
	HTTP2 HTTPVersion = iota
	HTTP11
)