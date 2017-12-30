package main

import (
	"net"
	"log"
	"fmt"
	cHttpServer "customHttpServer"
	templateLoader "customTemplateLoader"
)

const (
	internalServerErrorPath = "templates/serverError"
	notFoundErrorPath = "templates/notFound"
	badRequestErrorPath = "templates/badRequest"
)

var resources = make(map[cHttpServer.HttpMethod]map[string]string)

func init(){
	//Adding URI's of the GET method
	resources[cHttpServer.Get] = make(map[string]string)
	resources[cHttpServer.Get]["/"] = "templates/helloWorld"
	resources[cHttpServer.Get]["/helloworld"] = "templates/helloWorld"
	resources[cHttpServer.Get]["/byeworld"] = "templates/helloWorld"
}

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	fmt.Println("*********************")
	fmt.Println("******Listening******")
	fmt.Println("*********************")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go request(conn, cHttpServer.HttpServer{})
	}
}


func request(conn net.Conn, server cHttpServer.HttpServer){
	//Request
	defer conn.Close()

	method,uri,err := server.Request(conn)

	var body *string
	if err != nil {
		log.Println(err)
		body,_ = manageErrorTemplate(cHttpServer.ServerError)

		server.Respond(conn, cHttpServer.ServerError, cHttpServer.TextHtml, body)
	} else {

		fmt.Printf("HTTP Method: %s.\n", method)
		fmt.Printf("URi: %s.\n", uri)
		fmt.Printf("============================\n\n\n")

		switch method {
		//Manage GET Request
		case cHttpServer.Get:
			body,err = manageGetRequest(uri)
		default:
			//Bad Request
			body,_ = manageErrorTemplate(cHttpServer.BadRequest)
			server.Respond(conn, cHttpServer.BadRequest, cHttpServer.TextHtml, body)
		}
		if err == nil {
			//Success
			server.Respond(conn, cHttpServer.Sucess, cHttpServer.TextHtml, body)
		} else {
			//NotFound
			body,_ = manageErrorTemplate(cHttpServer.NotFound)
			server.Respond(conn, cHttpServer.NotFound, cHttpServer.TextHtml, body)
		}
	}
}


func manageGetRequest(uri string) (body *string, err error){
	var tp *templateLoader.Template

	p, ok := resources[cHttpServer.Get][uri]

	if ok {
		tp, err = getTemplate(p)
	} else {
		tp, err = getTemplate("templates/notFound")
	}
	return &tp.Content, err
}

func manageErrorTemplate(st cHttpServer.HttpStatus) (body *string, err error){
	var tp *templateLoader.Template

	switch st {
	case cHttpServer.ServerError:
		tp, err = getTemplate(internalServerErrorPath)
	case cHttpServer.NotFound:
		tp, err = getTemplate(notFoundErrorPath)
	case cHttpServer.BadRequest:
		tp, err = getTemplate(badRequestErrorPath)
	}

	return &tp.Content, err
}

func getTemplate(path string) (t *templateLoader.Template, err error){
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)

			if !ok {
				err = fmt.Errorf("error: %v", r)
			}
		}

	}()
	return templateLoader.Load(path), err
}

