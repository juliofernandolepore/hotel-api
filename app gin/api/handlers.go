package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofernandolepore/hotel-api/db"
)

type UserHandler struct {
	// implemento la interface
	userStore db.IAlmacenadorUsuario
	// cuando cree un objeto UserHandler se van a heredar los metodos
}

func NuevoUserHandler(userStore db.IAlmacenadorUsuario) *UserHandler {
	//metodo contructor
	//argumento pasar una interface
	// new user Handler
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *gin.Context) error {
	id := "id"
	usuario, err := h.userStore.GetUserByID(id)
	if err != nil {
		return err
	}
	c.JSON(200, usuario)
	return nil

}

func (h *UserHandler) HandleGetUsers(c *gin.Context) error {
	c.JSON(200, gin.H{
		"mensaje": "ferv1",
	})
	return nil
}
