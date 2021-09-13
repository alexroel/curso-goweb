package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWriter  http.ResponseWriter
}

//Crear un respuesta por defecto
func CreateDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWriter:  rw,
		contentType: "application/json",
	}
}

//Método para responder error
func (resp *Response) NoFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource Not Found"
}

//Responder al Cliente
func (resp *Response) Send() {
	//Modificar el encabezado
	resp.respWriter.Header().Set("Content-Type", resp.contentType)
	resp.respWriter.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWriter, string(output))
}

//Responde la data al cliente
func SendData(rw http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(rw)
	response.Data = data
	response.Send()
}

//Responder error al cliente
func SendNotFound(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.NoFound()
	response.Send()
}

func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity Not Found"
}
func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := CreateDefaultResponse(rw)
	response.UnprocessableEntity()
	response.Send()
}
