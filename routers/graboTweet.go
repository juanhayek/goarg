package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juanhayek/goarg/bd"
	"github.com/juanhayek/goarg/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha registrado el Tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
