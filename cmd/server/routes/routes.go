package routes

import (
	"PseudoTwitter/cmd/server/handler"
	"PseudoTwitter/internal/user"
	"PseudoTwitter/internal/publicaciones"
	"PseudoTwitter/internal/comentarios"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router interface{
	MapRoutes()

}

type router struct{
	r  *gin.Engine

	db *sql.DB
}
func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	

	r.buildUsersRoutes()
	r.buildPublicationRoutes()
	r.buildCommentsRoutes()
	
}



func (r *router) buildUsersRoutes() {
	// Example
	repo := user.NewRepository(r.db)
	service := user.NewService(repo)
	handler := handler.NewUser(service)
	r.r.GET("/users", handler.GetAll())
	r.r.GET("/users/:id", handler.Get())
	r.r.POST("/users", handler.Create())
	r.r.DELETE("/users/:id", handler.Delete())
	r.r.PATCH("/users/:id", handler.Update())
}
func (r *router) buildPublicationRoutes() {
	// Example
	repo := publicaciones.NewRepository(r.db)
	service := publicaciones.NewService(repo)
	handler := handler.NewPublicacion(service)
	r.r.GET("/publications/:id", handler.Get())
	r.r.POST("/publications", handler.CreatePublication())
	// r.r.DELETE("/publications/:id", handler.Delete())
	// r.r.PATCH("/publications/:id", handler.Update())
}

func (r *router) buildCommentsRoutes() {
	// Example
	repo := comentarios.NewRepository(r.db)
	service := comentarios.NewService(repo)
	handler := handler.NewComment(service)
	r.r.GET("/comments/users/:id", handler.GetAllByUsers())
	r.r.GET("/comments/publications/:id", handler.GetAllByPublications())
	r.r.POST("/comments", handler.CreateComment())
	// r.r.DELETE("/publications/:id", handler.Delete())
	// r.r.PATCH("/publications/:id", handler.Update())
}