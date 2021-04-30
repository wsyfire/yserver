package logger

import (
	"log"
)

type DefaultLogger struct {
}

func newDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (d *DefaultLogger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (d *DefaultLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (d *DefaultLogger) Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func (d *DefaultLogger) Debug(v ...interface{}) {
	log.Print(v...)
}

func (d *DefaultLogger) Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (d *DefaultLogger) Debugln(v ...interface{}) {
	log.Println(v...)
}

func (d *DefaultLogger) Error(v ...interface{}) {
	log.Fatal(v...)
}

func (d *DefaultLogger) Errorf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func (d *DefaultLogger) Errorln(v ...interface{}) {
	log.Fatalln(v...)
}

func (d *DefaultLogger) Info(v ...interface{}) {
	log.Print(v...)
}

func (d *DefaultLogger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func (d *DefaultLogger) Infoln(v ...interface{}) {
	log.Println(v...)
}

func (d *DefaultLogger) Warn(v ...interface{}) {
	log.Print(v...)
}

func (d *DefaultLogger) Warnf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (d *DefaultLogger) Warnln(v ...interface{}) {
	log.Println(v...)
}
