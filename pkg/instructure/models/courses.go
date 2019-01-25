package models

type Course struct {
	// The unique identifier for the course.
	ID 				uint 	`json:"id"`

	// The SIS identifier for the course, if defined. This field is only
	// included if the user has permission to view SIS information.
	SisCourseID		string 	`json:"sis_course_id,omitempty"`

	// The UUID of the course.
	UUID			string 	`json:"uuid"`

	// The integration identifier for the course, if defined. This field is only
	// included if the user has permission to view SIS information.
	IntegrationID	string	`json:"integration_id,omitempty"`

	// The unique identifier for the SIS import. This field is only included if
	// the user has permission to manage SIS information.
	SisImportID		uint 	`json:"sis_import_id,omitempty"`

	// The full name of the course.
	Name 			string 	`json:"name"`

	// The course code.
	CourseCode 		string 	`json:"course_code"`

	// The current state of the course one of 'unpublished', 'available',
	// 'completed', or 'deleted'.
	WorkflowState	string 	`json:"workflow_state"`

	// The account associated with the course.
	AccountID		uint	`json:"account_id,omitempty"`

	// The root account associated with the course.
	RootAccountID	uint	`json:"root_account_id,omitempty"`

	// The enrollment term associated with the course
	EnrollmentTermID	uint	`json:"enrollment_term_id,omitempty"`

	// The grading standard associated with the course
	GradingStandardID	uint	`json:"grading_standard_id,omitempty"`

	// The date the course was created.
	CreatedAt		string	`json:"created_at,omitempty"`

	// The start date for the course, if applicable.
	StartAt			string 	`json:"start_at,omitempty"`

	// The end date for the course, if applicable.
	EndAt			string 	`json:"end_at,omitempty"`

	// The course-set locale, if applicable.
	Locale			string	`json:"locale,omitempty"`

	// Optional: the total number of active and invited students in the course.
	TotalStudents	int32	`json:"total_students,omitempty"`

	SyllabusBody	string	`json:"syllabus_body,omitempty"`

	PublicDescription	string 	`json:"public_description,omitempty"`
	IsPublic		bool 	`json:"is_public"`
	TimeZone		string 	`json:"time_zone,omitempty"`
}