package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/juanhayek/goarg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("goarg")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	ObjID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": ObjID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}
	return perfil, nil
}
