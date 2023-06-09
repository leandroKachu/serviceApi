package rotas

import (
	"api/src/controllers"
	"net/http"
)

var routesPost = []Rota{
	{
		URI:                "/posts/create",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreatePost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/posts/{postID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPostbyID,
		RequerAutenticacao: true,
	},
	{
		URI:                "/posts",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPosts,
		RequerAutenticacao: true,
	},
	{
		URI:                "/posts/update/{postID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdatePost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/posts/delete/{postID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletePost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/posts/user/{userID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPostByIDuser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/post/{postID}/like",
		Metodo:             http.MethodPost,
		Funcao:             controllers.LikePost,
		RequerAutenticacao: true,
	},
	{
		URI:                "/post/{postID}/unlike",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Unlike,
		RequerAutenticacao: true,
	},
}
