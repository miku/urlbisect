package urlbisect

import (
	"fmt"
	"io"
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
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
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

func ScanHandle(min, max int, w io.Writer) error {
	for i := min; i < max; i++ {
		base := fmt.Sprintf("https://hdl.handle.net/%d/@", i)
		v, err := Bisect(base, "@", "", 0, 1000000)
		if err != nil {
			log.Printf("skipping failed: %s, %v", base, err)
			continue
		}
		if v == -1 {
			fmt.Fprintf(w, "ER %s\n", base)
		} else {
			fmt.Fprintf(w, "OK %s %d\n", base, v)
		}
	}
	return nil
}
