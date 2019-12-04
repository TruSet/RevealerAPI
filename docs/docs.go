// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-11 10:25:44.064663 +0000 GMT m=+0.903901074

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "ToS TBD",
        "contact": {
            "name": "API Support",
            "url": "https://www.truset.com/contact/",
            "email": "greg.taschuk@consensys.net"
        },
        "license": {
            "name": "License TBD",
            "url": "TBD"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/commitments/": {
            "post": {
                "description": "Save a vote and the matching hash commitment to that vote",
                "produces": [
                    "application/json"
                ],
                "summary": "Store a commitment privately, to ensure it can be revealed at a later date",
                "operationId": "store-commitment",
                "parameters": [
                    {
                        "description": "The (about to be) committed vote details data you would like to store",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/database.CommitmentBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/database.Response"
                        }
                    },
                    "406": {
                        "description": "Bad payload",
                        "schema": {
                            "$ref": "#/definitions/database.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.CommitmentBody": {
            "type": "object",
            "properties": {
                "commitHash": {
                    "type": "string",
                    "example": "0x12345678909876543210123456789012"
                },
                "pollID": {
                    "type": "string",
                    "example": "0x12345678901234567890123456789012"
                },
                "salt": {
                    "type": "integer",
                    "example": 5866984321541876564
                },
                "voteOption": {
                    "type": "integer",
                    "example": 1
                },
                "voterAddress": {
                    "type": "string",
                    "example": "0x11223344556677889900"
                }
            }
        },
        "database.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "vote will be revealed when voting closes"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "/revealer/v0.1",
	Schemes:     []string{},
	Title:       "TruSet Revealer API",
	Description: "A REST interface for submitting votes (e.g. Accept or Reject) on proposed data records.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
