package urlbisect

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	ValueError = errors.New("value error")
)

// Bisect runs binary search on a template url with autoincrement numbers and
// will report back the last number and url that did yield a 200 OK.
func Bisect(base, placeholder string, min, max int) (int, error) {
	return bisect(base, placeholder, min, max)
}

func bisect(base, placeholder string, min, max int) (int, error) {
	var (
		mid  = min + ((max - min) / 2)
		link = strings.Replace(base, placeholder, strconv.Itoa(mid), 1)
	)
	resp, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	log.Printf("%d %s", resp.StatusCode, link)
	if resp.StatusCode < 400 {
		if bytes.Contains(b, []byte("404")) {
			max = mid
		} else {
			min = mid
		}
	} else {
		max = mid
	}
	if min == max {
		return mid, nil
	}
	return bisect(base, placeholder, min, max)
}
