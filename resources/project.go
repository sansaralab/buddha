package resources

import "github.com/emicklei/go-restful"

type Project struct {
	Id string `json:"id"`
	Name string `json:"name"`
	GitUrl string `json:"git_url"`
}

type ProjectResource struct {
}

func (r *ProjectResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/projects").
		Doc("Manage projects").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{project-id}").To(r.findProject).
		Doc("get a project").
		Operation("findProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		Writes(Project{}))

	container.Add(ws)
}

func (r *ProjectResource) findProject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("project-id")
	prj := Project{Id: id, Name: "test project", GitUrl: "http"}
	response.WriteEntity(prj)
}