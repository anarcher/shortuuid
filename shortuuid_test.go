package shortuuid

import (
	"testing"
)

func Test_Alphabet(t *testing.T) {
	alphabet := "01"
	su1 := New("", alphabet)
	su2 := NewDefault()

	if su1.alphabet.String() != alphabet {
		t.Errorf("uuid.alphabet %v is not equal to %v", su1.alphabet, alphabet)
	}

	alphabet2 := "01010101010101"
	su1.SetAlphabet(alphabet2)
	if su1.alphabet.String() != alphabet {
		t.Errorf("uuid.alphabet %v is not equal to %v", su1.alphabet, alphabet2)
	}

	if su1.alphabet.Len() != len(alphabet) {
		t.Errorf("su1.alphabet.Len: %v", su1.alphabet.Len())
	}

	t.Log(su2)

}
