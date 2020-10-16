package main

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dlsniper/debugger"
)

func fakeTraffic() {
	// Wait for the server to start
	time.Sleep(1 * time.Second)

	pages := []string{"/", "/login", "/logout", "/products", "/product/{productID}", "/basket", "/about"}

	activeConns := make(chan struct{}, 10)

	c := &http.Client{
		Timeout: 10 * time.Second,
	}

	i := int64(0)

	for {
		activeConns <- struct{}{}
		i++

		page := pages[rand.Intn(len(pages))]

		// We need to launch this using a closure function to
		// ensure that we capture the correct value for the
		// two parameters we need: page and i
		go func(p string, rid int64) {
			// Step 2. Now that the debugger stopped, we can take a goroutines snapshot
			//         using the 📷 button at the bottom of the left debugger toolbar
			//         and then filter them, group them, or hide them
			makeRequest(activeConns, c, p, rid)
		}(page, i)
	}
}

func makeRequest(done chan struct{}, c *http.Client, page string, i int64) {
	defer func() {
		// Unblock the next request from the queue
		<-done
	}()

	debugger.SetLabels(func() []string {
		return []string{
			"request", "automated",
			"page", page,
			"rid", strconv.Itoa(int(i)),
		}
	})

	page = strings.Replace(page, "{productID}", "abc-"+strconv.Itoa(int(i)), -1)
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080"+page, nil)
	if err != nil {
		return
	}

	resp, err := c.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	_, _ = io.Copy(ioutil.Discard, resp.Body)

	debugger.SetLabels(func() []string {
		return []string{
			"request", "automated",
			"page", page,
			"rid", strconv.Itoa(int(i)),
			"status", "before sleep",
		}
	})

	time.Sleep(time.Duration(10+rand.Intn(40)) + time.Millisecond)

	debugger.SetLabels(func() []string {
		return []string{
			"request", "automated",
			"page", page,
			"rid", strconv.Itoa(int(i)),
			"status", "after sleep",
		}
	})
}
