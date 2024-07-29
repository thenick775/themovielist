package list

import (
	"fyne.io/fyne/v2/data/binding"
)

type ListData struct {
	Data    binding.ExternalStringList
	StrList []string
}

type ListItem struct {
	Name   string
	Rating int
	Tags   string
}
