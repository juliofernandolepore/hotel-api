package api

import "log"

type InterfaceNumberStore interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresDbConexion struct {
	//datos de conexion a postgres
}

func (p PostgresDbConexion) GetAll() ([]int, error) {
	log.Println("sin errores en postgres")
	return []int{1, 2, 3}, nil
}

func (p PostgresDbConexion) Put(numero int) error {
	return nil
}

type MongoDbConexion struct {
	//datos de coneccion
	//instanciar o preparar los metodos
}

func (m MongoDbConexion) GetAll() ([]int, error) {
	log.Println("sin errores en mongo")
	return []int{1, 2, 3}, nil
}

func (m MongoDbConexion) Put(numero int) error {
	return nil
}

type ApiServer struct {
	//instanciar la interfaz
	Almacenar InterfaceNumberStore
}
