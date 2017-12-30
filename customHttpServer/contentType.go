package customHttpServer

type contentType string

const (
	TextHtml = contentType("text/html")
	Json = contentType("application/json")
	Xml = contentType("application/xml")
	TextPlain = contentType("text/plain")
)
