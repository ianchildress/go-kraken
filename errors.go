package kraken

import (
	"fmt"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

type KrakenErrors []string

func (ke KrakenErrors) Errors() error {
	// if empty, return nil
	if len(ke) == 0 {
		return nil
	}

	// return errors
	var err string
	for _, s := range ke {
		err += s + " "
	}
	return errors.New(err)
}

func Wrap(err error) error {
	return errors.Wrap(err, callerLocation())
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func callerLocation() string {
	_, file, no, ok := runtime.Caller(2)
	if ok {
		return fmt.Sprintf("%s:%d", file, no)
	}
	return "failed to determine line number"
}
