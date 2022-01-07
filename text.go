package main
import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
)
var count int = 1
var filepath string
func showTextEditor() {

    w := myapp.NewWindow("NotePad")
	w.Resize(fyne.NewSize(500, 500))
    input := widget.NewMultiLineEntry()
    input.SetPlaceHolder("Enter text...")
    input.Move(fyne.NewPos(0, 0))
    input.Resize(fyne.NewSize(500, 500))
    new1 := fyne.NewMenuItem("New", func() {
        filepath = ""
        w.SetTitle("NotePad")
        input.Text = ""
        input.Refresh()
    })
    save1 := fyne.NewMenuItem("Save", func() {
        if filepath != "" {
            f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
            if err != nil {

            }
            defer f.Close()
            f.WriteString(input.Text)
        } else {
            saveFileDialog := dialog.NewFileSave(
                func(r fyne.URIWriteCloser, _ error) {
                    textData := []byte(input.Text)
                    r.Write(textData)
                    filepath = r.URI().Path()
                    w.SetTitle(filepath)
                }, w)
            saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
            saveFileDialog.Show()
        }
    })
    saveAs1 := fyne.NewMenuItem("Save as..", func() {
        saveFileDialog := dialog.NewFileSave(
            func(r fyne.URIWriteCloser, _ error) {
                textData := []byte(input.Text)
                r.Write(textData)
                filepath = r.URI().Path()
                w.SetTitle(filepath)
            }, w)
        saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
        saveFileDialog.Show()
    })
    open1 := fyne.NewMenuItem("Open", func() {
        openfileDialog := dialog.NewFileOpen(
            func(r fyne.URIReadCloser, _ error) {
                data, _ := ioutil.ReadAll(r)
                result := fyne.NewStaticResource("name", data)
                input.SetText(string(result.StaticContent))
                fmt.Println(result.StaticName + r.URI().Path())
                filepath = r.URI().Path()
                w.SetTitle(filepath)
            }, w)
        openfileDialog.SetFilter(
            storage.NewExtensionFileFilter([]string{".txt"}))
        openfileDialog.Show()
    })
    menuItem := fyne.NewMenu("File", new1, save1, saveAs1, open1)
    menux1 := fyne.NewMainMenu(menuItem)
    w.SetMainMenu(menux1)
    textContainer:=container.NewWithoutLayout(input,)
    r, _ := fyne.LoadResourceFromPath("media\\notepad.png")
	w.SetIcon(r)
    w.SetContent(container.NewBorder(nil,nil,nil,nil,textContainer))
    w.CenterOnScreen()
    w.Show()
}
