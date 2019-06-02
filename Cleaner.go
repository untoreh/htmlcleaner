package HtmlCleaner

import (
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"mvdan.cc/xurls"
)

type Cleaner struct {
	Strategies []string
	Strat      string
	whitelist  string
	imgAttr    map[string]bool
	elAttr     map[string]bool
	titleRgx   []*regexp.Regexp
}

var cln_v = Cleaner{
	Strategies: []string{"article"},
	Strat:      "article",
	whitelist:  "div,span,hr,p,a,img,strong,i,b,em,li,ul,h6,h5,h4,h3,h2,h1,blockquote,html",
	imgAttr: map[string]bool{
		"src":    true,
		"srcset": true,
		"width":  true,
		"height": true,
		"alt":    true,
	},
	elAttr: map[string]bool{"href": true, "title": true, "alt": true, "target": true},
	titleRgx: []*regexp.Regexp{
		regexp.MustCompile(`([\{\(\[][^()\[\]]*[\)\]\}])*`),
		regexp.MustCompile(`\*{2,}`),
		regexp.MustCompile(`(.)1{3,}`),
	},
}

func (cln *Cleaner) CleanBody(body io.Reader) string {
	// Load the HTML document
	dom, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Print(err)
	}
	// log.Printf("parsing...%d", start)
	dom.Not(cln.whitelist).Remove()
	dom.Find("*").Each(func(i int, s *goquery.Selection) {
		switch s.Nodes[0].Data {
		case "img":
			for _, a := range s.Nodes[0].Attr {
				if _, ok := cln.imgAttr[a.Key]; !ok {
					s.RemoveAttr(a.Key)
				}
			}
		default:
			for _, a := range s.Nodes[0].Attr {
				if _, ok := cln.elAttr[a.Key]; !ok {
					s.RemoveAttr(a.Key)
				}
			}
		}
	})
	html, _ := dom.Find("body").Html()
	// log.Printf("end...%d", time.Now().UnixNano()-start)
	return html
}

func (cln *Cleaner) CleanTitle(body io.Reader) (title string) {
	read, err := ioutil.ReadAll(body)
	title = string(read)
	if err != nil {
		log.Printf("couldn't read title: %s", err)
	}
	for _, r := range cln.titleRgx {
		title = r.ReplaceAllString(title, "")
	}
	return
}

func (cln *Cleaner) Linkify(body io.Reader) string {
	// Load the HTML document
	dom, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Print(err)
	}

	// loop over elements which are not anchors
	var c *html.Node
	for _, c = range dom.Find("*").Not("a").Contents().Nodes {
		// log.Print(c.Data)
		if c.Type == 1 {
			var wrap strings.Builder
			text := c.Data
			for _, url := range xurls.Relaxed().FindAllString(text, -1) {
				wrap.Reset()
				wrap.WriteString("<a href=\"")
				wrap.WriteString(url)
				wrap.WriteString("\">")
				wrap.WriteString(url)
				wrap.WriteString("</a>")
				text = strings.ReplaceAll(text, url, wrap.String())
			}
			dom.FindNodes(c).ReplaceWithHtml(text)
			// c.Data = text
		}
	}

	html, _ := dom.Html()
	return html
}

func New(options ...map[string]interface{}) *Cleaner {
	cln := cln_v
	return &cln
}
