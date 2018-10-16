// log writer test
package logger

import (
	"testing"
)

func TestWrite(t *testing.T) {
	val := "test log message"
	err := Write(val)
	if err != nil {
		t.Errorf(err.Error())
	}
}
