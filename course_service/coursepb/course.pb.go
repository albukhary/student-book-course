// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: course_service/coursepb/course.proto

package __coursepb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId   int32  `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Instructor string `protobuf:"bytes,2,opt,name=instructor,proto3" json:"instructor,omitempty"`
	Title      string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Course) GetInstructor() string {
	if x != nil {
		return x.Instructor
	}
	return ""
}

func (x *Course) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type CreateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *CreateCourseRequest) Reset() {
	*x = CreateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseRequest) ProtoMessage() {}

func (x *CreateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseRequest.ProtoReflect.Descriptor instead.
func (*CreateCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCourseRequest) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type CreateCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *CreateCourseResponse) Reset() {
	*x = CreateCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseResponse) ProtoMessage() {}

func (x *CreateCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseResponse.ProtoReflect.Descriptor instead.
func (*CreateCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{3}
}

func (x *GetCourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *GetCourseResponse) Reset() {
	*x = GetCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseResponse) ProtoMessage() {}

func (x *GetCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseResponse.ProtoReflect.Descriptor instead.
func (*GetCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{4}
}

func (x *GetCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCourseRequest) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type UpdateCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *UpdateCourseResponse) Reset() {
	*x = UpdateCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseResponse) ProtoMessage() {}

func (x *UpdateCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseResponse.ProtoReflect.Descriptor instead.
func (*UpdateCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type DeleteCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCourseRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *DeleteCourseResponse) Reset() {
	*x = DeleteCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseResponse) ProtoMessage() {}

func (x *DeleteCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseResponse.ProtoReflect.Descriptor instead.
func (*DeleteCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCourseResponse) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type GetAllCoursesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *GetAllCoursesResponse) Reset() {
	*x = GetAllCoursesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCoursesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCoursesResponse) ProtoMessage() {}

func (x *GetAllCoursesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCoursesResponse.ProtoReflect.Descriptor instead.
func (*GetAllCoursesResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllCoursesResponse) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

// Get Enrolled Students - started
type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{10}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Student) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetEnrolledStudentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int32 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *GetEnrolledStudentsRequest) Reset() {
	*x = GetEnrolledStudentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnrolledStudentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrolledStudentsRequest) ProtoMessage() {}

func (x *GetEnrolledStudentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrolledStudentsRequest.ProtoReflect.Descriptor instead.
func (*GetEnrolledStudentsRequest) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{11}
}

func (x *GetEnrolledStudentsRequest) GetCourseId() int32 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type GetEnrolledStudentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetEnrolledStudentsResponse) Reset() {
	*x = GetEnrolledStudentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_service_coursepb_course_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnrolledStudentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnrolledStudentsResponse) ProtoMessage() {}

func (x *GetEnrolledStudentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_service_coursepb_course_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnrolledStudentsResponse.ProtoReflect.Descriptor instead.
func (*GetEnrolledStudentsResponse) Descriptor() ([]byte, []int) {
	return file_course_service_coursepb_course_proto_rawDescGZIP(), []int{12}
}

func (x *GetEnrolledStudentsResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_course_service_coursepb_course_proto protoreflect.FileDescriptor

var file_course_service_coursepb_course_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x06, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x41,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x73, 0x22, 0x6b, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x39,
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xe6, 0x03, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_course_service_coursepb_course_proto_rawDescOnce sync.Once
	file_course_service_coursepb_course_proto_rawDescData = file_course_service_coursepb_course_proto_rawDesc
)

func file_course_service_coursepb_course_proto_rawDescGZIP() []byte {
	file_course_service_coursepb_course_proto_rawDescOnce.Do(func() {
		file_course_service_coursepb_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_service_coursepb_course_proto_rawDescData)
	})
	return file_course_service_coursepb_course_proto_rawDescData
}

var file_course_service_coursepb_course_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_course_service_coursepb_course_proto_goTypes = []interface{}{
	(*Course)(nil),                      // 0: course.Course
	(*CreateCourseRequest)(nil),         // 1: course.CreateCourseRequest
	(*CreateCourseResponse)(nil),        // 2: course.CreateCourseResponse
	(*GetCourseRequest)(nil),            // 3: course.GetCourseRequest
	(*GetCourseResponse)(nil),           // 4: course.GetCourseResponse
	(*UpdateCourseRequest)(nil),         // 5: course.UpdateCourseRequest
	(*UpdateCourseResponse)(nil),        // 6: course.UpdateCourseResponse
	(*DeleteCourseRequest)(nil),         // 7: course.DeleteCourseRequest
	(*DeleteCourseResponse)(nil),        // 8: course.DeleteCourseResponse
	(*GetAllCoursesResponse)(nil),       // 9: course.GetAllCoursesResponse
	(*Student)(nil),                     // 10: course.Student
	(*GetEnrolledStudentsRequest)(nil),  // 11: course.GetEnrolledStudentsRequest
	(*GetEnrolledStudentsResponse)(nil), // 12: course.GetEnrolledStudentsResponse
	(*empty.Empty)(nil),                 // 13: google.protobuf.Empty
}
var file_course_service_coursepb_course_proto_depIdxs = []int32{
	0,  // 0: course.CreateCourseRequest.course:type_name -> course.Course
	0,  // 1: course.CreateCourseResponse.course:type_name -> course.Course
	0,  // 2: course.GetCourseResponse.course:type_name -> course.Course
	0,  // 3: course.UpdateCourseRequest.course:type_name -> course.Course
	0,  // 4: course.UpdateCourseResponse.course:type_name -> course.Course
	0,  // 5: course.DeleteCourseResponse.course:type_name -> course.Course
	0,  // 6: course.GetAllCoursesResponse.courses:type_name -> course.Course
	10, // 7: course.GetEnrolledStudentsResponse.students:type_name -> course.Student
	1,  // 8: course.CourseService.CreateCourse:input_type -> course.CreateCourseRequest
	3,  // 9: course.CourseService.GetCourse:input_type -> course.GetCourseRequest
	5,  // 10: course.CourseService.UpdateCourse:input_type -> course.UpdateCourseRequest
	7,  // 11: course.CourseService.DeleteCourse:input_type -> course.DeleteCourseRequest
	13, // 12: course.CourseService.GetAllCourses:input_type -> google.protobuf.Empty
	11, // 13: course.CourseService.GetEnrolledStudents:input_type -> course.GetEnrolledStudentsRequest
	2,  // 14: course.CourseService.CreateCourse:output_type -> course.CreateCourseResponse
	4,  // 15: course.CourseService.GetCourse:output_type -> course.GetCourseResponse
	6,  // 16: course.CourseService.UpdateCourse:output_type -> course.UpdateCourseResponse
	8,  // 17: course.CourseService.DeleteCourse:output_type -> course.DeleteCourseResponse
	9,  // 18: course.CourseService.GetAllCourses:output_type -> course.GetAllCoursesResponse
	12, // 19: course.CourseService.GetEnrolledStudents:output_type -> course.GetEnrolledStudentsResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_course_service_coursepb_course_proto_init() }
func file_course_service_coursepb_course_proto_init() {
	if File_course_service_coursepb_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_service_coursepb_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCourseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCourseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCoursesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnrolledStudentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_service_coursepb_course_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnrolledStudentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_service_coursepb_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_service_coursepb_course_proto_goTypes,
		DependencyIndexes: file_course_service_coursepb_course_proto_depIdxs,
		MessageInfos:      file_course_service_coursepb_course_proto_msgTypes,
	}.Build()
	File_course_service_coursepb_course_proto = out.File
	file_course_service_coursepb_course_proto_rawDesc = nil
	file_course_service_coursepb_course_proto_goTypes = nil
	file_course_service_coursepb_course_proto_depIdxs = nil
}
