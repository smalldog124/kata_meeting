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

func Test_Split_Input_Fred_Corwill_Should_Be_Fred_Coreill(t *testing.T) {
	expected := "(CORWILL, ALFRED)(CORWILL, FRED)(CORWILL, RAPHAEL)(CORWILL, WILFRED)(TORNBULL, BARNEY)(TORNBULL, BETTY)(TORNBULL, BJON)"
	text := "Fred:Corwill;Wilfred:Corwill;Barney:Tornbull;Betty:Tornbull;Bjon:Tornbull;Raphael:Corwill;Alfred:Corwill"

	actual := meeting.GetMeeting(text)

	assert.Equal(t, expected, actual)
}
func Test_Split_Input_Fred_Corwill_Should_Be_Fred_Coreill_2(t *testing.T) {
	expected := "(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)"
	text := "John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell"

	actual := meeting.GetMeeting(text)

	assert.Equal(t, expected, actual)
}

func Test_Fomating_Input_List_Name_Shoud_Be_Formater(t *testing.T) {
	expected := []string{"(Corwill, Fred)", "(Corwill, Wilfred)", "(Tornbull, Barney)", "(Tornbull, Betty)", "(Tornbull, Bjon)", "(Corwill, Raphael)", "(Corwill, Alfred)"}
	text := []string{"Fred:Corwill", "Wilfred:Corwill", "Barney:Tornbull", "Betty:Tornbull", "Bjon:Tornbull", "Raphael:Corwill", "Alfred:Corwill"}

	actual := meeting.Formating(text)

	assert.Equal(t, expected, actual)
}
