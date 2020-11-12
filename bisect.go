package urlbisect

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Bisect runs binary search on a template url with autoincrement numbers and
// will report back the last number that did yield a 200 OK.
func Bisect(base, placeholder, indicate404 string, min, max int) (int, error) {
	return bisect(base, placeholder, indicate404, min, max)
}

func bisect(base, placeholder, indicate404 string, min, max int) (int, error) {
	var (
		mid  = min + (max-min)/2
		link = strings.Replace(base, placeholder, strconv.Itoa(mid), 1)
	)
	resp, err := http.Get(link)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	log.Printf("[%d] %s %s", resp.StatusCode, link, resp.Request.URL.String())
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode < 400 {
		if indicate404 != "" && strings.Contains(string(b), indicate404) {
			max = mid
			log.Printf("[404] found indicator string: %s", indicate404)
		} else {
			min = mid + 1
		}
	} else {
		max = mid
	}
	if min == max {
		if resp.StatusCode < 400 {
			return mid, nil
		} else {
			return mid - 1, nil
		}
	}
	return bisect(base, placeholder, indicate404, min, max)
}
