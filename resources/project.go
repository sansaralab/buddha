package resources

import "github.com/emicklei/go-restful"

type Project struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
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

	ws.Route(ws.GET("/").To(r.getProjectList).
		Doc("get all projects").
		Operation("getProjectList").
		Writes([]Project{}))

	ws.Route(ws.POST("/").To(r.createProject).
		Doc("create new project").
		Operation("createProject").
		Param(ws.FormParameter("name", "unique name of a new project").DataType("string")).
		Param(ws.FormParameter("git_url", "git url of a new project").DataType("string")).
		Writes(Project{}))

	ws.Route(ws.GET("/{project-id}").To(r.findProject).
		Doc("get a project").
		Operation("findProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		Writes(Project{}))

	ws.Route(ws.PUT("/{project-id}").To(r.updateProject).
		Doc("update a project").
		Operation("updateProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		Param(ws.FormParameter("name", "unique name of a project").DataType("string").Required(false)).
		Param(ws.FormParameter("git_url", "git url of a project").DataType("string").Required(false)).
		Writes(Project{}))

	ws.Route(ws.DELETE("/{project-id}").To(r.deleteProject).
		Doc("delete a project").
		Operation("deleteProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("string")).
		Writes(Project{}))

	container.Add(ws)
}

func (r *ProjectResource) getProjectList(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) createProject(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) updateProject(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) deleteProject(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) findProject(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("project-id")
	prj := Project{Id: id, Name: "test project", GitUrl: "http"}
	response.WriteEntity(prj)
}
