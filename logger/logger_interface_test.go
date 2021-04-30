package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetLogger(t *testing.T) {
	l := logrus.New()
	SetLogger(l)

	assert.Equal(t, l, Log)
}
