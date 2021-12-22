package netstring

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

type NetString []byte

const (
	defaultBufSize = 9 + 1 // we want to support max 100 MB large net strings - the "+ 1" is for the ":" character
)

func (n NetString) String() string {
	return string(n)
}

func FromReader(reader io.Reader) (NetString, error) {
	buffer := make([]byte, defaultBufSize)

	_, err := reader.Read(buffer)
	if err == io.EOF {
		return FromBytes(buffer)
	}

	if err != nil {
		return nil, err
	}

	chunks, err := GetChunks(buffer)

	if err != nil {
		return nil, err
	}

	left, right := chunks[0], chunks[1]

	length, err := strconv.Atoi(string(left))

	if err != nil {
		return nil, err
	}

	lengthRight := len(right)
	readme := length - lengthRight + 1 // "+ 1" is for the "," at the end

	rest := make([]byte, readme)

	n, err := io.ReadAtLeast(reader, rest, readme)
	if err != nil {
		return nil, err
	}
	if n != readme {
		return nil, errors.New("FromReader: io.ReadAtLeast returned n != readme")
	}
	return FromBytes(append(right, rest...))
}

func GetChunks(b []byte) ([][]byte, error) {
	chunks := bytes.SplitN(b, []byte{58}, 2)

	if len(chunks) != 2 {
		return nil, errors.New("FromBytes: len(chunks) != 2")
	}
	return chunks, nil
}

func FromBytes(b []byte) (NetString, error) {
	if len(b) == 0 {
		return make(NetString, 0), nil
	}
	if b[len(b)-1] != ',' {
		return nil, errors.New("FromBytes: netstring does not end with ','")
	}
	return b[:len(b)-1], nil
}
