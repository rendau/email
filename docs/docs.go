// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/send": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/types.SendReqSt"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dopTypes.ErrRep"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dopTypes.ErrRep": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "error_code": {
                    "type": "string"
                },
                "fields": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "types.SendReqSt": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "receivers": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "subject": {
                    "type": "string"
                },
                "sync": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
