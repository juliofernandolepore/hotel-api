package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juliofernandolepore/hotel-api/api"
	"github.com/juliofernandolepore/hotel-api/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// constantes para conexion mongo
const dburi = "mongodb://root:mongo@localhost:27017/admin"
const basededatos = "hotel-reservation"
const coleccionUsuarios = "usuarios"

func main() {
	// configuracion de conexxion mongo
	client, err := mongo.Connect(options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatalln("no se pudo efectuar conexion a db mongo", err)
	}
	ctx := context.Background()
	coll := client.Database(basededatos).Collection(coleccionUsuarios)
	usuario := types.Usuario{
		Nombre:   "fer",
		Apellido: "mi apellido",
	}
	res, err := coll.InsertOne(ctx, usuario)
	if err != nil {
		log.Fatalln("no pudo ingresar el usuario a la coleccion", res)
	}
	log.Println(res)
	app := gin.Default()

	//apiv1 := app.Group("/api/v1")

	app.GET("/fer", api.HandleGetUsers)
	app.Run()

}
