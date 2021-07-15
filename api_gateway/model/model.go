package model

// We need these structs for swagger modeling
type Student struct {
	Id         int
	First_Name string
	Last_Name  string
	Email      string
}

type CreateStudentModel struct {
	First_name string
	Last_name  string
	Email      string
}

type Course struct {
	CourseId  int32	`json:"course_id"`
	Instructor string
	Title      string
}

type Book struct {
	Id     int
	Title  string
	Author string
}

type CreateBookModel struct {
	Title  string
	Author string
}

type CreateCourseModel struct {
	Title      string
	Instructor string
}

type Student_Book_Ids struct {
	Book_id    int32
	Student_id int32
}

type Student_Course_Ids struct {
	Course_id  int32
	Student_id int32
}