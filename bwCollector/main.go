package main

// In this file I will write the code to create
// - Read a yaml configuration to find the logs folders with the regex to scan
// - Scan the directionaries ( I will actually try to explore different ways of doing that)
// - Start a thread to read the logs files
// - Implement a strategy to make sure I dont read what has already been read.

// Day 4: Today I am working getting data from the config file.
// the config file should have data on which files to take. So I will learn how to read a file and parse the information
import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type logsPath struct {
	LogPath []string `yaml:"log"`
}

func (l *logsPath) LoadConfig(configFile string) {
	log.Println("Loading configuration file")
	file, err := os.ReadFile(configFile)

	if err != nil {
		log.Panic(err)
	}

	err = yaml.Unmarshal(file, &l)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(*&l.LogPath)

}

func main() {

	var l logsPath
	l.LoadConfig("config.yaml")

}
