package processors

import (
	"crypto/md5"
	"errors"
)

var MD5 = &md5Container{}

type md5Container struct{}

// Pre for the Timestamp processor attaches a timestamp at the beginning of the data
func (md5Container) Pre(d *[]byte) (err error) {
	hash := md5.Sum(*d)
	*d = append(hash[:], *d...)
	return
}

// Post for the Timestamp processor doesn't do anything
func (md5Container) Post(d *[]byte) (err error) {
	var hash [md5.Size]byte
	d2 := *d

	copy(hash[:], d2[:md5.Size])

	*d = d2[md5.Size:]

	if hash != md5.Sum(*d) {
		err = errors.New("MD5 Mismatch")
	}

	return
}
