package graphql

import "infraestructure-user-action/models"

func CreateMutation(instance, user string) models.GraphQLQuery {
	return models.GraphQLQuery{
		Query: `
			mutation ($input: CreateInstanCredentialInput) {
				user: createInstanceCredential(input: $input) {
					_id
					user
					password
					key
					instance {
						_id
					}
				}
			}
		`,
		Variables: map[string]interface{}{
			"input": map[string]interface{}{
				"instance":         instance,
				"user":             user,
				"isRoot":           false,
				"createInInstance": true,
			},
		},
	}
}
