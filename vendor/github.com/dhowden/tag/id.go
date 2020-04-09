package tag

import (
	"fmt"
	"io"
)

// Identify identifies the format and file type of the data in the ReadSeeker.
func Identify(r io.ReadSeeker) (format Format, fileType FileType, err error) {
	b, err := readBytes(r, 11)
	if err != nil {
		return
	}

	_, err = r.Seek(-11, io.SeekCurrent)
	if err != nil {
		err = fmt.Errorf("could not seek back to original position: %v", err)
		return
	}

	switch {
	case string(b[0:4]) == "fLaC":
		return VORBIS, FLAC, nil

	case string(b[0:4]) == "OggS":
		return VORBIS, OGG, nil

	case string(b[4:8]) == "ftyp":
		b = b[8:11]
		fileType = UnknownFileType
		switch string(b) {
		case "M4A":
			fileType = M4A

		case "M4B":
			fileType = M4B

		case "M4P":
			fileType = M4P
		}
		return MP4, fileType, nil

	case string(b[0:3]) == "ID3":
		b := b[3:]
		switch ui