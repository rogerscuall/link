package main

import (
	"fmt"
	"strings"

	"github.com/rogerscuall/link"
)

var ehtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
  A link to another page
  </a>
  <a href="/other-page1">Another link</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(ehtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
