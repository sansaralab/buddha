package main

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/sansaralab/buddha/resources"
	"log"
	"net/http"
	"os/user"
)

func main() {
	//restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))

	wsContainer := restful.NewContainer()
	p := resources.ProjectResource{}
	p.Register(wsContainer)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	config := swagger.Config{
		WebServices:     wsContainer.RegisteredWebServices(),
		WebServicesUrl:  "http://localhost:8080",
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: usr.HomeDir + "/github/swagger-api/swagger-ui/dist",
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	log.Println("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
