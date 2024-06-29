package matchprefixes

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"dkgosql.com/match-prefix-service/configs"
	"dkgosql.com/match-prefix-service/internal/consts"
)

// PrefixFinder interface

//go:generate mockgen --destination=../../mocks/matchprefix.go --source=matchprefix.go -package=mocks

type PrefixFinder interface {
	FindPrefix(input string) error
	// FindDummy(input string) error
}

// Prefix struct
type Prefix struct {
	PrefixList []string
}

// FindPrefix
func (p *Prefix) FindPrefix(input string) error {
	prefixData := configs.Data
	if len(prefixData) == 0 {
		return errors.New(consts.ERROR_EMPTY_DATA)
	}

	prefixes := bytes.Split(prefixData, []byte{'\n'})
	var wg sync.WaitGroup
	var prefixesMatchedList []string
	var resultPrefixesMatched []string
	var prefixesMaxLength int = 0
	prefixesLength := len(prefixes)

	if prefixesLength == 0 {
		return errors.New(consts.ERROR_DATA_NOT_FOUND)
	}

	// Find Remainder
	remainder := prefixesLength % consts.WORKER
	pages := prefixesLength / consts.WORKER
	skip, take := 0, pages // skip and take
	if consts.WORKER == 1 || (consts.WORKER > 0 && pages == 0) {
		wg.Add(1)
		go findMatchedPrefix(&wg, 1, input, &prefixes, &prefixesMatchedList, &prefixesMaxLength)

	} else if pages >= consts.WORKER {
		for counter := 0; counter < consts.WORKER; counter++ {
			wg.Add(1)
			data_list := prefixes[skip:take]
			skip = skip + pages
			take = take + pages
			go findMatchedPrefix(&wg, counter, input, &data_list, &prefixesMatchedList, &prefixesMaxLength)
		}

		if remainder > 0 {
			wg.Add(1)
			skip = pages * consts.WORKER
			take = pages*consts.WORKER + remainder
			data_list := prefixes[skip:take]
			go findMatchedPrefix(&wg, consts.WORKER, input, &data_list, &prefixesMatchedList, &prefixesMaxLength)
		}
	} else {
		err := fmt.Sprintf("Number of worker and pages does not matches: worker %v, pages %v\n", consts.WORKER, pages)
		log.Println(err)
		return errors.New(err)
	}

	wg.Wait()

	findLongestMatches(&resultPrefixesMatched, &prefixesMatchedList, &prefixesMaxLength)
	if len(resultPrefixesMatched) == 0 {
		return errors.New(consts.ERROR_DATA_NOT_FOUND)
	}
	p.PrefixList = resultPrefixesMatched
	return nil
}

// findMatchedPrefix
func findMatchedPrefix(wg *sync.WaitGroup, counter int, searchStr string, prefixes *[][]byte, prefixesMatchedList *[]string, prefixesMaxLength *int) {
	for _, prefix := range *prefixes {
		// fmt.Println("all: ",string(prefix))
		prefixString := string(prefix) // fmt.Println("prefix:", string(prefix))
		if strings.HasPrefix(prefixString, searchStr) && len(prefixString) > 0 {
			*prefixesMatchedList = append(*prefixesMatchedList, prefixString)
			prefixLength := len(prefixString)
			if prefixLength > *prefixesMaxLength {
				*prefixesMaxLength = prefixLength
			}
		}
	}
	wg.Done()
}

// findLongestMatches
func findLongestMatches(resultPrefixesMatched *[]string, prefixesMatchedList *[]string, prefixesMaxLength *int) {
	for _, prefix := range *prefixesMatchedList {
		if len(prefix) == *prefixesMaxLength {
			*resultPrefixesMatched = append(*resultPrefixesMatched, prefix)
		}
	}
}

func (p *Prefix) FindDummy(input string) error {
	return nil
}
