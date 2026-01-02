package utils

import (
	"regexp"
	"strings"
	"time"
)

// Clean whitespace and line breaks
func TrimLineBreak(text string) string {
	runes := []rune(text)
	for idx, ch := range runes { // character normalization
		if ch == '\v' || ch == '\f' || ch == '\r' {
			runes[idx] = '\n'
		}
	}
	text = string(runes)

	multiNewlinesRegex := regexp.MustCompile(`\n\n\n+`)
	text = multiNewlinesRegex.ReplaceAllString(text, "\n\n")

	// normalize unicode dashes matches
	text = strings.ReplaceAll(text, "−", "-") // U+2212
	text = strings.ReplaceAll(text, "–", "-") // U+2013
	return text
}

// Step 2: Replace airport codes with names
func ConvertNames(text, filepath string) (string, error) {
	iata_code := regexp.MustCompile(`\*?#[A-Z]{3}`)
	icao_code := regexp.MustCompile(`\*?##[A-Z]{4}`)

	icaoFound := icao_code.FindAllString(text, -1)
	var err error
	text, err = checkTags(text, filepath, "icao_code", icaoFound)
	if err != nil {
		return "", err
	}

	iataFound := iata_code.FindAllString(text, -1)
	text, err = checkTags(text, filepath, "iata_code", iataFound)
	if err != nil {
		return "", err
	}

	return text, nil
}

func checkTags(text, filepath, tagtype string, matches []string) (string, error) {
	for _, tag := range matches {
		find := "name"
		// determine whether requester asked for city (municipality) or airport name
		tagTrimmed, hasStar := strings.CutPrefix(tag, "*")
		if hasStar {
			find = "municipality"
		}
		// remove prefix # or ##
		code := strings.TrimPrefix(tagTrimmed, "##")
		code = strings.TrimPrefix(code, "#")
		code = strings.TrimSpace(code)
		if code == "" {
			// nothing to lookup; skip
			continue
		}
		airportName, found, err := lookupInCSV(tagtype, code, find, filepath)

		// propagate CSV errors
		if err != nil {
			return "", err
		}
		// If not found, skip this tag and leave the original text unchanged
		if !found || airportName == "" {
			continue
		}
		// replace found tag
		text = strings.ReplaceAll(text, tag, HighlightDest+airportName+Reset)
	}
	return text, nil
}

// Step 3: Format timestamps
func FormatTime(input string) string {
	reDate := regexp.MustCompile(`D\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[-+][0-9]{2}:[0-9]{2})\)`)
	reT12 := regexp.MustCompile(`T12\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[-+][0-9]{2}:[0-9]{2})\)`)
	reT24 := regexp.MustCompile(`T24\([0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}(Z|[-+][0-9]{2}:[0-9]{2})\)`)

	dateLayout := "02 Jan 2006"
	t12Layout := "03:04PM (-07:00)"
	t24Layout := "15:04 (-07:00)"

	input = findAndReplace(input, dateLayout, "D", HighlightDate, reDate)
	input = findAndReplace(input, t12Layout, "T12", HighlightTime, reT12)
	input = findAndReplace(input, t24Layout, "T24", HighlightTime, reT24)
	return input
}

func stampToTime(input, stamp string) time.Time {
	input = strings.TrimPrefix(input, stamp)
	input = strings.Trim(input, "()")
	input = strings.ReplaceAll(input, "Z", "+00:00")
	layout := "2006-01-02T15:04-07:00"
	t, _ := time.Parse(layout, input)
	return t
}

func findAndReplace(input, layout, stamp, format string, re *regexp.Regexp) string {
	for {
		datetime := re.FindString(input)
		if datetime == "" {
			break
		}
		t := stampToTime(datetime, stamp).Format(layout)
		if stamp == "T12" || stamp == "T24" {
			t = strings.Replace(t, "(", HighlightOffset+"(", 1)
		}
		input = strings.ReplaceAll(input, datetime, format+t+Reset)
	}
	return input
}
