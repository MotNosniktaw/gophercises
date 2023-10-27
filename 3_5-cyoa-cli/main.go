package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/exp/slices"
)

func main() {
	j, _ := os.Open("gopher.json")
	b, _ := io.ReadAll(j)

	m := make(map[string]struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		}
	})
	json.Unmarshal(b, &m)

	arc := "intro"

	for {
		next := m[arc]

		fmt.Println(next.Title)
		fmt.Println("--------------------------")
		fmt.Println(next.Story)
		fmt.Println()
		if len(next.Options) > 0 {
			choices := []int{}

			for i, o := range next.Options {
				fmt.Println()
				fmt.Printf("%d. %s", i, o.Text)
				fmt.Println()
				choices = append(choices, i)
			}

			for {
				// very buggy but i did what i came here to do :D
				fmt.Println("which one?")
				var x int
				_, err := fmt.Scanf("%d", &x)
				if err == nil && slices.Contains(choices, x) {
					arc = next.Options[x].Arc
					break
				}
			}
		} else {
			fmt.Println("You have completed your adventure")
			break
		}
	}
}
