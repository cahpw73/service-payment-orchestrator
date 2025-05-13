package errors

import "fmt"

type AppError struct {
	Code       string
	Message    string
	ErrorTrace error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// Map de errores
var ErrorMessages = map[string]AppError{
	"ERR001": {Code: "ERR001", Message: "Error al obtener categorías"},
	"ERR002": {Code: "ERR002", Message: "Error al obtener Ciudades"},
	"ERR003": {Code: "ERR003", Message: "Parámetros Inválidos - Solicitud de Deuda"},
	"ERR004": {Code: "ERR004", Message: "Error al obtener las deudas"},
	"ERR005": {Code: "ERR005", Message: "Parámetros Inválidos - Pago de Deuda"},
	"ERR006": {Code: "ERR006", Message: "Error al procesar el pago de la dueda"},
	"ERR007": {Code: "ERR007", Message: "Error al obtener Criterios de Búsqueda"},
	"ERR008": {Code: "ERR008", Message: "Error al obtener servicios por subcategorias"},
	"ERR009": {Code: "ERR009", Message: "Error al obtener SubCategorias"},
	"ERR010": {Code: "ERR010", Message: "Parámetros Inválidos - Validación de Datos para Afiliación"},
	"ERR011": {Code: "ERR011", Message: "Error al Validar Datos para Afiliación"},
	"ERR012": {Code: "ERR012", Message: "Error al Conectarse a la Base de Datos"},
	"ERR013": {Code: "ERR013", Message: "Error al obtener Nro. de Persona"},
	"ERR014": {Code: "ERR014", Message: "Error al obtener Afiliaciones por Nro. de Persona"},
	"ERR015": {Code: "ERR015", Message: "Error al registrar Afiliacion"},
	"ERR016": {Code: "ERR016", Message: "Error al eliminar Afiliacion"},
	"ERR017": {Code: "ERR017", Message: "Error al actualizar Afiliacion"},
}

func GetError(code string) AppError {
	if err, exists := ErrorMessages[code]; exists {
		return err
	}

	return AppError{Code: "ERR000", Message: "Ocurrió un error en la aplicación"}
}
