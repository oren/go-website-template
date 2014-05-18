ToDo
========

This application has been written for a talk on Go, TDD, and REST APIs.

The slides for the talk can be seen [here](http://go-talks.appspot.com/github.com/campoy/todo/talk/talk.slide).

TODO provides a simple task management application.

The purpose of this project is to show how some general techniques can be applied to Go code.

##### Run

    cd todo
    go run local.go
    open 0.0.0.0:8080

##### Test Driven Development

The task package has been implemented using TDD techniques, starting writing tests, seeing them failing, writing the code to make the test pass.

Once all tests pass the code is refactored.

##### Implementation of a REST API in Go

The server package handles requests on `"/task/"` providing a REST API.

This REST API is then exposed as a stand-alone http server and a Google App Engine app.
Please take into account that the TaskManager is only in-memory at this point, which means that tasks are
lost if the application is restarted.

##### Consuming a REST API with AngularJS

With a minimum of HTML, CSS, and AngularJS we create a UI for the REST API.
The code is in the `server/static` directory.
