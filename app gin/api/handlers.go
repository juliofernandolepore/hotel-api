package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juliofernandolepore/hotel-api/types"
)

func HandleGetUsers(c *gin.Context) {
	u := types.Usuario{
		Nombre:   "fernando",
		Apellido: "lepore",
	}
	c.JSON(200, gin.H{
		"mensaje": u,
	})
}

func Ferv1(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "ferv1",
	})
}
