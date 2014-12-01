# Web Service

Code for a REST service from this nice tutorial: http://thenewstack.io/make-a-restful-json-api-go/

* [Run](#run)
* [API](#api)
* [Code structure](#code-structure)

## Run

  	go get
  	go build
  	./rest
  	open localhost:3000

## API

* GET /
* GET /todos
* GET /todos/1
* POST /todos

### Usage

		curl localhost:3000

Welcome!

  	curl localhost:3000/todos

[{"id":1,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"},{"id":2,"name":"Host meetup","completed":false,"due":"0001-01-01T00:00:00Z"}]

  	curl localhost:3000/todos/1

{"id":1,"name":"Write presentation","completed":false,"due":"0001-01-01T00:00:00Z"}

  	curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:3000/todos

{"id":3,"name":"New Todo","completed":false,"due":"0001-01-01T00:00:00Z"}

## Code Structure

		main -> router -> routes -> handlers -> repo -> todo
			 -> logger                       -> error


main, router, routes and handlers holds the webservice functionality. We create a slice of structs and insert it into our router.
repo and todo represents our Database. for now it's just a slice of structs.
