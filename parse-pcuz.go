package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
)

var (
	u = flag.String("u", "http://www.pc.uz/trade/orgs/cat1013", "url to pc.uz cat to parse")
)

func init() {
	flag.Parse()
	flag.Usage()
}

func main() {
	// get all list in one go
	suf := "?&sort=0&limit=10000"
	orgs := make([]Org, 0)
	url := *u + suf
	// get html page
	r, err := http.Get(url)
	fatalOnErr(err)
	// find core element to parse
	defer r.Body.Close()
	// start time count
	start := time.Now()
	d, err := html.Parse(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// find table containing list
	slTable := selector{
		Tag:  atom.Table,
		Attr: "width",
		Val:  "98%",
	}

	tables := findNodes(d, slTable)
	// @todo should be more flexible
	if len(tables) != 2 {
		log.Fatalln("html structure changed, can't find required file", len(tables))
	}
	// get required table node to process
	tbl := tables[1]
	// find org titles
	slT := selector{
		Tag: atom.Strong,
	}
	elmsT := findNodes(tbl, slT)
	for _, v := range elmsT {
		orgs = append(orgs, Org{Name: getText(v)})
	}
	// select about td nodes
	sls := selector{
		Tag:  atom.Td,
		Attr: "class",
		Val:  "line_about",
	}
	elms := findNodes(tbl, sls)
	// about elms
	slD := selector{
		Tag:  atom.Div,
		Attr: "style",
		Val:  "padding-bottom:1px",
	}

	for k, v := range elms {
		divs := findNodes(v, slD)
		orgs[k].Tel = getPhone(divs[0])
	}

	b, err := json.Marshal(orgs)
	fatalOnErr(err)
	// create file
	cn := strings.Split(url, "/")
	fname := "pcuz-cat-parse-" + cn[len(cn)-1] + ".json"
	f, err := os.Create(fname)
	fatalOnErr(err)
	defer f.Close()

	_, err = f.Write(b)
	fatalOnErr(err)
	fmt.Println("html page parsed and saved as json object in " + time.Since(start).String())

}

type Org struct {
	Name string `json:name;`
	Tel  string `json:tel;`
}

func getPhone(n *html.Node) string {
	var s string
	if n.LastChild != nil {
		s = n.LastChild.FirstChild.Data
	}
	return s
}

func getText(n *html.Node) string {
	var s string
	// simplest will be
	if n.FirstChild != nil {
		s = n.FirstChild.Data
	}
	return s
}

type selector struct {
	Tag  atom.Atom
	Attr string
	Val  string
}

func findNodes(h *html.Node, s selector) []*html.Node {
	nodes := make([]*html.Node, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom == s.Tag {
			// if only tag given find don't search for attr
			if s.Attr != "" {
				for _, a := range n.Attr {
					if a.Key == s.Attr && a.Val == s.Val {
						// save pointer to nodes
						nodes = append(nodes, n)
					}
				}
			} else {
				nodes = append(nodes, n)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(h)
	return nodes
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func logOnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
