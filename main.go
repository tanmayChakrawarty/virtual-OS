package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myapp fyne.App = app.New()
var mywindow fyne.Window = myapp.NewWindow("Virtual OS");
var drawer_toggle bool = false; 
var theme_toggle bool = false;
var theme_color int = 0;
var rect *canvas.Rectangle;
var rect2 *canvas.Rectangle;
var bg_color *canvas.Rectangle;


func updateTime(datetime *canvas.Text){
	datetime.Text = time.Now().Format("Jan 02  15:04:05")
	datetime.Refresh()
}

func main() {

	myapp.Settings().SetTheme(theme.DarkTheme())

	dt := time.Now()
	datetime := canvas.NewText(dt.Format("Jan 02  15:04:05"),color.White)
	datetime.TextSize = 12
	datetime.Alignment = fyne.TextAlignCenter
	datetime.TextStyle = fyne.TextStyle{Bold: true}
	
	img := canvas.NewImageFromFile("media\\wallpaper5.jpg")
	img.FillMode = canvas.ImageFillStretch

	icon1 :=canvas.NewImageFromFile("media\\calclogo.png")
	icon1.FillMode = canvas.ImageFillContain
	calc := widget.NewButton("",func ()  {
		showCalc()
	})
	icon2 :=canvas.NewImageFromFile("media\\weathericon.png")
	icon2.FillMode = canvas.ImageFillContain
	weather := widget.NewButton("",func ()  {
		showWeatherApp()
	})
	icon3 :=canvas.NewImageFromFile("media\\gallery1.png")
	icon3.FillMode = canvas.ImageFillContain
	gallery := widget.NewButton("",func ()  {
		showGalleryApp()
	})
	icon4 :=canvas.NewImageFromFile("media\\notepad.png")
	icon4.FillMode = canvas.ImageFillContain
	notepad := widget.NewButton("",func ()  {
		showTextEditor()
	})
	icon5 :=canvas.NewImageFromFile("media\\music2.png")
	icon5.FillMode = canvas.ImageFillContain
	music := widget.NewButton("",func ()  {
		showAudioPlayer()
	})

	icon6 :=canvas.NewImageFromFile("media\\newslogo.png")
	icon6.FillMode = canvas.ImageFillContain
	news := widget.NewButton("",func ()  {
		showNews()
	})

	icon7 :=canvas.NewImageFromFile("media\\shutdown.png")
	icon7.FillMode = canvas.ImageFillContain
	Shutdown := widget.NewButton("", func() {
		mywindow.Close()
	})

	rect= canvas.NewRectangle(color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180})
	rect.Hide()
	apps1 := container.New(layout.NewGridLayoutWithRows(5),
	layout.NewSpacer(),
	container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(calc, icon1)),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(weather, icon2)),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(gallery, icon3)),
		layout.NewSpacer(),
	),
	layout.NewSpacer(),
	container.New(layout.NewHBoxLayout(),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(notepad, icon4)),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(music, icon5)),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(60,60), container.NewPadded(news, icon6)),
		layout.NewSpacer(),
	),	
	layout.NewSpacer(),
	)
	apps1.Hide()
	rect_container :=container.NewGridWrap(fyne.NewSize(500,400), container.NewPadded(rect,apps1))

	icon0 :=canvas.NewImageFromFile("media\\drawer.png")
	icon0.FillMode = canvas.ImageFillContain
	app_drawer := widget.NewButton("",func ()  {
		drawer_toggle = !drawer_toggle
		if drawer_toggle != true{
			rect.Hide()
			apps1.Hide()
		}else{
			rect.Show()
			apps1.Show()
		}
	})
	
	bg_color = canvas.NewRectangle(color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:150})

	
	icon8 :=canvas.NewImageFromFile("media\\lighttheme.png")
	icon8.FillMode = canvas.ImageFillContain
	lt_theme := widget.NewButton("", func() {
		myapp.Settings().SetTheme(theme.LightTheme())
		theme_color = 255
		rect.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		rect2.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		bg_color.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		rect.Refresh()
		rect2.Refresh()
		bg_color.Refresh()
	})
	icon9 :=canvas.NewImageFromFile("media\\darktheme.png")
	icon9.FillMode = canvas.ImageFillContain
	drk_theme := widget.NewButton("", func() {
		myapp.Settings().SetTheme(theme.DarkTheme())
		theme_color = 0
		rect.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		rect2.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		bg_color.FillColor= color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180}
		rect.Refresh()
		rect2.Refresh()
		bg_color.Refresh()
	})

	rect2= canvas.NewRectangle(color.NRGBA{R:uint8(theme_color),G:uint8(theme_color),B:uint8(theme_color),A:180})
	rect2.Hide()

	themes := container.New(layout.NewVBoxLayout(),
		container.NewGridWrap(fyne.NewSize(160,90), container.NewPadded(lt_theme, icon8)),
		layout.NewSpacer(),
		container.NewGridWrap(fyne.NewSize(160,90), container.NewPadded(drk_theme, icon9)),
	)
	themes.Hide()
	theme_container :=container.NewGridWrap(fyne.NewSize(170,200), container.NewPadded(rect2,themes))
	settings_icon :=canvas.NewImageFromFile("media\\setting.png")
	settings_icon.FillMode = canvas.ImageFillContain
	setting := widget.NewButton("",func ()  {
		theme_toggle = !theme_toggle
		if theme_toggle != true{
			rect2.Hide()
			themes.Hide()
			rect.Hide()
			apps1.Hide()
		}else{
			rect2.Show()
			themes.Show()
			rect.Hide()
			apps1.Hide()
		}
	})
	rect3:= canvas.NewRectangle(color.RGBA{R:29,G:29,B:29,A:255})
	datetime_bar:= container.New(
		layout.NewMaxLayout(),
		rect3,
		datetime,
	)
	apps := container.New(layout.NewVBoxLayout(),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(app_drawer, icon0)),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(setting, settings_icon)),
	layout.NewSpacer(),
	container.NewGridWrap(fyne.NewSize(45,45), container.NewPadded(Shutdown, icon7)),
	)

	pane:= container.New(
		layout.NewMaxLayout(),
		bg_color,
		apps,
	)
	mywindow.Resize(fyne.NewSize(1280,720))
	mywindow.CenterOnScreen()
	r, _ := fyne.LoadResourceFromPath("media\\oslogo.png")
	mywindow.SetIcon(r)
	mywindow.SetContent(
		container.New(layout.NewMaxLayout(), img,
		container.New(layout.NewBorderLayout(datetime_bar,nil,pane, nil), datetime_bar, pane, rect_container, theme_container)))
	go func(){
		t := time.NewTicker(time.Second)

		for range t.C{
			updateTime(datetime)
		}
	}()
	mywindow.ShowAndRun()
}