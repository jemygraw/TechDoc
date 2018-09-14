package userservice

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id, Name string
}

func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/users").Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(FindUser))
	service.Route(service.POST("").To(UpdateUser))
	service.Route(service.PUT("/{user-id}").To(CreateUser))
	service.Route(service.DELETE("/{user-id}").To(RemoveUser))
	return service
}

func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	// here you would fetch user from some persistence system
	usr := &User{Id: id, Name: "John Doe"}
	response.WriteEntity(usr)
}

func UpdateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	// here you would update the user with some persistence system
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func CreateUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	// here you would create the user with some persistence system
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func RemoveUser(request *restful.Request, response *restful.Response) {
	// here you would delete the user from some persistence system
}
