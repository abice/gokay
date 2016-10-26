package gkgen

import (
	"fmt"

	"golang.org/x/net/context"
)

// PlaybackServicer is the interface for the Playback service.
type PlaybackServicer interface {
	Playback(context.Context, interface{}) (interface{}, error)
}

// PlaybackService is responsible for tranforming a Once Playback request into a Bolt Playback request
type PlaybackService struct{}

// NewPlaybackService creates a Playback Service with default options.
func NewPlaybackService() PlaybackServicer {
	return &PlaybackService{}
}

// PlaybackRequest is empty because a PlaybackRequest is a GET request.
type PlaybackRequest struct {
	// URL Deconstructed params
	DeliveryType      string            `valid:"NotEmpty"`
	RequestedFileType *string           `valid:"NotNil"`
	DomainID          string            `valid:"BCP47"`
	ApplicationID     string            `valid:"NotNil"`
	MediaItemID       string            `valid:"NotNil"`
	VirtualFileName   string            `valid:"NotNil"`
	Params            map[string]string `valid:"NotNil"`

	// These are maybes?

	// VisitID is some Guid for a visit?
	VisitID string
	// ContextID is the ContextGuid in UX
	ContextID string
}

// PlaybackResponse
type PlaybackResponse struct {
	Err error  `json:"err" valid:"NotNil"`
	MSG string `json:"msg"`
}

// Playback is responsible for handling Once API customers existing calls into the Once platform and
// proxying them into BOLT backend calls.
func (h *PlaybackService) Playback(ctx context.Context, i interface{}) (interface{}, error) {
	_, ok := i.(*PlaybackRequest)
	if !ok {
		return nil, fmt.Errorf("Unable to convert request to PlaybackRequest type")
	}

	return PlaybackResponse{MSG: "Playback"}, nil
}
