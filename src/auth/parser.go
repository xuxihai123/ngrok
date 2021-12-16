package auth

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	Path    string
	Content string
	Tokens  map[string]string
}

func NewParser(path string) *Parser {
	fmt.Println(path)
	return &Parser{
		Path:   path,
		Tokens: make(map[string]string),
	}
}

func (this *Parser) Parse() error {
	file, err := os.Open(this.Path)
	if err != nil {
		err = fmt.Errorf("Failed to read configuration file %s: %v", this.Path, err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		match, _ := regexp.MatchString(`^\s*#`, text)
		if match {
			continue
		}

		match2, _ := regexp.MatchString(`^\s*[^\s:]+:[^s]+`, text)
		if match2 == false {
			continue
		}
		parts := strings.Split(text, ":")
		this.Tokens[parts[0]] = parts[1]
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
