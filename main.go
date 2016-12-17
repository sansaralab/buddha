package main

import (
)
import (
	"github.com/emicklei/go-restful"
	"github.com/sansaralab/buddha/resources"
	"log"
	"net/http"
	"github.com/emicklei/go-restful/swagger"
)

func main() {
	// to see what happens in the package, uncomment the following
	//restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))

	wsContainer := restful.NewContainer()
	p := resources.ProjectResource{}
	p.Register(wsContainer)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost:8080",
		ApiPath:        "/apidocs.json",

		//// Optionally, specifiy where the UI is located
		//SwaggerPath:     "/apidocs/",
		//SwaggerFilePath: "/Users/emicklei/xProjects/swagger-ui/dist"
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Println("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
