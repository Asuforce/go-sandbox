package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Asuforce/gogo/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			fmt.Printf("Failed to search %q's synonim: %v\n", word, err)
			return
		}
		if len(syns) == 0 {
			fmt.Printf("There is no synonim for %q\n ", word)
			return
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}

}
