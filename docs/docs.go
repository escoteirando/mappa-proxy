// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/": {
            "get": {
                "description": "route: /",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Index handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.IndexResponse"
                        }
                    }
                }
            }
        },
        "/hc": {
            "get": {
                "description": "Service healthcheck",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Healthcheck handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.HealthCheckResponse"
                        }
                    }
                }
            }
        },
        "/mappa": {
            "get": {
                "description": "Send a request to Mappa API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Mappa Generic Handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/mappa/escotista/{userId}": {
            "get": {
                "description": "Detalhes do escotista",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "MappaEscotista handler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaDetalhesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/login": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "Mappa Login handler",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.MappaLoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        },
        "/mappa/progressoes/{ramo}": {
            "get": {
                "description": "Lista de progressões do ramo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "mappa-proxy"
                ],
                "summary": "MappaProgressoes handler",
                "parameters": [
                    {
                        "enum": [
                            "L",
                            "E",
                            "S",
                            "P"
                        ],
                        "type": "string",
                        "description": "Ramo",
                        "name": "ramo",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.MappaProgressaoResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ReplyMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ReplyMessage": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "infra.MemoryStatus": {
            "type": "object",
            "properties": {
                "alloc": {
                    "type": "integer"
                },
                "heap_alloc": {
                    "type": "integer"
                },
                "total_alloc": {
                    "type": "integer"
                }
            }
        },
        "requests.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.HealthCheckResponse": {
            "type": "object",
            "properties": {
                "mappa_server": {
                    "$ref": "#/definitions/responses.MappaServerResponse"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "responses.IndexResponse": {
            "type": "object",
            "properties": {
                "app": {
                    "type": "string"
                },
                "running-by": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "responses.MappaAssociadoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoCategoria": {
                    "type": "integer"
                },
                "codigoEquipe": {
                    "type": "integer"
                },
                "codigoFoto": {
                    "type": "integer"
                },
                "codigoRamo": {
                    "type": "integer"
                },
                "codigoRamoAdulto": {
                    "type": "integer"
                },
                "codigoSegundaCategoria": {
                    "type": "integer"
                },
                "codigoTerceiraCategoria": {
                    "type": "integer"
                },
                "dataAcompanhamento": {
                    "type": "string"
                },
                "dataNascimento": {
                    "type": "string"
                },
                "dataValidade": {
                    "type": "string"
                },
                "linhaFormacao": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "nomeAbreviado": {
                    "type": "string"
                },
                "numeroDigito": {
                    "type": "integer"
                },
                "sexo": {
                    "type": "string"
                },
                "username": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaDetalhesResponse": {
            "type": "object",
            "properties": {
                "associado": {
                    "$ref": "#/definitions/responses.MappaAssociadoResponse"
                },
                "escotista": {
                    "$ref": "#/definitions/responses.MappaEscotistaResponse"
                },
                "grupos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.MappaGrupoResponse"
                    }
                }
            }
        },
        "responses.MappaEscotistaResponse": {
            "type": "object",
            "properties": {
                "ativo": {
                    "$ref": "#/definitions/types.Bool"
                },
                "codigo": {
                    "type": "integer"
                },
                "codigoAssociado": {
                    "type": "integer"
                },
                "codigoFoto": {
                    "type": "integer"
                },
                "codigoGrupo": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "nomeCompleto": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "responses.MappaGrupoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoModalidade": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "responses.MappaLoginResponse": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "ttl": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "responses.MappaProgressaoResponse": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "integer"
                },
                "codigoCaminho": {
                    "type": "integer"
                },
                "codigoCompetencia": {
                    "type": "integer"
                },
                "codigoDesenvolvimento": {
                    "type": "integer"
                },
                "codigoRegiao": {
                    "type": "string"
                },
                "codigoUeb": {
                    "type": "string"
                },
                "descricao": {
                    "type": "string"
                },
                "numeroGrupo": {
                    "type": "integer"
                },
                "ordenacao": {
                    "type": "integer"
                },
                "segmento": {
                    "type": "string"
                }
            }
        },
        "responses.MappaServerResponse": {
            "type": "object",
            "properties": {
                "mappa_server_url": {
                    "type": "string"
                },
                "memory": {
                    "$ref": "#/definitions/infra.MemoryStatus"
                },
                "status": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "types.Bool": {
            "type": "object"
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
