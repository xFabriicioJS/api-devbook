package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Routes{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.FindAllUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
}
