package meeting_test

import (
	meeting "kata"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpperCase_Input_Small_Should_Be_SMALL(t *testing.T) {
	expected := "SMALL"
	text := "Small"

	actual := meeting.UpperCase(text)

	assert.Equal(t, expected, actual)
}
