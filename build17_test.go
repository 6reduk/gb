// +build go1.1
// +build !go1.8

package gb

import "go/build"

func nogoerr(dir string) error {
	return &build.NoGoError{Dir: dir}
}
