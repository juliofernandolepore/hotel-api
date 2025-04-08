package types

type Usuario struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre   string `bson:"nombre" json:"nombre"`
	Apellido string `bson:"apellido" json:"apellido"`
}
