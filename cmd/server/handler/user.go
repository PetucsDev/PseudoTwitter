package handler

import (
	
	"PseudoTwitter/internal/domain"
	"PseudoTwitter/internal/user"
	"PseudoTwitter/pkg/web"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"net/http"
)


type User struct {
	

	service user.Service
}

func NewUser(p user.Service) *User {
	return &User{
		
		service: p,
	}
}


func (s *User) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		


		id, err := strconv.ParseInt(c.Param("id"),10, 64)
		if err != nil  {
			c.JSON(400, gin.H{ "error":  "invalid ID"})
			return
		}

		p,err := s.service.Get(c.Request.Context(), int(id))

		if err != nil{
			c.JSON(404, web.NewResponse(404, nil, "No se encuentra el user con el id ingresado"))
			return
		}

		c.JSON(200, p)


	}
}

func (s *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
	

		p, err := s.service.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if len(p) == 0 {
			c.JSON(404, web.NewResponse(404, nil, "No hay users"))
			return
		}
		c.JSON(200, web.NewResponse(200, p, ""))

	}
}

func (s *User) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Users

		if err := c.Bind(&req); err != nil {
			c.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		if s.service.Exists(c.Request.Context(), req.Mail) {
			web.Error(c, 409, "%s", "Ya existe un seller con ese Mail")
			return
		}

		if req.Mail == "" {
			c.JSON(422, web.NewResponse(422, nil, "El mail es requerido"))
			return
		}

		if req.UserName == "" {
			c.JSON(422, web.NewResponse(422, nil, "El nombre de usuario es requerido"))
			return
		}

		if req.Password == "" {
			c.JSON(422, web.NewResponse(422, nil, "La password es requerida"))
			return
		}

		p, err := s.service.Save(c.Request.Context(), req)
		if err != nil {
			c.JSON(422, web.NewResponse(422, nil, err.Error()))
			return
		}
		c.JSON(201, web.NewResponse(201, p, ""))
	}

}

func (s *User) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", "El id no es valido ")
			return
		}

		req := domain.Users{}

		if err := c.ShouldBindJSON(&req); err != nil{
			web.Error(c, http.StatusBadRequest, "Error en la peticion: %s", err)
			return
		}
		lastS, err := s.service.Get(c.Request.Context(), int(id))
		if err != nil{
			web.Error(c, http.StatusNotFound, "No se encuentra el user con el id ingresado")
			return
		}

		newS := updateUserFields(lastS, req, int(id))
		err = s.service.Update(c.Request.Context(), newS)
		if err != nil{
			web.Error(c, http.StatusInternalServerError, "Ocurrio un error al actualizar el user")
			return
		}

		web.Success(c, http.StatusOK, newS)
	
	}
	
}


func updateUserFields(lastSeller domain.Users, newSeller domain.Users, id int) domain.Users {
	if newSeller.UserName != lastSeller.UserName && newSeller.UserName != "" {
		lastSeller.UserName = newSeller.UserName
	}
	if newSeller.Password != lastSeller.Password && newSeller.Password != "" {
		lastSeller.Password = newSeller.Password
	}
	if newSeller.Mail != lastSeller.Mail && newSeller.Mail != "" {
		lastSeller.Mail = newSeller.Mail
	}
	
	lastSeller.ID = id
	return lastSeller
}

func (s *User) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		id, err := strconv.ParseInt(c.Param("id"),10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", "El id ingresado no es valido")
			return
		}
		user, err := s.service.Get(c.Request.Context(), int(id))
		if user.ID == 0 {
			web.Error(c, http.StatusNotFound, "No se encuentra el user con id: %d", id)
			return
		}


		err = s.service.Delete(c.Request.Context(),int(id))
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Ocurrio un error al eliminar el warehouse")
			return
		}
		c.JSON(200, gin.H{ "data": fmt.Sprintf("El user %d ha sido eliminado", id) })


	}
}
