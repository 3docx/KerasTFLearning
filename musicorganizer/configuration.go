
package musicorganizer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hjson/hjson-go"
)

type Configuration struct {
	MusicIn, MusicOut            string
	UnorganizedFiles             string
	Preview, Move, DeleteMusicIn bool
	TreeTemplate, Replacement    string
}

var config Configuration

func init() {
	config = Configuration{
		MusicIn:          ".",
		MusicOut:         "organizedMusicLibrary",
		UnorganizedFiles: "unorganizedFiles",
		Preview:          false,
		Move:             false,
		DeleteMusicIn:    false,
		Replacement:      "_",
		TreeTemplate: `
		{{if .Genre}}
			{{.Genre}}
		{{else}}
			Unknonw genre
		{{end}}
		/
		{{if .AlbumArtist}}
			{{.AlbumArtist}}
		{{else if .Artist}}
			{{.Artist}}
		{{else}}
			Unknonw artist
		{{end}}
		/
		{{if .Album}}
			{{.Album}}
		{{else}}
			Unknonw album
		{{end}}
		/
		{{if gt .DiscTotal 1}}
			{{.Disc | printf "%02d"}}-
		{{end}}
		{{if .Track}}
			{{.Track | printf "%02d"}} {{/**/}}
		{{end}}
		{{if .Title}}
			{{.Title}}{{.Ext}}
		{{else}}
			{{.OriginalFilename}}
		{{end}}
		`,
	}
}

func LoadConfigurationFromFile(path string) error {
	configString, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	hjson.Unmarshal(configString, &data)

	b, _ := json.Marshal(data)
	json.Unmarshal(b, &config)

	return nil
}

func GetConfig() *Configuration {
	return &config
}

func PrintDefaultConfiguration() {
	fmt.Println(`{
	musicIn: .
	musicOut: organizedMusicLibrary
	unorganizedFiles: unorganizedFiles,
	preview: false
	move: false
	deleteMusicIn: false
	replacement: "_"
	treeTemplate:
		'''
		{{if .Genre}}
			{{.Genre}}
		{{else}}
			Unknonw genre
		{{end}}
		/
		{{if .AlbumArtist}}
			{{.AlbumArtist}}
		{{else if .Artist}}
			{{.Artist}}
		{{else}}
			Unknonw artist
		{{end}}
		/
		{{if .Album}}
			{{.Album}}
		{{else}}
			Unknonw album
		{{end}}
		/
		{{if gt .DiscTotal 1}}
			{{.Disc | printf "%02d"}}-
		{{end}}
		{{if .Track}}
			{{.Track | printf "%02d"}} {{/**/}}
		{{end}}
		{{if .Title}}
			{{.Title}}{{.Ext}}
		{{else}}
			{{.OriginalFilename}}
		{{end}}
		'''
}`)
}