package voice


// UploadMediaFile is the payload type of the UploadMedia method.
type UploadMediaFile struct {

	// Your Africa’s Talking application username.
	Username string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`

	// The url of the file to upload.
	URL string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}