package main

import (
	"fmt"
	"infraestructure-user-action/graphql"
)

func main() {

	queryInstance := graphql.GetInstanceQuery(instanceName)
	instanceResponse, err := handleResponse(queryInstance)
	if err != nil {
		fmt.Println("Error obteniendo ID de la instancia:", err)
		return
	}

	instanceId, err := extractInstanceID(instanceResponse)
	if err != nil {
		fmt.Println("No se pudo extraer el ID de la instancia:", err)
		return
	}

	query := graphql.CreateMutation(instanceId, user)
	output, err := handleResponse(query)
	if err != nil {
		fmt.Println("Error handling error response:", err)
		return
	}

	errCode, err := handleError(output)
	if err != nil {
		fmt.Println("Error handling error:", err)
		return
	}
	if errCode != "" {
		jsonOutput := handleErrorCode(errCode)
		fmt.Println(jsonOutput)
		return
	}

	jsonOutput, _ := requestToJson(output)
	fmt.Println(jsonOutput)
	return
}
