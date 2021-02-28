package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	prgName      = "ignore"
	prgVersion   = "0.1.0"
	prgRepo      = "https://github.com/ajunior/ignore"
	githubApiUrl = "https://api.github.com/repos/github/gitignore/contents/"
)

type Template struct {
	Name string `json:"name"`
	Url  string `json:"download_url"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isValidPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func getTemplates() []Template {
	var templates []Template

	response, err := http.Get(githubApiUrl)
	checkError(err)
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkError(err)

	err = json.Unmarshal(content, &templates)
	checkError(err)
	return templates
}

func getListOfTemplates() []string {
	templates := getTemplates()

	var list []string

	for _, template := range templates {
		if !strings.Contains(template.Name, ".gitignore") {
			continue
		}

		s := strings.Split(template.Name, ".")[0]
		list = append(list, s)
	}

	return list
}

func createGitIgnoreFile(dir string, tpls set) {
	templates := getTemplates()

	p := path.Join(filepath.ToSlash(dir), ".gitignore")
	f, err := os.Create(p)
	checkError(err)
	defer f.Close()

	for _, template := range templates {
		t := strings.Split(template.Name, ".")[0]
		t = strings.ToLower(t)

		if tpls.has(t) {
			response, err := http.Get(template.Url)
			checkError(err)

			content, err := ioutil.ReadAll(response.Body)
			checkError(err)

			w := bufio.NewWriter(f)
			_, err = w.WriteString(string(content))
			checkError(err)
			w.Flush()

			response.Body.Close()
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		printUsageMessage()
		os.Exit(0)
	}

	switch command := strings.ToLower(os.Args[1]); command {
	case "-v", "--version":
		printVersionMessage()

	case "-h", "--help":
		printHelpMessage()

	case "-t", "--templates":
		fmt.Printf("Available .gitignore templates:\n")

		languages := getListOfTemplates()
		for i, language := range languages {
			if i%4 == 0 {
				fmt.Println()
			}
			fmt.Printf("\t%-30s", strings.ToLower(language))
		}
		fmt.Printf("\n\n")

	case "-c", "--create":
		if len(os.Args) < 4 {
			printUsageMessage()
			os.Exit(0)
		}

		dir := os.Args[2]
		if !isValidPath(dir) {
			fmt.Printf("%s: the path is invalid\n", prgName)
			printUsageMessage()
			os.Exit(0)
		}

		templates := os.Args[3:]
		if len(templates) == 0 {
			fmt.Println("error: you have to pass at least one template.")
			printUsageMessage()
			os.Exit(0)
		}

		templateSet := set{}
		for _, template := range templates {
			template := strings.ToLower(template)
			fmt.Println(template)
			templateSet.add(template)
		}

		createGitIgnoreFile(dir, templateSet)

	default:
		fmt.Printf("error: you entered a command that doesn't exist.\n\n")
		printUsageMessage()
	}
}
