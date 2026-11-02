package main

import (
	"errors"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Deflnam/code-research/internal/simulation"
)

func main() {

	a := app.NewWithID("com.deflnam.coderesearch")
	w := a.NewWindow("Исследование помехоустойчивого кодирования")
	w.Resize(fyne.NewSize(500, 650))

	entryLength := widget.NewEntry()
	entryLength.SetText("32")

	entryP := widget.NewEntry()
	entryP.SetText("0.01")

	entryMaxErrors := widget.NewEntry()
	entryMaxErrors.SetText("1000")

	entryMaxBlocks := widget.NewEntry()
	entryMaxBlocks.SetText("100000")

	lblTotal := widget.NewLabel("Передано блоков: 0")
	lblCorrect := widget.NewLabel("Корректно принято: 0")
	lblDetected := widget.NewLabel("Обнаружено ошибок: 0")
	lblMissed := widget.NewLabel("Пропущено ошибок (ложный прием): 0")

	var lastResult simulation.Result
	var lastConfig simulation.Config

	var btnRun *widget.Button

	btnRun = widget.NewButton("Запустить моделирование", func() {

		length, err := strconv.Atoi(entryLength.Text)
		if err != nil || length < 1 || length > 1048576 {
			dialog.ShowError(errors.New("Длина блока должна быть от 1 до 1 048 576 байт (1 МБ)"), w)
			return
		}

		p, err := strconv.ParseFloat(entryP.Text, 64)
		if err != nil || p < 0 || p > 1 {
			dialog.ShowError(errors.New("Вероятность ошибки должна быть в диапазоне от 0.0 до 1.0"), w)
			return
		}

		maxErr, err := strconv.Atoi(entryMaxErrors.Text)
		if err != nil || maxErr < 1 {
			dialog.ShowError(errors.New("Лимит ошибок должен быть целым числом больше 0"), w)
			return
		}

		maxBlk, err := strconv.Atoi(entryMaxBlocks.Text)
		if err != nil || maxBlk < 1 {
			dialog.ShowError(errors.New("Лимит блоков должен быть целым числом больше 0"), w)
			return
		}

		lastConfig = simulation.Config{
			Length:    length,
			P:         p,
			MaxErrors: maxErr,
			MaxBlocks: maxBlk,
		}

		btnRun.Disable()
		lblTotal.SetText("Выполняется симуляция, подождите...")

		go func() {

			defer btnRun.Enable()

			lastResult = simulation.Run(lastConfig)

			lblTotal.SetText(fmt.Sprintf("Передано блоков: %d", lastResult.Total))
			lblCorrect.SetText(fmt.Sprintf("Корректно принято: %d", lastResult.Correct))
			lblDetected.SetText(fmt.Sprintf("Обнаружено ошибок: %d", lastResult.Detected))
			lblMissed.SetText(fmt.Sprintf("Пропущено ошибок (ложный прием): %d", lastResult.Missed))
		}()
	})

	btnExport := widget.NewButton("Сохранить как (CSV)...", func() {
		if lastResult.Total == 0 {
			dialog.ShowInformation("Внимание", "Сначала запустите моделирование", w)
			return
		}

		fileSaveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if writer == nil {
				return
			}
			defer writer.Close()

			writer.Write([]byte("P,Length,MaxErrors,MaxBlocks,Total,Correct,Detected,Missed\n"))
			dataRow := fmt.Sprintf("%f,%d,%d,%d,%d,%d,%d,%d\n",
				lastConfig.P, lastConfig.Length, lastConfig.MaxErrors, lastConfig.MaxBlocks,
				lastResult.Total, lastResult.Correct, lastResult.Detected, lastResult.Missed)
			writer.Write([]byte(dataRow))

			dialog.ShowInformation("Успех", "Данные успешно сохранены", w)
		}, w)

		fileSaveDialog.SetFileName("results.csv")
		fileSaveDialog.Show()
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
		lblCorrect,
		lblDetected,
		lblMissed,

		widget.NewSeparator(),
		btnExport,
	)

	w.SetContent(form)
	w.ShowAndRun()
}
