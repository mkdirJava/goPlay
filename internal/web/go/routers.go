/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * Documentation of a simple state machine that handles applications work state
 *
 * API version: 0.0.1
 * Contact: projects@mkdirdev.co.uk
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddCompany",
		strings.ToUpper("Post"),
		"/company",
		AddCompany,
	},

	Route{
		"DeleteCompany",
		strings.ToUpper("Delete"),
		"/company/{companyId}",
		DeleteCompany,
	},

	Route{
		"GetCompanyById",
		strings.ToUpper("Get"),
		"/company/{companyId}",
		GetCompanyById,
	},

	Route{
		"DeleteSubtask",
		strings.ToUpper("Delete"),
		"/task/subtask/{subtaskId}",
		DeleteSubtask,
	},

	Route{
		"GetSubtaskById",
		strings.ToUpper("Get"),
		"/task/subtask/{subtaskId}",
		GetSubtaskById,
	},

	Route{
		"UpdateSubtaskById",
		strings.ToUpper("Put"),
		"/task/subtask/{subtaskId}",
		UpdateSubtaskById,
	},

	Route{
		"AddTask",
		strings.ToUpper("Post"),
		"/task",
		AddTask,
	},

	Route{
		"DeleteTask",
		strings.ToUpper("Delete"),
		"/task/{taskId}",
		DeleteTask,
	},

	Route{
		"GetTaskById",
		strings.ToUpper("Get"),
		"/task/{taskId}",
		GetTaskById,
	},

	Route{
		"AddUser",
		strings.ToUpper("Post"),
		"/user",
		AddUser,
	},

	Route{
		"DeleteUser",
		strings.ToUpper("Delete"),
		"/user/{userId}",
		DeleteUser,
	},

	Route{
		"GetUserById",
		strings.ToUpper("Get"),
		"/user/{userId}",
		GetUserById,
	},

	Route{
		"AddWorkflow",
		strings.ToUpper("Post"),
		"/workflow",
		AddWorkflow,
	},

	Route{
		"DeleteWorkflow",
		strings.ToUpper("Delete"),
		"/workflow/{workflowId}",
		DeleteWorkflow,
	},

	Route{
		"GetWorkflowById",
		strings.ToUpper("Get"),
		"/workflow/{workflowId}",
		GetWorkflowById,
	},

	Route{
		"UpdateWorkflowById",
		strings.ToUpper("Put"),
		"/workflow/{workflowId}",
		UpdateWorkflowById,
	},

	Route{
		"GetWorkflowInstanceById",
		strings.ToUpper("Get"),
		"/workflow/running/{workflowId}",
		GetWorkflowInstanceById,
	},

	Route{
		"StartWorkflow",
		strings.ToUpper("Post"),
		"/workflow/running",
		StartWorkflow,
	},

	Route{
		"UpdateWorkflowInstanceById",
		strings.ToUpper("Put"),
		"/workflow/running/{workflowId}",
		UpdateWorkflowInstanceById,
	},
}
