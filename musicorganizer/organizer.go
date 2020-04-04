package musicorganizer

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/JosephCaillet/boris/fs"
	"github.com/JosephCaillet/boris/log"
	"github.com/cleversoap/go-cp"
	"github.com/dhowden/tag"
)

type readTagsError struct {
	error
	filepath string
}

func (e readTagsError) Error() string {
	return fmt.Sprintf("opening tags of file %q: %v", e.filepath, e.error)
}

var tmpl *template.Template

func Reorganize() error {
	var err error
	tmpl, err = template.New("TreeTemplate").Parse(config.TreeTemplate)
	if err != nil {
		return fmt.Errorf("parsing template: %v", err)
	}

	if config.Preview {
		log.Mode("Preview mode")
	} else {
		if config.Move {
			log.Mode("Move mode")
		} else {
			log.Mode("Copy mode")
		}
	}

	fmt.Print("\n🔎  Listing files...\n\n")
	exploredPathes, err := fs.Tree(config.MusicIn)
	if err != nil {
		return err
	}

	err = reorganizeFiles(&exploredPathes)
	if err != nil {
		return fmt.Errorf("error reorganizing file: %v", err)
	}

	if !config.Preview && config.DeleteMusicIn {
		if err = os.RemoveAll(config.MusicIn); err != nil {
			return fmt.Errorf("delete input music directory: %v", err)
		}

	}

	return nil
}

func computeNewFilePath(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("opening file %q: %v", filePath, err)
	}
	defer file.Close()

	tags, err := tag.ReadFrom(file)
	if err != nil {
		return "", readTagsError{filepath: filePath, error: err}
	}

	var pathBuilder strings.Builder
	err = tmpl.Execute(&pathBuilder, newMetadata(tags, filePath))
	if err != nil {
		return "", fmt.Errorf("executing template: %v", err)
	}

	newPath := strings.Replace(pathBuilder.String(), "\n", "", -1)
	newPath = strings.Replace(newPath, "\t", "", -1)
	return path.Clean(config.MusicOut + "/" + newPath), nil
}

func reorganizeFiles(exploredPathes *[]fs.FilePathInfos) error {
	nonMusicFile := make([]string, 0)
	lastDestinationDir := config.MusicOut
	musicFoundInDir := false
	log.StartOperation(