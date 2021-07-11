package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

const latexDateTemplate = `\DatestampYMD{2006}{01}{02}`

var (
	inFileName         string
	outFileName        string
	experienceFileName string
	projectsFileName   string

	escapeLatex = regexp.MustCompile(`\$`)
)

func init() {
	flag.StringVar(&inFileName, "if", "resume.tex.tmpl", "Input LaTeX template file")
	flag.StringVar(&outFileName, "of", "resume.tex", "Output LaTeX file")
	flag.StringVar(&experienceFileName, "ef", "experience.json", "JSON file containing work experience")
	flag.StringVar(&projectsFileName, "pf", "projects.json", "JSON file containing projects")

	flag.Parse()
}

type experience struct {
	Company struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"company"`

	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Position    string    `json:"position"`
	Description []string  `json:"description"`

	// If the position is still ongoing
	current bool
}

func (e *experience) UnmarshalJSON(data []byte) error {
	var err error

	type alias experience
	aux := struct {
		Start string `json:"start"`
		End   string `json:"end"`
		*alias
	}{
		alias: (*alias)(e),
	}

	err = json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	e.Start, err = time.Parse("Jan 2006", aux.Start)
	if err != nil {
		return err
	}

	if strings.ToLower(aux.End) == "present" {
		e.End = time.Now()
		e.current = true
	} else {
		e.End, err = time.Parse("Jan 2006", aux.End)

		if err != nil {
			return err
		}
	}

	for i, line := range e.Description {
		e.Description[i] = escapeLatex.ReplaceAllLiteralString(line, `\$`)
	}

	return nil
}

func (e experience) FormatDateRange() string {
	start := e.Start.Format(latexDateTemplate)
	end := "Present"

	if !e.current {
		end = e.End.Format(latexDateTemplate)
	}

	return `\hfill ` + start + " -- " + end
}

func parseExperience() ([]experience, error) {
	experienceFile, err := os.Open(experienceFileName)
	if err != nil {
		return nil, err
	}

	defer experienceFile.Close()

	var work []experience
	err = json.NewDecoder(experienceFile).Decode(&work)
	if err != nil {
		return nil, err
	}

	return work, nil
}

type project struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Language    string `json:"language"`
	Description string `json:"description"`
}

func (p *project) UnmarshalJSON(data []byte) error {
	type alias project
	aux := struct {
		Language []string `json:"language"`
		*alias
	}{
		alias: (*alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	p.Language = strings.Join(aux.Language, ", ")

	return nil
}

func parseProjects() ([]project, error) {
	projectFile, err := os.Open(projectsFileName)
	if err != nil {
		return nil, err
	}

	defer projectFile.Close()

	var projects []project
	err = json.NewDecoder(projectFile).Decode(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

type resume struct {
	Work     []experience
	Projects []project
}

func main() {
	work, err := parseExperience()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded experience JSON (%s)", experienceFileName)

	projects, err := parseProjects()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Loaded projects JSON (%s)", projectsFileName)

	tmpl, err := template.New(inFileName).ParseFiles(inFileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parsed resume template file (%s)", inFileName)

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer outFile.Close()

	if err := tmpl.Execute(outFile, resume{
		Work:     work,
		Projects: projects,
	}); err != nil {
		log.Fatal(err)
	}

	log.Printf("Sucessfully wrote out template to %s", outFileName)
}
