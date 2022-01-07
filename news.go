package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var num_article int = 1
var news News
var news_url *url.URL

func showNews() {
	w := myapp.NewWindow("News App")
	w.Resize(fyne.NewSize(400, 200))
	URL := "https://gnews.io/api/v4/top-headlines?token=f4242aaddbb109f07654431e207e2412&lang=en&max=100"
	//API
	res, _ := http.Get(URL)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	news, _ = UnmarshalNews(body)
	fmt.Println(news)
	// num_article = int(news.TotalArticles)

	//show title
	label3 := widget.NewLabel(news.Articles[1].Title)
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label3.Wrapping = fyne.TextWrapBreak
	// show articles
	entry1 := widget.NewLabel(news.Articles[1].Description)
	//entry1.MultiLine = true
	entry1.Wrapping = fyne.TextWrapBreak

	news_url, _ = url.Parse(news.Articles[1].URL)
	new_link := widget.NewHyperlink("", news_url)
	next_icon := canvas.NewImageFromFile("media\\next-icon.jpg")
	link_icon := canvas.NewImageFromFile("media\\link.jpg")
	img := canvas.NewImageFromFile("media\\news2.png")
	img.FillMode = canvas.ImageFillOriginal
	img_container := container.NewGridWrap(fyne.NewSize(400, 250), container.NewPadded(img))
	btn := widget.NewButton("", func() {
		num_article += 1
		label3.Text = news.Articles[num_article].Title
		entry1.Text = news.Articles[num_article].Description
		news_url, _ = url.Parse(news.Articles[num_article].URL)
		new_link.URL = news_url
		new_link.Refresh()
		label3.Refresh()
		entry1.Refresh()
	})
	x := container.New(layout.NewVBoxLayout(), label3, layout.NewSpacer(), entry1, layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(),
			container.NewGridWrap(fyne.NewSize(25, 25), container.NewPadded(new_link, link_icon)),
			layout.NewSpacer(),
			container.NewGridWrap(fyne.NewSize(25, 25), container.NewPadded(btn, next_icon)),
		),
	)

	r, _ := fyne.LoadResourceFromPath("media\\newslogo.png")
	w.SetIcon(r)
	// e.Resize(fyne.NewSize(300, 300))
	w.SetContent(container.NewBorder(img_container, nil, nil, nil, x))
	w.CenterOnScreen()
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()
func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	TotalArticles int64     `json:"totalArticles"`
	Articles      []Article `json:"articles"`
}
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
	Source      Source `json:"source"`
}
type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
