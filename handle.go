// Manejo de errores
package main

import (
	"encoding/json"
	"fmt"
	"infraestructure-user-action/graphql"
	"infraestructure-user-action/models"
	"infraestructure-user-action/utils"
)

func handleResponse(query models.GraphQLQuery) (response []byte, err error) {
	response, err = utils.SendGraphQLQuery(url, query)
	if err != nil {
		fmt.Println("Error sending GraphQL query for existing user:", err)
		return
	}
	return
}

func handleError(response []byte) (errorCode string, err error) {
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(response, &jsonResponse)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}
	if errors, ok := jsonResponse["errors"].([]interface{}); ok {
		for _, e := range errors {
			if errMap, ok := e.(map[string]interface{}); ok {
				if extensions, ok := errMap["extensions"].(map[string]interface{}); ok {
					if code, ok := extensions["code"].(string); ok {
						errorCode = code
						return
					}
				}
			}
		}
		return
	}
	return
}
func handleQuery(request []byte) (response []byte, err error) {
	var res map[string]interface{}
	err = json.Unmarshal(request, &res)
	if err != nil {
		return
	}
	if data, ok := res["data"].(map[string]interface{}); ok {
		if credentials, ok := data[alias].(map[string]interface{}); ok {
			if edges, ok := credentials["edges"].([]interface{}); ok && len(edges) > 0 {
				if firstEdge, ok := edges[0].(map[string]interface{}); ok {
					res["data"].(map[string]interface{})[alias] = firstEdge
				}
			}
		}
	}
	response, _ = json.Marshal(res)
	return
}

func requestToJson(request []byte) (jsonOutput string, err error) {
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(request, &jsonResponse)
	if err != nil {
		errMessage := "Error unmarshaling response:"
		jsonOutput = errMessage + err.Error()
		err = fmt.Errorf(errMessage, err)
		return
	}
	output, err := json.MarshalIndent(jsonResponse, "", "  ")
	if err != nil {
		jsonOutput = "Error formatting JSON response:" + err.Error()
		return
	}
	jsonOutput = string(output)
	return
}

func handleErrorCode(errCode string) (message string) {
	switch errCode {
	case catchErrCode:
		query := graphql.GetQuery(user)
		output, err := handleResponse(query)
		if err != nil {
			message = fmt.Sprint("Error handling error response:", err)
			return
		}
		errNewCode, err := handleError(output)
		if err != nil || errNewCode != "" {
			message = fmt.Sprint("Error handling error response:", err)
			return
		}
		response, err := handleQuery(output)
		if err != nil {
			message = fmt.Sprint("Error handling query:", err)
			return
		}
		jsonOutput, _ := requestToJson(response)
		message = fmt.Sprint(jsonOutput)
		return
	default:
		message = fmt.Sprint("Error handling error response code:", errCode)
		return
	}
}

func extractInstanceID(response []byte) (string, error) {
	var jsonResponse map[string]interface{}
	err := json.Unmarshal(response, &jsonResponse)
	if err != nil {
		return "", fmt.Errorf("error parseando JSON: %v", err)
	}

	// Acceder a la estructura "data" -> "instance" -> "edges"
	if data, ok := jsonResponse["data"].(map[string]interface{}); ok {
		if instance, ok := data["instance"].(map[string]interface{}); ok {
			if edges, ok := instance["edges"].([]interface{}); ok && len(edges) > 0 {
				if firstEdge, ok := edges[0].(map[string]interface{}); ok {
					if id, ok := firstEdge["_id"].(string); ok {
						return id, nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("no se encontró el ID de la instancia")
}
