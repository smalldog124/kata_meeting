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

func Test_Fomating_Input_List_Name_Shoud_Be_Formater(t *testing.T) {
	expected := []string{"(Corwill, Fred)", "(Corwill, Wilfred)", "(Tornbull, Barney)", "(Tornbull, Betty)", "(Tornbull, Bjon)", "(Corwill, Raphael)", "(Corwill, Alfred)"}
	text := []string{"Fred:Corwill", "Wilfred:Corwill", "Barney:Tornbull", "Betty:Tornbull", "Bjon:Tornbull", "Raphael:Corwill", "Alfred:Corwill"}

	actual := meeting.Formating(text)

	assert.Equal(t, expected, actual)
}
