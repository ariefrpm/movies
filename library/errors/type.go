package errors

import "errors"

//Please do not use Newr() here, use only errors.New()
//because will confuse the callers stack with the original callers
var ErrTrackingNotExist = errors.New("last tracking data not exist")
var ErrCreateDeliveryLocked = errors.New("create delivery already locked")
