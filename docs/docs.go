// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "support@thiel-inet.de"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://opensource.org/license/bsd-3-clause/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/hero": {
            "get": {
                "description": "Get all heroes.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hero"
                ],
                "summary": "get all heroes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.HeroSet"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorData"
                        }
                    }
                }
            }
        },
        "/v1/hero/{id}": {
            "get": {
                "description": "Get hero by given ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hero"
                ],
                "summary": "get hero by given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 2,
                        "description": "Hero ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.HeroSet"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.HeroSet": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string",
                    "example": "Wade"
                },
                "heroName": {
                    "type": "string",
                    "example": "Deadpool"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "lastName": {
                    "type": "string",
                    "example": "Wilson"
                }
            }
        },
        "response.ErrorData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "10001"
                },
                "msg": {
                    "type": "string",
                    "example": "something went wrong"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "This is a demonstration application showing a working example of several functions and performance features.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
