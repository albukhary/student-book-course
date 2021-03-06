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
        "/book/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Gets details of the Book from User input book ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Gets the list of all the books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Book"
                            }
                        }
                    }
                }
            }
        },
        "/books/student/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Student Book Related"
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
                                "$ref": "#/definitions/model.Book"
                            }
                        }
                    }
                }
            }
        },
        "/course/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Gets details of the Course from User input book ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                }
            }
        },
        "/courses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Gets the list of all the courses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Course"
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
                    "Student Course Related"
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
                                "$ref": "#/definitions/model.Course"
                            }
                        }
                    }
                }
            }
        },
        "/create/book": {
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
                "summary": "Creates a Book record with user input details and writes into database",
                "parameters": [
                    {
                        "description": "Book details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateBookModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                }
            }
        },
        "/create/course": {
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
                "summary": "Creates a Course record with user input details and writes into database",
                "parameters": [
                    {
                        "description": "Course details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateCourseModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
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
                            "$ref": "#/definitions/model.CreateStudentModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                }
            }
        },
        "/delete/book/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Deletes the book with the given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                }
            }
        },
        "/delete/course/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Deletes the course with the given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
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
                            "$ref": "#/definitions/model.Student"
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
                    "Student Book Related"
                ],
                "summary": "Borrows the book of a given Id for a given student ID",
                "parameters": [
                    {
                        "description": "Book and Student IDs",
                        "name": "student_book_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student_Book_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
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
                    "Student Course Related"
                ],
                "summary": "Drop the student of the given Id from the given course(ID)",
                "parameters": [
                    {
                        "description": "Student and Course IDs",
                        "name": "student_course_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student_Course_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
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
                    "Student Course Related"
                ],
                "summary": "Enrolls the student of the given Id to the given course ID",
                "parameters": [
                    {
                        "description": "Student and Course IDs",
                        "name": "student_course_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student_Course_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
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
                    "Student Book Related"
                ],
                "summary": "Hands in the book of a given Id for a given student ID",
                "parameters": [
                    {
                        "description": "Book and Student IDs",
                        "name": "student_book_ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Student_Book_Ids"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
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
                            "$ref": "#/definitions/model.Student"
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
                                "$ref": "#/definitions/model.Student"
                            }
                        }
                    }
                }
            }
        },
        "/students/book/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Gets the list of all students who borrowed a particular book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Book ID",
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
                                "$ref": "#/definitions/model.Student"
                            }
                        }
                    }
                }
            }
        },
        "/students/course/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Gets the list of all students who are enrolled to a particular course",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
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
                                "$ref": "#/definitions/model.Student"
                            }
                        }
                    }
                }
            }
        },
        "/update/book": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book Related"
                ],
                "summary": "Updates the book record with user input details",
                "parameters": [
                    {
                        "description": "Book details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Book"
                        }
                    }
                }
            }
        },
        "/update/course": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course Related"
                ],
                "summary": "Updates the course record with user input details",
                "parameters": [
                    {
                        "description": "Course details",
                        "name": "details",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Course"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Course"
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
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Student"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.Course": {
            "type": "object",
            "properties": {
                "course_id": {
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
        "model.CreateBookModel": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CreateCourseModel": {
            "type": "object",
            "properties": {
                "instructor": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CreateStudentModel": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        },
        "model.Student": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_Name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_Name": {
                    "type": "string"
                }
            }
        },
        "model.Student_Book_Ids": {
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
        "model.Student_Course_Ids": {
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
