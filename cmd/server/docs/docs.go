// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/profile/{username}": {
            "get": {
                "description": "Retrieve a profile from clubhouse by given username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieve a profile from clubhouse by given username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ClubhouseProfileResponse"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ClubhouseProfileResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2021-03-15T15:20:18.066663Z"
                },
                "id": {
                    "type": "string",
                    "example": "604f7b3253883ed1649ed13a"
                },
                "profile": {
                    "type": "object",
                    "properties": {
                        "bio": {
                            "type": "string",
                            "example": "BackEnd Programmer. 선린인터넷고 정보보호과 3학년 재학중입니다.\n꽈뚜룹, 이근 성대모사 잘합니다.\n\nPython, Node, Go 관심 많습니다."
                        },
                        "clubs": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "properties": {
                                    "club_id": {
                                        "type": "integer",
                                        "example": 12345
                                    },
                                    "description": {
                                        "type": "string",
                                        "example": "this is a strip club"
                                    },
                                    "enable_private": {
                                        "type": "boolean",
                                        "example": true
                                    },
                                    "is_community": {
                                        "type": "boolean",
                                        "example": false
                                    },
                                    "is_follow_allowed": {
                                        "type": "boolean",
                                        "example": true
                                    },
                                    "is_membership_private": {
                                        "type": "boolean",
                                        "example": false
                                    },
                                    "name": {
                                        "type": "string",
                                        "example": "strip club"
                                    },
                                    "num_followers": {
                                        "type": "integer",
                                        "example": 100
                                    },
                                    "num_members": {
                                        "type": "integer",
                                        "example": 20
                                    },
                                    "num_online": {
                                        "type": "integer",
                                        "example": 3
                                    },
                                    "photo_url": {
                                        "type": "string",
                                        "example": "https://clubhouseprod.s3.amazonaws.com:443/club_\u003cclub_id\u003e_\u003cguid\u003e_thumbnail_250x250"
                                    },
                                    "rules": {
                                        "type": "array",
                                        "items": {
                                            "type": "object",
                                            "properties": {
                                                "desc": {
                                                    "type": "string",
                                                    "example": "description"
                                                },
                                                "title": {
                                                    "type": "string",
                                                    "example": "title"
                                                }
                                            }
                                        }
                                    },
                                    "url": {
                                        "type": "string",
                                        "example": "https://www.joinclubhouse.com/club/\u003cclubname\u003e"
                                    }
                                }
                            }
                        },
                        "displayname": {
                            "type": "string"
                        },
                        "instagram": {
                            "type": "string",
                            "example": "yeon.gyu.kim"
                        },
                        "invited_by_user_profile": {
                            "type": "object",
                            "properties": {
                                "name": {
                                    "type": "string",
                                    "example": "Go Taegeon"
                                },
                                "photo_url": {
                                    "type": "string",
                                    "example": "https://clubhouseprod.s3.amazonaws.com/5652354_2c1840a1-02c0-40ee-b98c-f45598712bcd_thumbnail_250x250"
                                },
                                "user_id": {
                                    "type": "integer",
                                    "example": 5652354
                                },
                                "username": {
                                    "type": "string",
                                    "example": "gtg7784"
                                }
                            }
                        },
                        "name": {
                            "type": "string",
                            "example": "YeonGyu Kim"
                        },
                        "num_followers": {
                            "type": "integer",
                            "example": 84
                        },
                        "num_following": {
                            "type": "integer",
                            "example": 163
                        },
                        "photo_url": {
                            "type": "string",
                            "example": "https://clubhouseprod.s3.amazonaws.com/711498010_dcbb5803-9985-4ec5-b36a-0ac86284ff92"
                        },
                        "time_created": {
                            "type": "string",
                            "example": "2021-02-08T15:03:08.132077+00:00"
                        },
                        "twitter": {
                            "type": "string"
                        },
                        "url": {
                            "type": "string",
                            "example": "https://www.joinclubhouse.com/@yeon.gyu.kim"
                        },
                        "username": {
                            "type": "string",
                            "example": "yeon.gyu.kim"
                        }
                    }
                },
                "updated_at": {
                    "type": "string",
                    "example": "2021-03-15T15:20:18.066664Z"
                },
                "user_id": {
                    "type": "integer",
                    "example": 711498010
                },
                "username": {
                    "type": "string",
                    "example": "yeon.gyu.kim"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
