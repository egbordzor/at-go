package voice

// CallTransferPayload is the payload type of the TransferCall method.
type CallTransferPayload struct {

	// Session Id of the ongoing calls, it must have two legs
	SessionID string `form:"sessionId" json:"sessionId" xml:"sessionId"`

	// Phone Number to transfer the calls to
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`

	// Call leg to transfer the calls to either caller or callee
	CallLeg string `form:"callLeg,omitempty" json:"callLeg,omitempty" xml:"callLeg,omitempty"`

	// The url of the media file to be played when the user is on hold.
	HoldMusicURL string `form:"holdMusicUrl,omitempty" json:"holdMusicUrl,omitempty" xml:"holdMusicUrl,omitempty"`
}

// CallTransferResponse is the result type of the TransferCall method.
type CallTransferResponse struct {

	// can be either Success or Aborted
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`

	// Why the transfer ws aborted None is successful
	ErrorMessage string `form:"errorMessage,omitempty" json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
}
