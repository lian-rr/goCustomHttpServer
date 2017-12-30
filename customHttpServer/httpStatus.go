package customHttpServer

type HttpStatus string

const (
	Sucess = HttpStatus("200 OK")
	BadRequest = HttpStatus("400 Bad Request")
	NotFound = HttpStatus("404 Not Found")
	ServerError = HttpStatus("500 Internal Server Error")
)
