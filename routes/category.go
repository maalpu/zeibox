package routes

import (
	"encoding/json"
	"strconv"

	//"github.com/aws/aws-lambda-go/events"
	"github.com/maalpu/zeibox/bd"
	"github.com/maalpu/zeibox/models"
)

func InsertCategory(body string, User string) (int, string) {
	var c models.Category
	err := json.Unmarshal([]byte(body), &c)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}
	if len(c.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}
	if len(c.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertarCategory(c)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar registrar una categoría " + c.CategName + " > " + err2.Error()
	}

	return 200, "{ CategId: " + strconv.Itoa(int(result)) + "}"
}
