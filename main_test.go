package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, 1, 1, "sollten aber gleich sein")
}

func TestMessageLexer(t *testing.T) {

	input := []string{"hello", "world", "+tag01", "+tag02", "proj:project01", "ref:externalReference", "!"}

	interval := &Interval{}
	lexInterval(input, interval)

	assert.Equal(t, "hello world !", interval.Annotation)
	assert.Equal(t, []string{"tag01", "tag02"}, interval.Tags)
	assert.Equal(t, "project01", interval.Project)
	assert.Equal(t, "externalReference", interval.Ref)
	assert.Equal(t, strings.Join(input, " "), interval.Raw)

}
