package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getOmegaStatsForPlayer(player string) (*http.Response, error) {
	return http.Get(fmt.Sprintf("https://stats.omegastrikers.gg/player/%s", player))
}

func getOmegaPage(player string) error {
	response, err := getOmegaStatsForPlayer(player)
	if err != nil {
		return err
	}
	reader := response.Body

	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	foundID := ""

	document.Find("td[style*='background: #e29595'], td[style*='background: #496599'] > a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		for _, n := range s.Nodes {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					id, _ := strings.CutPrefix(attr.Val, "/get_username/")
					foundID = id
					return false
				}
			}
		}

		return true
	})

	if foundID == "" {
		return fmt.Errorf("Couldn't find player id for '%s'", player)
	}

	fmt.Printf("%s\n", foundID)

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var names []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		names = append(names, line)
	}
	for _, playerName := range names {
		err := getOmegaPage(strings.TrimSpace(playerName))
		if err != nil {
			fmt.Printf("Unknown\n")
		}
	}
}
