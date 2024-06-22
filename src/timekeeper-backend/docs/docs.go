// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/get-remote": {
            "get": {
                "description": "Get a remote by name and version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "remote"
                ],
                "summary": "Get a remote by name and version",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Remote Name",
                        "name": "remoteName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Version",
                        "name": "version",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/remote.RemoteResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/remote.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/health-check": {
            "get": {
                "description": "Returns the status of the server",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/push-remote": {
            "post": {
                "description": "Push a new remote",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "remote"
                ],
                "summary": "Push a new remote",
                "parameters": [
                    {
                        "description": "Remote",
                        "name": "remote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/remote.PushRemoteRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/remote.RemoteResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/remote.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.SuccessResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "remote.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "remote.PushRemoteRequest": {
            "type": "object",
            "properties": {
                "baseUrl": {
                    "type": "string"
                },
                "remoteName": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "remote.RemoteResponse": {
            "type": "object",
            "properties": {
                "remoteName": {
                    "type": "string"
                },
                "remoteURL": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Timekeeper Backend API",
	Description:      "This is a server for timekeeper backend.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}