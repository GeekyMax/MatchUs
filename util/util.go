package util

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

func MustToString(i interface{}) string {
	result, _ := jsoniter.MarshalToString(i)
	return result
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

func PanicIfError(err error) {
	if err != nil {
		Panicf("panic %v", err)
	}
}
func PanicIf(b bool) {
	if b {
		Panicf("panic!")
	}
}
