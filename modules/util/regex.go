package util

import "regexp"

//GetRegexParams ...
func GetRegexParams(compRegEx *regexp.Regexp, url string) (params map[string]string) {

	match := compRegEx.FindStringSubmatch(url)

	params = map[string]string{}
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[name] = match[i]
		}
	}
	return
}
