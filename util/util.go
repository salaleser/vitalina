package util

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// Config is a configuration
var Config = make(map[string]string)

// Regexp patterns
const (
	ASAppIDPattern   = "^\\d{9,10}$"
	GPAppIDPattern   = "^[a-zA-Z][a-zA-Z0-9_]*(\\.[a-zA-Z0-9_]+)+[0-9a-zA-Z_]$"
	TimestampPattern = "[MWN]" // FIXME
)

// Store
const (
	NA = iota
	GooglePlay
	AppStore
)

const separator = "=" // разделитель для парсинга файлов

// ReadConfig reads lines from config file into the Config map
func ReadConfig() {
	file, err := os.Open("config")
	if err != nil {
		log.Printf("Не удалось найти конфигурационный файл (%s)", err)
		return
	}
	defer file.Close()

	var line []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = strings.Split(scanner.Text(), separator)
		if len(line) == 2 {
			key := line[0]
			value := line[1]
			Config[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// IsAppID returns 0 if appID is not an application ID
func IsAppID(s string) int {
	gpAppIDRegexp, _ := regexp.Compile(GPAppIDPattern)
	asAppIDRegexp, _ := regexp.Compile(ASAppIDPattern)

	if gpAppIDRegexp.MatchString(s) {
		return GooglePlay
	}

	if asAppIDRegexp.MatchString(s) {
		return AppStore
	}

	return NA
}

func IsLocation(s string) bool {
	switch s {
	case "ru":
	case "us":
	case "au":
	case "fr":
	case "de":
	case "no":
	case "pl":
	case "gb":
	case "es":
	case "pt":
	case "ca":
	case "br":
		return true
	}

	return false
}

func IsTimestamp(s string) bool {
	timestampRegexp, _ := regexp.Compile(TimestampPattern)

	return timestampRegexp.MatchString(s)
}
