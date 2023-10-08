package main

import (
	"fmt"
	"regexp"
)

// Regex work nice to match or find specific custom patterns in a string

func RegexMatch(r, s string) bool {
	re := regexp.MustCompile(r)
	match := re.MatchString(s)
	fmt.Printf("\"%s\" has %s -> %v\n", s, r, match)
	return match
}

func RegexFind(r, s string) {
	re := regexp.MustCompile(r)
	found := re.FindAllString(s, -1)
	fmt.Printf("Finds %s in \"%s\" -> %v\n", r, s, found)
}

func main() {

	r := "s"
	re := regexp.MustCompile(r)
	// Match requires bytes, MatchString requires string
	s := "antiagos"
	matchRes := re.Match([]byte(s))
	fmt.Printf("\"%s\" has %s -> %v\n", s, r, matchRes)

	s = "kika"
	RegexMatch(r, s)

	RegexFind(r, s)

	s = "santi y kikas"
	RegexFind(r, s)

	// Regex for rune with runes at both sides
	r = ".a."
	s = "santiago y calypsa aman a kika"
	RegexFind(r, s)

	// Regex to match one or more @
	r = "@+"
	s = "hol@"
	RegexMatch(r, s)

	// Regex to match one or more digits
	r = "[0-9]+"
	s = "cl2a3ve"
	RegexMatch(r, s)

	// Regex to match one or more lowercase runes
	r = "[a-z]+"
	s = "23784"
	RegexMatch(r, s)

	// Regex to allow only 10 digits
	// double \\, first one as escape character
	r = "^\\d{10}$"
	s = "3103347776"
	RegexMatch(r, s)

	// Regex to my pass
	// regexp package does not support lookaheads so this could be an approach
	// strings.ContainsAny could work too
	var matchNums bool
	var matchChars bool
	var matchSpec bool
	var matchLen bool
	rNums := "[0-9]+"
	rChars := "[a-zA-Z]+"
	// Define special characters to be allowed
	rSpec := "[$#@&?!=%]+"
	// len(s) will work better
	rLen := "^[a-zA-Z0-9$#@&?!=]{12,}$"
	s = "AsffdQ#34!34dxx878"
	matchNums = RegexMatch(rNums, s)
	if matchNums {
		matchChars = RegexMatch(rChars, s)
	}
	if matchChars {
		matchSpec = RegexMatch(rSpec, s)
	}
	if matchSpec {
		matchLen = RegexMatch(rLen, s)
	}
	if matchLen {
		fmt.Printf("\"%s\" is valid password\n", s)
	} else {
		fmt.Printf("\"%s\" is invalid password\n", s)
	}

	// Regex to validate url
	r = "^(https?://(www\\.)?|www\\.)([a-z]{3,20})\\.([a-z]{3,5})\\.([a-z]{2,3})?$"
	s = "https://calypsa.com.co"
	RegexMatch(r, s)
	RegexFind(r, s)

	// Regex to validate email
	r = "^([a-z0-9\\._]{3,20})@([a-z0-9]{3,10}\\.([a-z0-9]{3,15})(\\.[a-z]{2-3})?)$"
	s = "jesucristo_garcia@miempresa.com"
	RegexMatch(r, s)
	RegexFind(r, s)

	// Regex to validate colombian celphone number
	r = "^(3(00|01|02|03|04|05|06|07|08|09|10|11|12|13|14|15|16|17|18|19|20|21|22|23|24))[0-9]{7}$"
	s = "3154678760"
	RegexMatch(r, s)
	RegexFind(r, s)

	// Regex to search .pdf files
	r = "^[a-zA-Z0-9-_]+\\.pdf$"
	files := []string{
		"file.exe",
		"file1_2",
		"myfile-final.txt",
		"file_for-tes-2023.pdf",
		".pdf",
		"a.pd",
		"a.pdf",
	}
	pdfs := []string{}
	for _, file := range files {
		if RegexMatch(r, file) {
			pdfs = append(pdfs, file)
		}
	}
	fmt.Printf("%d pdf files found: %v\n", len(pdfs), pdfs)

}
