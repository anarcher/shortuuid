package shortuuid

import (
	"github.com/satori/go.uuid"
	"testing"
)

func Test_Alphabet(t *testing.T) {
	alphabet := "01"
	su1 := NewWithAlphabet(alphabet)

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

}

func Test_Encode(t *testing.T) {
	id := "3b1f8b40-222c-4a6e-b77e-779d5a94e21c"
	sid := "bYRT25J5s7Bniqr4b58cXC"

	u, err := uuid.FromString(id)
	su := New()

	if err != nil {
		t.Error(err)
	}
	if su.Encode(u) != sid {
		t.Errorf("They should be equal. %v %v", sid, su.String())
	}
}

func Test_Decode(t *testing.T) {
	id := "3b1f8b40-222c-4a6e-b77e-779d5a94e21c"
	sid := "bYRT25J5s7Bniqr4b58cXC"

	su := New()

	u1, u1_err := su.Decode(sid)
	u2, u2_err := uuid.FromString(id)
	if u1_err != nil {
		t.Error(u1_err)
	}
	if u2_err != nil {
		t.Error(u2_err)
	}

	if u1.String() != u2.String() {
		t.Errorf("They should be equal. %v %v", u1.String(), u2.String())
	}
}
