package main

import (
	"fmt"
	"os"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Deflnam/code-research/internal/simulation"
)

func main() {
	a := app.New()
	w := a.NewWindow("Исследование помехоустойчивого кодирования")
	w.Resize(fyne.NewSize(500, 600))

	entryLength := widget.NewEntry()
	entryLength.SetText("32")

	entryP := widget.NewEntry()
	entryP.SetText("0.01")

	entryMaxErrors := widget.NewEntry()
	entryMaxErrors.SetText("1000")
	entryMaxBlocks := widget.NewEntry()
	entryMaxBlocks.SetText("100000")

	lblTotal := widget.NewLabel("Передано блоков: 0")
	lblDetected := widget.NewLabel("Обнаружено ошибок: 0")
	lblMissed := widget.NewLabel("Пропущено ошибок (ложный прием): 0")

	var lastResult simulation.Result
	var lastConfig simulation.Config

	btnRun := widget.NewButton("Запустить моделирование", func() {
		length, _ := strconv.Atoi(entryLength.Text)
		p, _ := strconv.ParseFloat(entryP.Text, 64)
		maxErr, _ := strconv.Atoi(entryMaxErrors.Text)
		maxBlk, _ := strconv.Atoi(entryMaxBlocks.Text)

		lastConfig = simulation.Config{
			Length:    length,
			P:         p,
			MaxErrors: maxErr,
			MaxBlocks: maxBlk,
		}

		lastResult = simulation.Run(lastConfig)

		lblTotal.SetText(fmt.Sprintf("Передано блоков: %d", lastResult.Total))
		lblDetected.SetText(fmt.Sprintf("Обнаружено ошибок: %d", lastResult.Detected))
		lblMissed.SetText(fmt.Sprintf("Пропущено ошибок (ложный прием): %d", lastResult.Missed))
	})

	btnExport := widget.NewButton("Экспорт результатов в CSV", func() {
		file, err := os.Create("results.csv")
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		defer file.Close()

		file.WriteString("P,Length,MaxErrors,MaxBlocks,Total,Detected,Missed\n")
		file.WriteString(fmt.Sprintf("%f,%d,%d,%d,%d,%d,%d\n",
			lastConfig.P, lastConfig.Length, lastConfig.MaxErrors, lastConfig.MaxBlocks,
			lastResult.Total, lastResult.Detected, lastResult.Missed))

		dialog.ShowInformation("Экспорт", "Данные успешно сохранены в results.csv", w)
	})

	form := container.NewVBox(
		widget.NewLabelWithStyle("Параметры кодирования:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewGridWithColumns(2, widget.NewLabel("Длина блока (байт):"), entryLength),

		widget.NewLabelWithStyle("Свойства канала связи:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewGridWithColumns(2, widget.NewLabel("Вероятность ошибки (P):"), entryP),

		widget.NewLabelWithStyle("Условия остановки:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewGridWithColumns(2, widget.NewLabel("Макс. зафиксированных ошибок:"), entryMaxErrors),
		container.NewGridWithColumns(2, widget.NewLabel("Лимит переданных блоков:"), entryMaxBlocks),

		widget.NewSeparator(),
		btnRun,
		widget.NewSeparator(),

		widget.NewLabelWithStyle("Результаты симуляции:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		lblTotal,
		lblDetected,
		lblMissed,

		widget.NewSeparator(),
		btnExport,
	)

	w.SetContent(form)
	w.ShowAndRun()
}
