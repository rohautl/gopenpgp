package internal

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/rohautl/go-crypto/openpgp/armor"
)

// Unarmor unarmors an armored string.
func Unarmor(input string) (*armor.Block, error) {
	io := strings.NewReader(input)
	b, err := armor.Decode(io)
	if err != nil {
		return nil, errors.Wrap(err, "gopenpgp: unable to unarmor")
	}
	return b, nil
}
