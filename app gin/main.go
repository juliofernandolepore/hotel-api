package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juliofernandolepore/hotel-api/api"
	"github.com/juliofernandolepore/hotel-api/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// constantes para conexion mongo
const dburi = "mongodb://root:mongo@localhost:27017/admin"

func main() {
	// configuracion de conexxion mongo
	client, err := mongo.Connect(options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatalln("no se pudo efectuar conexion a db mongo", err)
	}
	log.Println("conexion exitosa")

	//instanciar un nuevo objeto handleUsers (del paquete api creado) para utilizar sis metodos
	UserHandler := api.NuevoUserHandler(db.NuevoMongoAlmacenador(client))

	app := gin.Default()

	app.GET("/user", api.UserHandler.HandleGetUser)
	app.Run()

}
