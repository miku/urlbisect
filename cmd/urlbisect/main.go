package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	indicate404 = flag.String("r", "", "a regular expression to use as 404 indicator")
	verbose     = flag.Bool("v", false, "be verbose")
	showVersion = flag.Bool("version", false, "show version and exit")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("urlbisect %s %s\n", Version, Buildtime)
		os.Exit(0)
	}
	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}
	v, err := urlbisect.Bisect(*base, *placeholder, *indicate404, *from, *to)
	if err != nil {
		log.Fatal(err)
	}
	link := strings.Replace(*base, *placeholder, strconv.Itoa(v), 1)
	fmt.Printf("%d %s\n", v, link)
}
