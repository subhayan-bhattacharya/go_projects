package main

import (
	"fmt"
	"htmlparser"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/first-page">A link to first page</a>
  <a href="/second-page">A link to second page</a>
  <a href="/third-page">A link to third page</a>
</body>
</html>
`

func main() {
	reader := strings.NewReader(exampleHtml)
	links, err := htmlparser.Parse(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
