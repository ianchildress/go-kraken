package kraken

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

func wrap(err error) error {
	return errors.Wrap(err, callerLocation())
}

func callerLocation() string {
	_, file, no, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d", file, no)
	}
	return "failed to determine line number"
}

func mergeErrors(multiErrors []error) error {
	if len(multiErrors) > 0 {
		var errString string
		for _, err := range multiErrors {
			errString += err.Error() + " "
		}
		return wrap(errors.New(errString))
	}
	return nil
}
