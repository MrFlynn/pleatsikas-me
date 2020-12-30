package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	inFileName         string
	outFileName        string
	experienceFileName string
	projectsFileName   string
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
}

func (e *experience) UnmarshalJSON(data []byte) error {
	type alias experience
	aux := struct {
		Start string `json:"start"`
		End   string `json:"end"`
		*alias
	}{
		alias: (*alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	start, err := time.Parse("Jan 2006", aux.Start)
	if err != nil {
		return err
	}

	end, err := time.Parse("Jan 2006", aux.End)
	if err != nil {
		return err
	}

	e.Start = start
	e.End = end

	return nil
}

func (e experience) FormatDateRange() string {
	var (
		yearStart  = e.Start.Format("2006")
		monthStart = e.Start.Format("01")
		dayStart   = e.Start.Format("02")

		yearEnd  = e.End.Format("2006")
		monthEnd = e.End.Format("01")
		dayEnd   = e.End.Format("02")
	)

	return fmt.Sprintf(
		"\\hfill \\DatestampYMD{%s}{%s}{%s} -- \\DatestampYMD{%s}{%s}{%s}",
		yearStart, monthStart, dayStart,
		yearEnd, monthEnd, dayEnd,
	)
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
