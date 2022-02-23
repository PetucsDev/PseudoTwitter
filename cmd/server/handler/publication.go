package handler

import (
	
	"PseudoTwitter/internal/domain"
	"PseudoTwitter/internal/publicaciones"
	"PseudoTwitter/pkg/web"

	//"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	//"net/http"
)


type Publicaciones struct {
	

	service publicaciones.Service
}

func NewPublicacion(p publicaciones.Service) *Publicaciones {
	return &Publicaciones{
		
		service: p,
	}
}


func (s *Publicaciones) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		


		id, err := strconv.ParseInt(c.Param("id"),10, 64)
		if err != nil  {
			c.JSON(400, gin.H{ "error":  "invalid ID"})
			return
		}

		p,err := s.service.GetAll(c.Request.Context(), int(id))

		if err != nil{
			c.JSON(404, web.NewResponse(404, nil, "No se encuentra la publication con el id ingresado"))
			return
		}

		c.JSON(200, p)


	}
}

func (s *Publicaciones) CreatePublication() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req1 domain.Publications

		if err := c.Bind(&req1); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

	

		if req1.Titulo == "" {
			c.JSON(422, web.NewResponse(422, nil, "El mail es requerido"))
			return
		}

		if req1.Fecha == "" {
			c.JSON(422, web.NewResponse(422, nil, "El nombre de usuario es requerido"))
			return
		}

		if req1.UsuariosId == 0 {
			c.JSON(422, web.NewResponse(422, nil, "La password es requerida"))
			return
		}

		p, err := s.service.Save(c.Request.Context(), req1)
		if err != nil {
			c.JSON(422, web.NewResponse(422, nil, err.Error()))
			return
		}
		c.JSON(201, web.NewResponse(201, p, ""))
	}

}