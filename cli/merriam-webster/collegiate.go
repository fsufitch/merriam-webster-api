package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	a "github.com/logrusorgru/aurora"

	mwapi "github.com/fsufitch/merriam-webster-api"
	"github.com/fsufitch/merriam-webster-api/types"
)

func collegiateLookup(terms []string, url string, key string, json bool, verbose bool, debugf printfFunc) error {
	debugf("looking up %v in collegiate", terms)
	client := mwapi.NewClient(key, userAgent, &mwapi.BaseURLs{Collegiate: url})
	client.SetDebugf(debugf)
	results, suggestions, err := client.SearchCollegiate(strings.Join(terms, " "))
	if err != nil {
		debugf("error during lookup: %v", err)
		return err
	}

	if len(results) > 0 {
		debugf("found %d results; formatting...", len(results))
		return collegiateFormatResults(results, json, verbose)
	} else if len(suggestions) > 0 {
		debugf("found no results, but have %d suggestions; formatting...", len(suggestions))
		return collegiateFormatSuggestions(suggestions, json)
	}

	debugf("no results, no suggestions, yet no error... wtf?")
	return errors.New("no results, suggestions, or error")
}

type collegiateResults struct {
	Raw []types.CollegiateResult `json:"raw,omitempty"`
}

func collegiateFormatResults(results []types.CollegiateResult, toJSON bool, verbose bool) error {
	if toJSON && verbose {
		cr := collegiateResults{Raw: results}
		bytes, err := json.Marshal(cr)
		if err != nil {
			return err
		}
		bytes = append(bytes, '\n')
		_, err = os.Stdout.Write(bytes)
		return err
	}

	for i, entry := range results {
		hw := strings.ReplaceAll(entry.HeadwordInfo.Headword, "*", "")
		fmt.Printf("%s %s", a.Underline(a.Bold(hw)), entry.Function)
		if len(results) > 1 {
			fmt.Printf(" (%d of %d)", i+1, len(results))
		}
		fmt.Println()

		prsStrings := []string{}
		prsSep := "; "
		for _, prs := range entry.HeadwordInfo.Pronounciations {
			prsStrings = append(prsStrings, fmt.Sprintf("%s %s %s", prs.LabelBefore, prs.MerriamWebsterFormat, prs.LabelAfter))
		}
		if len(prsStrings) > 0 {
			fmt.Printf("\\%s\\\n", strings.Join(prsStrings, prsSep))
		}

		for j, def := range entry.Definitions {
			if len(entry.Definitions) > 1 {
				fmt.Printf("%d ", j+1)
			}

			senses, err := def.SenseSequence.Contents()
			if err != nil {
				return err
			}

			for _, sense := range senses {
				fmt.Printf("Sense: %+v\n", sense)
				if sense.SubSequence != nil {
					subSenses, err := sense.SubSequence.Contents()
					if err != nil {
						return err
					}
					for _, subsense := range subSenses {
						fmt.Printf(" - Subsense: %+v\n", subsense)
						if subsense.Sense != nil {
							fmt.Printf("   - sense DT: %+v\n", subsense.Sense.DefiningText)
						}
					}
				}
				// fmt.Printf("%+v\n", sense.Sense.DefiningText)
			}
		}

	}
	return nil
}

type collegiateSuggestions struct {
	Suggestions []string `json:"suggestions"`
}

func collegiateFormatSuggestions(suggestions []string, toJSON bool) error {
	cs := collegiateSuggestions{Suggestions: suggestions}
	if toJSON {
		bytes, err := json.Marshal(cs)
		if err != nil {
			return err
		}
		bytes = append(bytes, '\n')
		_, err = os.Stdout.Write(bytes)
		return err
	}

	fmt.Println("No results found. Did you mean one of these?")
	for _, s := range suggestions {
		fmt.Printf("- %s\n", s)
	}
	return nil
}
