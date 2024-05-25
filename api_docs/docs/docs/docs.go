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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/accounts": {
            "get": {
                "description": "Get account by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get account by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetAccountResDto"
                        }
                    }
                }
            },
            "post": {
                "description": "Create account",
                "produces": [
                    "application/json"
                ],
                "summary": "Create account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateAccountResDto"
                        }
                    }
                }
            }
        },
        "/v1/ping": {
            "get": {
                "description": "Ping handler",
                "produces": [
                    "application/json"
                ],
                "summary": "Ping handler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.PongResponse"
                        }
                    }
                }
            }
        },
        "/v1/transactions": {
            "post": {
                "description": "Create transaction",
                "produces": [
                    "application/json"
                ],
                "summary": "Create transaction",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateTransactionResDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.PongResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.CreateAccountResDto": {
            "type": "object",
            "required": [
                "account_id",
                "balance"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "error": {
                    "type": "string"
                },
                "failure_code": {
                    "$ref": "#/definitions/models.FailureCodeType"
                }
            }
        },
        "dto.CreateTransactionResDto": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "destination_account_id": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                },
                "failure_code": {
                    "$ref": "#/definitions/models.FailureCodeType"
                },
                "object": {
                    "type": "string"
                },
                "source_account_id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.GetAccountResDto": {
            "type": "object",
            "required": [
                "account_id",
                "balance"
            ],
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "error": {
                    "type": "string"
                },
                "failure_code": {
                    "$ref": "#/definitions/models.FailureCodeType"
                }
            }
        },
        "models.FailureCodeType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                6,
                7
            ],
            "x-enum-comments": {
                "FailureCodeNil": "FailureCode"
            },
            "x-enum-varnames": [
                "FailureCodeNil",
                "FailureCodeParseRequest",
                "FailureCodeNotFound",
                "FailureCodeCreateResponse",
                "FailureCodeServiceFailed",
                "FailureCodeCreateUser",
                "FailureCodePaymentMethodNotFound",
                "FailureCodeInvalidPartnerConfig"
            ]
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Transfers System",
	Description:      "This is an http-request collection of Transfers System.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
