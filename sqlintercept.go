package sqlintercept

import (
	"fmt"
	ui "github.com/gizak/termui"
	"github.com/narita-takeru/tcpstream"
	"os"
	"regexp"
	"strings"
	"sync"
)

func Start(src, dst string) {

	mutex := new(sync.Mutex)

	counts := map[string]int{}
	go doCurses(counts)

	t := tcpstream.Thread{}
	t.SrcToDstHook = func(b []byte) []byte {
		regexTableName := regexp.MustCompile("[f|F][r|R][o|O][m|M] ([^ ]+)")
		group := regexTableName.FindSubmatch(b)
		if 0 < len(group) {
			tblName := string(group[1])
			mutex.Lock()
			counts[tblName]++
			mutex.Unlock()
			ui.StopLoop()
		}

		return b
	}

	t.Do(src, dst)
}

type outputItem struct {
	Text  string
	Count int
}

func format(counts map[string]int) []outputItem {
	maxLen := len(counts)
	if maxLen <= 0 {
		return nil
	}

	items := make([]outputItem, 0, maxLen)
	for text, count := range counts {
		appended := false
		for i, item := range items {
			if item.Count < count {
				items = append(items[:i+1], items[i:]...)
				items[i] = outputItem{Text: text, Count: count}
				appended = true
				break
			}
		}

		if !appended {
			items = append(items, outputItem{Text: text, Count: count})
		}
	}

	return items
}

func doCurses(counts map[string]int) {

	filterText := ""

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	requireStop := false

	registerFilter := func(f string) {
		filterText = filterText + f
		ui.StopLoop()
	}

	delOneFilter := func() {
		if 0 < len(filterText) {
			filterText = filterText[0:(len(filterText) - 1)]
			ui.StopLoop()
		}
	}

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
		requireStop = true
	})

	ui.Handle("/sys/kbd/a", func(ui.Event) { registerFilter("a") })
	ui.Handle("/sys/kbd/b", func(ui.Event) { registerFilter("b") })
	ui.Handle("/sys/kbd/c", func(ui.Event) { registerFilter("c") })
	ui.Handle("/sys/kbd/d", func(ui.Event) { registerFilter("d") })
	ui.Handle("/sys/kbd/e", func(ui.Event) { registerFilter("e") })
	ui.Handle("/sys/kbd/f", func(ui.Event) { registerFilter("f") })
	ui.Handle("/sys/kbd/g", func(ui.Event) { registerFilter("g") })
	ui.Handle("/sys/kbd/h", func(ui.Event) { registerFilter("h") })
	ui.Handle("/sys/kbd/i", func(ui.Event) { registerFilter("i") })
	ui.Handle("/sys/kbd/j", func(ui.Event) { registerFilter("j") })
	ui.Handle("/sys/kbd/k", func(ui.Event) { registerFilter("k") })
	ui.Handle("/sys/kbd/l", func(ui.Event) { registerFilter("l") })
	ui.Handle("/sys/kbd/m", func(ui.Event) { registerFilter("m") })
	ui.Handle("/sys/kbd/n", func(ui.Event) { registerFilter("n") })
	ui.Handle("/sys/kbd/o", func(ui.Event) { registerFilter("o") })
	ui.Handle("/sys/kbd/p", func(ui.Event) { registerFilter("p") })
	ui.Handle("/sys/kbd/q", func(ui.Event) { registerFilter("q") })
	ui.Handle("/sys/kbd/r", func(ui.Event) { registerFilter("r") })
	ui.Handle("/sys/kbd/s", func(ui.Event) { registerFilter("s") })
	ui.Handle("/sys/kbd/t", func(ui.Event) { registerFilter("t") })
	ui.Handle("/sys/kbd/u", func(ui.Event) { registerFilter("u") })
	ui.Handle("/sys/kbd/v", func(ui.Event) { registerFilter("v") })
	ui.Handle("/sys/kbd/w", func(ui.Event) { registerFilter("w") })
	ui.Handle("/sys/kbd/x", func(ui.Event) { registerFilter("x") })
	ui.Handle("/sys/kbd/y", func(ui.Event) { registerFilter("y") })
	ui.Handle("/sys/kbd/z", func(ui.Event) { registerFilter("z") })
	ui.Handle("/sys/kbd/_", func(ui.Event) { registerFilter("_") })
	ui.Handle("/sys/kbd/C-8", func(ui.Event) { delOneFilter() })

	filterHeader := ui.NewPar("")
	filterHeader.Height = 3
	filterHeader.Width = 60
	filterHeader.TextFgColor = ui.ColorWhite
	filterHeader.BorderLabel = "filter"
	filterHeader.BorderFg = ui.ColorBlue

	p := ui.NewPar("")
	p.Height = 35
	p.Width = 60
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "access counter"
	p.BorderFg = ui.ColorCyan

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(6, 0, filterHeader)),
		ui.NewRow(ui.NewCol(6, 0, p)),
	)

	ui.Body.Align()

	for !requireStop {
		filterHeader.Text = filterText
		p.Text = ""
		items := format(counts)
		for _, item := range items {
			line := fmt.Sprintf("%30s: %d\n", item.Text, item.Count)
			if "" != filterText {
				if !strings.Contains(line, filterText) {
					continue
				}
			}

			p.Text = p.Text + line
		}

		ui.Render(ui.Body)
		ui.Loop()
	}

	ui.Close()
	os.Exit(0)
}
