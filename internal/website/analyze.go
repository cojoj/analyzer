package website

import (
	"bytes"
	"io"
	"net/url"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/cojoj/analyzer/internal/models"
	"golang.org/x/net/html"
)

// Analyze performs full analysis od the provided website's URL. It tokenize HTML and iterates
// through it in order to find all necessary components.
func Analyze(url *url.URL, r io.Reader) (*models.Report, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	report := models.Report{
		URL:               url,
		DTD:               documentType(doc),
		Title:             pageTitle(doc),
		Headings:          headings(doc),
		ContainsLoginForm: hasLoginForm(doc),
		Links:             links(url, doc),
	}

	return &report, nil
}

func documentType(d *goquery.Document) string {
	firstElement := d.Selection.Get(0).FirstChild
	if firstElement.Type != html.DoctypeNode {
		return "Not specified"
	}

	buf := new(bytes.Buffer)
	err := html.Render(buf, firstElement)
	if err != nil {
		return ""
	}

	return buf.String()
}

func pageTitle(d *goquery.Document) string {
	return d.Find("title").Contents().Text()
}

func headings(d *goquery.Document) []models.Heading {
	var h []models.Heading
	for i := 1; i <= 6; i++ {
		lvl := "h" + strconv.Itoa(i)
		number := d.Find(lvl).Length()
		if number != 0 {
			h = append(h, models.Heading{
				Level:  lvl,
				Amount: number,
			})
		}
	}

	return h
}

func hasLoginForm(d *goquery.Document) bool {
	return d.Find("input[type='password']").Length() >= 1
}

func links(u *url.URL, d *goquery.Document) []models.Link {
	var linksMap = make(map[string]models.Link)

	d.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		url, err := url.Parse(href)
		if err != nil {
			return
		}

		if url.Host == "" {
			url.Host = u.Host
		}

		if url.Scheme == "" {
			url.Scheme = u.Scheme
		}

		linksMap[url.String()] = models.Link{
			URL:      url,
			Count:    linksMap[url.String()].Count + 1,
			Internal: url.Host == u.Host,
		}
	})

	var links []models.Link
	for _, value := range linksMap {
		links = append(links, value)
	}

	check(links)

	return links
}

func check(links []models.Link) {
	var wg sync.WaitGroup
	maxParallel := 25
	semaphore := make(chan struct{}, maxParallel)

	for i := range links {
		wg.Add(1)
		go func(l *models.Link) {
			semaphore <- struct{}{}
			reachable, code := Reachable(l.URL.String())
			l.Status = &models.Status{
				Reachable:  reachable,
				StatusCode: code,
			}
			wg.Done()
			<-semaphore
		}(&links[i])
	}

	wg.Wait()
}
