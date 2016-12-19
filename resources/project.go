package resources

import (
	"github.com/emicklei/go-restful"
	"github.com/sansaralab/buddha/models"
	"github.com/sansaralab/buddha/repository"
	"strconv"
	"net/http"
)

type ProjectResource struct {
	Repo *repository.Repository
}

func NewProjectResource(repo *repository.Repository) *ProjectResource {
	return &ProjectResource{Repo: repo}
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
		Writes([]models.Project{}))

	ws.Route(ws.POST("/").To(r.createProject).
		Doc("create new project").
		Operation("createProject").
		//Param(ws.FormParameter("name", "unique name of a new project").DataType("string")).
		//Param(ws.FormParameter("git_url", "git url of a new project").DataType("string")).
		Writes(models.Project{}))

	ws.Route(ws.GET("/{project-id}").To(r.findProject).
		Doc("get a project").
		Operation("findProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("integer")).
		Writes(models.Project{}))

	ws.Route(ws.PUT("/{project-id}").To(r.updateProject).
		Doc("update a project").
		Operation("updateProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Param(ws.FormParameter("name", "unique name of a project").DataType("string").Required(false)).
		Param(ws.FormParameter("git_url", "git url of a project").DataType("string").Required(false)).
		Writes(models.Project{}))

	ws.Route(ws.DELETE("/{project-id}").To(r.deleteProject).
		Doc("delete a project").
		Operation("deleteProject").
		Param(ws.PathParameter("project-id", "identifier of the project").DataType("int")).
		Writes(models.Project{}))

	container.Add(ws)
}

func (r *ProjectResource) getProjectList(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) createProject(request *restful.Request, response *restful.Response) {
	project := models.NewProject()
	err := request.ReadEntity(project)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	createdProject := r.Repo.AddProject(project.Name, project.GitUrl)
	response.WriteEntity(createdProject)
}

func (r *ProjectResource) updateProject(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) deleteProject(request *restful.Request, response *restful.Response) {

}

func (r *ProjectResource) findProject(request *restful.Request, response *restful.Response) {
	idString := request.PathParameter("project-id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic("can't convert string to int")
	}
	prj := models.Project{Id: id, Name: "test project", GitUrl: "http"}
	response.WriteEntity(prj)
}
