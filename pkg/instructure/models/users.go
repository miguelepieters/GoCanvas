package models


// This mini-object is used for secondary user responses, when we just want to
// provide enough information to display a user.
type UserDisplay struct {
	// The ID of the user.
	ID				uint 	`json:"id"`

	// A short name the user has selected, for use in conversations or other
	// less formal places through the site.
	ShortName		string	`json:"short_name"`

	// If avatars are enabled, this field will be included and contain a url to
	// retrieve the user's avatar.
	AvatarImageURL	string	`json:"avatar_image_url"`

	// URL to access user, either nested to a context or directly.
	HtmlUrl			string 	`json:"html_url"`
}

// A Canvas user, e.g. a student, teacher, administrator, observer, etc.
type User struct {
	// The ID of the user.
	ID				uint 	`json:"id"`

	// The name of the user.
	Name 			string 	`json:"name,omitempty"`

	// The name of the user that is should be used for sorting groups of users,
	// such as in the gradebook.
	SortableName	string 	`json:"sortable_name"`

	// A short name the user has selected, for use in conversations or other
	// less formal places through the site.
	ShortName		string	`json:"short_name"`

	// The SIS ID associated with the user.  This field is only included if the
	// user came from a SIS import and has permissions to view SIS information.
	SisUserID		string 	`json:"sis_user_id,omitempty"`

	// The id of the SIS import.  This field is only included if the user came
	// from a SIS import and has permissions to manage SIS information.
	SisImportID		uint	`json:"sis_import_id,omitempty"`

	// The integration_id associated with the user.  This field is only included
	// if the user came from a SIS import and has permissions to view SIS
	// information.
	IntegrationID	string	`json:"integration_id,omitempty"`

	// The unique login id for the user.  This is what the user uses to log in
	// to Canvas.
	LoginID			string	`json:"login_id,omitempty"`

	// If avatars are enabled, this field will be included and contain a url to
	// retrieve the user's avatar.
	AvatarURL		string	`json:"avatar_url"`

	// Optional: This field can be requested with certain API calls, and will
	// return the users primary email address.
	Email			string 	`json:"email,omitempty"`

	// Optional: This field can be requested with certain API calls, and will
	// return the users locale in RFC 5646 format.
	Locale 			string 	`json:"locale,omitempty"`

	// Optional: This field is only returned in certain API calls, and will
	// return a timestamp representing the last time the user logged in to
	// canvas.
	LastLogin		string 	`json:"last_login,omitempty"`

	// Optional: This field is only returned in certain API calls, and will
	// return the IANA time zone name of the user's preferred timezone.
	TimeZone		string	`json:"time_zone,omitempty"`

	// Optional: The user's bio.
	Bio				string 	`json:"bio,omitempty"`
}

type Profile struct {
	// The ID of the user.
	ID				uint 	`json:"id"`

	// The name of the user.
	Name 			string 	`json:"name,omitempty"`

	// The name of the user that is should be used for sorting groups of users,
	// such as in the gradebook.
	SortableName	string 	`json:"sortable_name"`

	// A short name the user has selected, for use in conversations or other
	// less formal places through the site.
	ShortName		string	`json:"short_name"`

	// The SIS ID associated with the user.  This field is only included if the
	// user came from a SIS import and has permissions to view SIS information.
	SisUserID		string 	`json:"sis_user_id,omitempty"`

	// The unique login id for the user.  This is what the user uses to log in
	// to Canvas.
	LoginID			string	`json:"login_id,omitempty"`

	// If avatars are enabled, this field will be included and contain a url to
	// retrieve the user's avatar.
	AvatarURL		string	`json:"avatar_url"`

	// Optional: This field can be requested with certain API calls, and will
	// return the users primary email address.
	PrimaryEmail	string 	`json:"primary_email,omitempty"`

	// Optional: This field can be requested with certain API calls, and will
	// return the users locale in RFC 5646 format.
	Locale 			string 	`json:"locale,omitempty"`

	// Optional: This field is only returned in certain API calls, and will
	// return the IANA time zone name of the user's preferred timezone.
	TimeZone		string	`json:"time_zone,omitempty"`

	// Optional: The user's bio.
	Bio				string 	`json:"bio,omitempty"`
}