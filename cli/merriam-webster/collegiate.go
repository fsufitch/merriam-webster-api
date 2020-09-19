package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

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

	return errors.New("Not implemented")
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
		fmt.Printf("- %s", s)
	}
	return nil
}
