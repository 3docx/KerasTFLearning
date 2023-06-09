
package main

import (
	"flag"
	"fmt"

	"github.com/JosephCaillet/boris/log"
	"github.com/JosephCaillet/boris/musicorganizer"
)

func main() {
	var configFile string
	var showDefaultConfig bool
	conf := musicorganizer.GetConfig()

	flag.Usage = borisUsage

	flag.StringVar(&configFile, "c", "", "config file")
	flag.BoolVar(&showDefaultConfig, "s", false, "show default config, in hjson config file format")

	flag.StringVar(&conf.MusicIn, "i", conf.MusicIn, "input music directory")
	flag.StringVar(&conf.MusicOut, "o", conf.MusicOut, "output music directory")
	flag.StringVar(&conf.UnorganizedFiles, "u", conf.UnorganizedFiles, "unorganized files directory where folders will be moved if they contains files that does not have any tags")

	flag.BoolVar(&conf.Preview, "p", false, "preview change")
	flag.BoolVar(&conf.Move, "m", false, "move file instead of copying them")
	flag.BoolVar(&conf.DeleteMusicIn, "d", false, "delete input music folder after organizing")

	flag.StringVar(&conf.Replacement, "r", conf.Replacement, "replacement for <>:\"\\/|?*")
	flag.StringVar(&conf.TreeTemplate, "t", conf.TreeTemplate, "organizing template")

	flag.Parse()

	if showDefaultConfig {
		musicorganizer.PrintDefaultConfiguration()
		return
	}

	if configFile != "" {
		if err := musicorganizer.LoadConfigurationFromFile(configFile); err != nil {
			log.Error("failure loading configuration", err)
			return
		}
		flag.Parse() //needed to overide config file by command line
	}

	if err := musicorganizer.Reorganize(); err != nil {
		log.Error("failure reorganizing files", err)
	}
}

func borisUsage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: boris [options]\n")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
	fmt.Print(`
About boris:
  boris is a tool to reorganize automatically your music library, using tags defined in your music files and templating.
  boris is made by Joseph Caillet and is released under GPL-3.0.
  View source code, report bug, contribute here: https://github.com/JosephCaillet/boris.
`)
}