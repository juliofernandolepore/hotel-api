package db

import (
	"github.com/juliofernandolepore/hotel-api/types"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const NombreDB = "hotel-reservation"
const coleccionUsuarios = "usuarios"

type IAlmacenadorUsuario interface {
	GetUserByID(string) (*types.Usuario, error)
}

type AlmacenarUsuarioMongo struct {
	//implementacion de la interfaz para mongo
	mongoClient *mongo.Client
	coleccion   *mongo.Collection
}

func NuevoMongoAlmacenador(c *mongo.Client) *AlmacenarUsuarioMongo {
	//contructor de objeto mongo
	//argumento es una conexion, en este caso mongo
	// devuelve un objeto con todos los datos de la conexion
	return &AlmacenarUsuarioMongo{
		mongoClient: c,
		coleccion:   c.Database(NombreDB).Collection(coleccionUsuarios),
	}
}

func (s *AlmacenarUsuarioMongo) GetUserByID(id string) (*types.Usuario, error) {
	return nil, nil
}
