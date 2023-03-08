package lunch

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Sanz struct {
	Menu string
}

func Lunch(day string, how string) string {
	// Request
	url := "http://school.gyo6.net/gbsw/food/" + "2023/" + day + "/" + how
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// HTML 읽기
	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 현황판 파싱

	items := html.Find("tr")

	// items 순회하면서 출력
	div := items.Find("div").Eq(0).Text()

	return div
}
