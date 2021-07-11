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
        "/books/student/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Retrieves the borrowed books of the student based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
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
                                "$ref": "#/definitions/main.Book"
                            }
                        }
                    }
                }
            }
        },
        "/courses/student/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Retrieves the enrolled courses of the student based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
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
                                "$ref": "#/definitions/main.Course"
                            }
                        }
                    }
                }
            }
        },
        "/create/student": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Related"
                ],
                "summary": "Creates a student record with user input details and writes into database",
                "parameters": [
                    {
                        "description": "Student details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                }
            }
        },
        "/delete/student/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Related"
                ],
                "summary": "Deletes a student with the specified ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                }
            }
        },
        "/student/borrow/book": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Borrows the book of a given Id for a given student ID",
                "parameters": [
                    {
                        "description": "Book and Student IDs",
                        "name": "student_book_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student_Book_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/student/drop/course": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Drop the student of the given Id from the given course(ID)",
                "parameters": [
                    {
                        "description": "Student and Course IDs",
                        "name": "student_course_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student_Course_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Course"
                        }
                    }
                }
            }
        },
        "/student/enroll/course": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Enrolls the student of the given Id to the given course ID",
                "parameters": [
                    {
                        "description": "Student and Course IDs",
                        "name": "student_course_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student_Course_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Course"
                        }
                    }
                }
            }
        },
        "/student/handin/book": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Hands in the book of a given Id for a given student ID",
                "parameters": [
                    {
                        "description": "Book and Student IDs",
                        "name": "student_book_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student_Book_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Book"
                        }
                    }
                }
            }
        },
        "/student/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Related"
                ],
                "summary": "Retrieves user based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                }
            }
        },
        "/students": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Related"
                ],
                "summary": "Retrieves the list of all students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Student"
                            }
                        }
                    }
                }
            }
        },
        "/update/student/{id}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Related"
                ],
                "summary": "Updates a student record with user input details and writes into database",
                "parameters": [
                    {
                        "description": "Updated Student Details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Student"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "bookId": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "main.Course": {
            "type": "object",
            "properties": {
                "courseId": {
                    "type": "integer"
                },
                "instructor": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "main.Student": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "main.Student_Book_Ids": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "student_id": {
                    "type": "integer"
                }
            }
        },
        "main.Student_Course_Ids": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "integer"
                },
                "student_id": {
                    "type": "integer"
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
