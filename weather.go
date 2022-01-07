package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"

	"net/http"

	"fyne.io/fyne/v2/widget"
)

var city string = "mumbai"

func showWeatherApp() {
	w := myapp.NewWindow("Weather App")
	w.Resize(fyne.NewSize(600, 500))
	r, _ := fyne.LoadResourceFromPath("media\\weathericon.png")
	w.SetIcon(r)

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&APPID=b2d327a4ff36fdcee7cd18945283e86f")
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Println(err)
	}
	label1 := canvas.NewText("Weather Details:", color.White)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextSize = 15
	label2 := canvas.NewText(fmt.Sprintf("%s %s", weather.Name, weather.Sys.Country), color.White)
	label2.Alignment = fyne.TextAlignCenter
	label2.TextSize = 30
	label3 := canvas.NewText(fmt.Sprintf("%s", weather.Weather[0].Description), color.White)
	label3.Alignment = fyne.TextAlignCenter
	weather_icon := canvas.NewImageFromFile("media\\w_ico.png")
	weather_icon.FillMode = canvas.ImageFillOriginal
	label4 := canvas.NewText(fmt.Sprintf("\t%.2f  °C", weather.Main.Temp/10), color.White)
	label4.Alignment = fyne.TextAlignCenter
	label4.TextSize = 25
	label5 := canvas.NewText(fmt.Sprintf("Wind Speed: \t%.2f mph", weather.Wind.Speed), color.White)
	label5.Alignment = fyne.TextAlignCenter
	label6 := canvas.NewText(fmt.Sprintf("Humidity: \t%.2d", weather.Main.Humidity), color.White)
	label6.Alignment = fyne.TextAlignCenter
	label7 := canvas.NewText(fmt.Sprintf("Longitude: \t%f", weather.Coord.Lat), color.White)
	label7.Alignment = fyne.TextAlignCenter
	label8 := canvas.NewText(fmt.Sprintf("Latitude: \t%f", weather.Coord.Lon), color.White)
	label8.Alignment = fyne.TextAlignCenter

	combo := widget.NewSelect([]string{"Mumbai", "Delhi", "Ahmedabad", "Chennai", "Bangalore", "Kolkata",
		"Surat",
		"Pune",
		"Thane",
		"Hyderabad",
		"Jaipur",
		"Lucknow",
		"Kanpur",
		"Nagpur",
		"Indore",
		"Bhopal",
		"Amritsar",
		"Ranchi",
		"Gurgaon",
		"Noida"}, func(value string) {

		res, err = http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + value + "&APPID=b2d327a4ff36fdcee7cd18945283e86f")
		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		weather, err = UnmarshalWeather(body)
		if err != nil {
			fmt.Println(err)
		}
		label2.Text = weather.Name + " " + weather.Sys.Country
		label3.Text = weather.Weather[0].Description
		label4.Text = fmt.Sprintf("%.2f", weather.Main.Temp/10) + " °C"
		label5.Text = fmt.Sprintf("Wind Speed: \t%.2f", weather.Wind.Speed) + " mph"
		label6.Text = fmt.Sprintf("Humidity: \t%.2d", weather.Main.Humidity) + " %"
		label7.Text = fmt.Sprintf("Longitude: \t%f", weather.Coord.Lat)
		label8.Text = fmt.Sprintf("Latitude: \t%f", weather.Coord.Lon)

		label2.Refresh()
		label3.Refresh()
		label4.Refresh()
		label5.Refresh()
		label6.Refresh()
		label7.Refresh()
		label8.Refresh()

	})
	img := canvas.NewImageFromFile("media\\weather1.jpg")
	img.FillMode = canvas.ImageFillStretch

	rect1 := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 150})
	c1 := container.New(layout.NewHBoxLayout(),
		container.NewGridWrap(fyne.NewSize(200, 35), combo),
		layout.NewSpacer(),

		container.NewGridWrap(fyne.NewSize(400, 500), container.NewPadded(rect1,
			container.New(layout.NewVBoxLayout(),
				container.NewPadded(label1),
				container.NewPadded(label2),
				container.NewPadded(label3),
				container.New(layout.NewCenterLayout(), container.NewGridWrap(fyne.NewSize(100, 100), container.NewPadded(weather_icon))),
				container.NewPadded(label4),
				layout.NewSpacer(),
				container.NewPadded(label5),
				container.NewPadded(label6),
				container.NewPadded(label7),
				container.NewPadded(label8),
			))),
	)
	weatherContainer := container.New(layout.NewMaxLayout(), img,
		container.New(layout.NewBorderLayout(nil, nil, nil, nil), c1))
	w.SetContent(container.NewBorder(nil, nil, nil, nil, weatherContainer))
	w.CenterOnScreen()
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
