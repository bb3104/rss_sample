package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ParseRss interface {
	ParseRss(url string)
}

type HatenaXML struct {
	Bookmarks []struct {
		Title    string `xml:"title"`
		Link     string `xml:"link"`
		Date     string `xml:"date"`
		Count    int    `xml:"bookmarkcount"`
		ImageUrl string `xml:"imageurl"`
	} `xml:"item"`
}

type ItmediaXML struct {
	Itmedias []struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		PubDate     string `xml:"pubDate"`
		Description string `xml:"discription"`
	} `xml:"item"`
}

func (h HatenaXML) ParseRss(url string) {

	data := httpGet(url)

	result := HatenaXML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, bookmark := range result.Bookmarks {
		datetime, _ := time.Parse(time.RFC3339, bookmark.Date)

		fmt.Printf("%v\n", datetime.Format("2006/01/02 15:04:05"))
		fmt.Printf("%s - %duser\n", bookmark.Title, bookmark.Count)
		fmt.Printf("%v\n", bookmark.Link)
		fmt.Printf("%v\n", bookmark.ImageUrl)
		fmt.Println()
	}

}

func (i ItmediaXML) ParseRss(url string) {

	data := httpGet(url)

	result := ItmediaXML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, itmedia := range result.Itmedias {
		datetime, _ := time.Parse(time.RFC3339, itmedia.PubDate)

		fmt.Printf("%v\n", datetime.Format("2006/01/02 15:04:05"))
		fmt.Printf("%v\n", itmedia.Title)
		fmt.Printf("%v\n", itmedia.Link)
		fmt.Printf("%v\n", itmedia.Description)
		fmt.Println()
	}

}

func main() {
	var hatena_xml HatenaXML
	var itmedia_xml ItmediaXML
	var rss ParseRss

	rss = hatena_xml
	rss.ParseRss("http://b.hatena.ne.jp/hotentry/it.rss")

	rss = itmedia_xml
	rss.ParseRss("https://rss.itmedia.co.jp/rss/1.0/news_bursts.xml")

}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}
