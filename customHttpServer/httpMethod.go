package customHttpServer

type HttpMethod string

const (
	Get    = HttpMethod("GET")
	Post   = HttpMethod("POST")
	Put    = HttpMethod("PUT")
	Delete = HttpMethod("DELETE")
)
