package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/miku/urlbisect"
)

var (
	Buildtime string
	Version   string

	base        = flag.String("u", "", "base url template")
	placeholder = flag.String("p", "@", "placeholder in the URL to be replaced with an integer")
	from        = flag.Int("f", 0, "count from")
	to          = flag.Int("t", 1000000, "count to")
)

func main() {
	flag.Parse()

	v, err := urlbisect.Bisect(*base, *placeholder, *from, *to)
	if err != nil {
		log.Fatal(err)
	}
	link := strings.Replace(*base, *placeholder, strconv.Itoa(v), 1)
	fmt.Sprintf("%d %s", v, link)
}
