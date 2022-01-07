package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false
var maxLen int = 80

func showAudioPlayer() {
	go func(msg string) {
		fmt.Println(msg)
		if streamer == nil {
		} else {
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")
	time.Sleep(time.Second)
	w := myapp.NewWindow("Audio Player")
	w.Resize(fyne.NewSize(500, 500))
	r, _ := fyne.LoadResourceFromPath("media\\music2.png")
	w.SetIcon(r)
	logo := canvas.NewImageFromFile("media\\music2.png")
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause {
				pause = true
				speaker.Lock()
			} else if pause {
				pause = false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func() {
			speaker.Clear()
		}),
		widget.NewToolbarSpacer(),
	)
	label := widget.NewLabel("Audio Player")
	label.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}
	label2 := widget.NewLabel("Play")
	label2.Alignment = fyne.TextAlignCenter
	browse_files := widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(uc fyne.URIReadCloser, _ error) {
			streamer, format, _ = mp3.Decode(uc)
			if len(uc.URI().Name()) >= maxLen {
				label2.Text = uc.URI().Name()[:maxLen] // slicing is a constant time operation in go
			} else {
				label2.Text = uc.URI().Name()
			}
			label2.Refresh()
		}, w)
		fd.Show()
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})
	img := canvas.NewImageFromFile("media\\itunes_bg.png")
	img.FillMode = canvas.ImageFillStretch
	new_rect := canvas.NewRectangle(color.NRGBA{R: uint8(theme_color), G: uint8(theme_color), B: uint8(theme_color), A: 150})
	main_container := container.New(layout.NewMaxLayout(), img,
		container.New(layout.NewMaxLayout(), new_rect,
			container.New(layout.NewVBoxLayout(),
				label,
				browse_files,
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(),
					layout.NewSpacer(), container.NewGridWrap(fyne.NewSize(150, 150), container.NewPadded(logo)), layout.NewSpacer()),
				layout.NewSpacer(),
				label2,
				toolbar,
			),
		),
	)
	w.SetContent(
		container.NewBorder(nil, nil, nil, nil, main_container),
	)
	w.CenterOnScreen()
	w.Show()
}
