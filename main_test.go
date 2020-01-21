package main

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/actions-go/toolkit/core"
)

const (
	errorOutput = `::debug::Waiting  milliseconds
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::error::strconv.Atoi: parsing "": invalid syntax
`
	successOutput = `::debug::Waiting 10 milliseconds
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::debug::2020-01-10 20:10:20.000000001 +0000 UTC
::set-output name=time::2020-01-10 20:10:20.000000001 +0000 UTC
`
)

func TestRunMain(t *testing.T) {
	w := bytes.NewBuffer(nil)
	now = func() time.Time {
		return time.Date(2020, 01, 10, 20, 10, 20, 1, time.UTC)
	}
	core.SetStdout(w)
	runMain()
	w.String()
	assert.Equal(t, errorOutput, w.String())
	os.Setenv("INPUT_MILLISECONDS", "10")

	w = bytes.NewBuffer(nil)
	core.SetStdout(w)
	runMain()
	w.String()
	assert.Equal(t, successOutput, w.String())
}
