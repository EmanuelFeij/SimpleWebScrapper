package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/gocolly/colly/v2"
)

type Name struct {
	name    string
	meaning string
}

func (n Name) String() string {
	return fmt.Sprintf("------%v------\n -> %v", n.name, strings.Split(n.meaning, "\n")[0])
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFromShell() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")

	text, err := reader.ReadString('\n')
	checkError(err)

	return strings.ReplaceAll(text, "\n", "")
}

func main() {
	name := readFromShell()
	var nameMeaning Name

	c := colly.NewCollector()

	c.OnHTML(`div[class=nameitem]`, func(e *colly.HTMLElement) {
		str := e.ChildText("p")
		if str != "" {
			nameMeaning = Name{
				name:    name,
				meaning: e.ChildText("p"),
			}
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	err := c.Visit("https://www.babynames.com/name/" + name)
	checkError(err)
	fmt.Println(nameMeaning)

}
