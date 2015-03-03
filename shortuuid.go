package shortuuid

import (
	"github.com/satori/go.uuid"
	"math"
	"math/big"
	"strings"
)

const (
	DEFAULT_ALPHABET = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

type ShortUUID struct {
	name     string
	alphabet *StringSet
	uuid     uuid.UUID
}

func NewDefault() *ShortUUID {
	return New("", "")
}

func NewWithName(name string) *ShortUUID {
	return New(name, "")
}

func New(name, alphabet string) *ShortUUID {

	var _uuid uuid.UUID
	if name == "" {
		_uuid = uuid.NewV4()
	} else if strings.HasPrefix(name, "http") {
		_uuid = uuid.NewV5(uuid.NamespaceDNS, name)
	} else {
		_uuid = uuid.NewV5(uuid.NamespaceURL, name)
	}

	suuid := &ShortUUID{uuid: _uuid}

	if alphabet == "" {
		alphabet = DEFAULT_ALPHABET
	}
	suuid.SetAlphabet(alphabet)
	return suuid
}

func (s *ShortUUID) SetAlphabet(alphabet string) {
	set := NewStringSet()
	for _, a := range alphabet {
		set.Add(string(a))
	}
	s.alphabet = set
}

func (s ShortUUID) String() string {
	return ""
}

func (s *ShortUUID) encode() string {
	padLen := s.encodeLen(len(s.uuid.Bytes()))
	return s.numToString(1, padLen)
}

func (s *ShortUUID) encodeLen(numBytes int) int {
	factor := math.Log(float64(25)) / math.Log(float64(s.alphabet.Len()))
	length := math.Ceil(factor * float64(numBytes))
	return int(length)
}

//Covert a number to a string, using the given alphabet.
func (s *ShortUUID) numToString(number int, padToLen int) string {

	return ""
}

func uuidToInt(_uuid uuid.UUID) *big.Int {
	var i big.Int
	i.SetString(strings.Replace(_uuid.String(), "-", "", 4), 16)
	return &i
}
