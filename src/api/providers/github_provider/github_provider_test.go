package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc1234")
	assert.EqualValues(t, "token abc1234", header)
}
