package resetstate

import (
	"github.com/bizflycloud/gophercloud"
)

// ResetResult is the response of a ResetState operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type ResetResult struct {
	gophercloud.ErrResult
}
