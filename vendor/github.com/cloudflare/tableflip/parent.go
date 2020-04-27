package tableflip

import (
	"encoding/gob"
	"io"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

const (
	sentinelEnvVar = "TABLEFLIP_HAS_PARENT_7DIU3"
	notifyReady    = 42
)

type parent struct {
	wr     *os.File
	result <-chan error
	exited <-chan struct{}
}

func newParent(env *env) (*parent, map[fileName]*file, error) {
	if env.getenv(sentinelEnvVar) == "" {
		return nil, make(map[fileName]*file), nil
	}

	wr := env.newFile(3, "write")
	rd := env.newFile(4, "read")

	var names [][]string
	dec := gob.NewDecoder(rd)
	if err := dec.Decode(&names); err != nil {
		return nil, nil, errors.Wrap(err, "can't decode names from parent process")
	}

	files := make(map[fileName]*file)
	for i, parts := range names {
		var key fileName
		copy(key[:], parts)

		// Start at 5 to account for stdin, etc. and write
		// and read pipes.
		fd := 5 + i
		env.closeOnExec(fd)
		files[key] = &file{
			env.newFile(uintptr(fd), key.String()),
			uintptr(fd),
		}
	}

	result := make(chan error, 1)
	exited := make(chan struct{})
	go func() {
		defer rd.Close()

		n, err := io.Copy(ioutil.Discard, rd)
		if n != 0 {
			err = errors.New("unexpected data from parent process")
		} else if err != nil {
			err = errors.Wrap(err, "unexpected error while waiting for parent to exit")
		}
		result <- err
		close(exited)
	}()

	return &parent{
		wr:     wr,
		result: result,
		exited: exited,
	}, files, nil
}

func (ps *parent) sendReady() error {
	defer ps.wr.Close()
	if _, err := ps.wr.Write([]byte{notifyReady}); err != nil {
		return errors.Wrap(err, "can't notify parent process")
	}
	return nil
}
