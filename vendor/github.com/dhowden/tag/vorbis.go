// Copyright 2015, David Howden
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tag

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func newMetadataVorbis() *metadataVorbis {
	return &metadataVorbis{
		c: make(map[string]string),
	}
}

type metadataVorbis struct {
	c map[string]string // the vorbis comments
	p *Picture
}

func (m *metadataVorbis) readVorbisComment(r io.Reader) error {
	vendorLen, err := readInt32LittleEndian(r)
	if err != nil {
		return err
	}

	if vendorLen < 0 {
		return fmt.Errorf("invalid encoding: expected positive length, got %d", vendorLen)
	}

	vendor, err := readString(r, vendorLen)
	if err != nil {
		return err
	}
	m.c["vendor"] = vendor

	commentsLen, err := readInt32LittleEndian(r)
	if err != nil {
		return err
	}

	for i := 0; i < commentsLen; i++ {
		l, err := readInt32LittleEndian(r)
		if err != nil {
			return err
		}
		s, err := readString(r, l)
		if err != nil {
			return err
		}
		k, v, err := parseComment(s)
		if err != nil {
			return err
		}
		m.c[strings.ToLower(k)] = v
	}
	return nil
}

func (m *metadataVorbis) readPictureBlock(r io.Reader) error {
	b, err := readInt(r, 4)
	if err != nil {
		return err
	}
	pictureType, ok := pictureTypes[byte(b)]
	if !ok {
		return fmt.Errorf("invalid picture type: %v", b)
	}
	mimeLen, err := readInt(r, 4)
	if err != nil {
		return err
	}
	mime, err := readString(r, mimeLen)
	if err != nil {
		return err
	}

	ext := ""
	switch mime {
	case "image/jpeg":
		ext = "jpg"
	case "image/png":
		ext = "png"
	case "image/gif":
		ext = "gif"
	}

	descLen, err := readInt(r, 4)
	if err != nil {
		return err
	}
	desc, err := readString(r, descLen)
	if err != nil {
		return err
	}

	// We skip width <32>, height <32>, colorDepth <32>, coloresUsed <32>
	_, err = readInt(r, 4) // width
	if err != nil {
		return err
	}
	_, err = readInt(r, 4) // height
	if err != nil {
		return err
	}
	_, err = readInt(r, 4) // color depth
	if err != nil {
		return err
	