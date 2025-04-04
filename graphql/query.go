package graphql

import "infraestructure-user-action/models"

func GetQuery(user string) models.GraphQLQuery {
	return models.GraphQLQuery{
		Query: `
			query getInstanceCredentials($input: [SearchInput]) {
				user: getInstanceCredentials@search(input: $input) {
					edges {
						_id
						user
						password
						key
						instance {
							_id
						}
					}
				}
			}
		`,
		Variables: map[string]interface{}{
			"input": []map[string]interface{}{
				{
					"field": "user",
					"value": []map[string]interface{}{
						{"value": user},
					},
				},
			},
		},
	}
}

func GetInstanceQuery(instance string) models.GraphQLQuery {
	return models.GraphQLQuery{
		Query: `
			query getInstance($input: [SearchInput]) {
				instance: getInstances@search(input: $input) {
					edges {
						_id
						name
					}
				}
			}
		`,
		Variables: map[string]interface{}{
			"input": []map[string]interface{}{
				{
					"field": "name",
					"value": []map[string]interface{}{
						{"value": instance},
					},
				},
			},
		},
	}
}
