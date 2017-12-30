package customHttpServer

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type server interface {
	Request(conn net.Conn) (HttpMethod, string, error)
	Respond(net.Conn, HttpStatus, contentType, string)
}

type HttpServer struct{}

func toHttpMethod(method string) HttpMethod {
	switch method {
	case "GET":
		return Get
	case "POST":
		return Post
	case "DELETE":
		return Delete
	case "PUT":
		return Put
	default:
		panic("Unsupported HTTP Method")
	}
}

func (HttpServer) Request(conn net.Conn) (m HttpMethod, uri string, err error) {
	var i int
	var method string

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// Request Line
			fl := strings.Fields(ln)
			method = fl[0]
			uri = fl[1]
		}

		if ln == "" {
			//Headers are done
			break
		}
		i++
	}

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("error: %v", r)
			}

		}
	}()

	return toHttpMethod(method), uri, err
}

func (HttpServer) Respond(conn net.Conn, status HttpStatus, cType contentType, body *string) {
	fmt.Fprintf(conn, "HTTP/1.1 %s\r\n", status)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(*body))
	fmt.Fprintf(conn, "Content-Type: %s\r\n", cType)
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, *body)
}
