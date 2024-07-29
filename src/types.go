package main

import (
	"the-list/list"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// misc structures
type MenuPageLink struct {
	View func(w fyne.Window) fyne.CanvasObject
}

// form structures
type submitEntry struct {
	widget.Entry
	onSubmit func()
}

// inquiry structures
type inquiryEntry struct {
	widget.Entry
	list_loc int // move this to the listData struct??
}
type userList struct {
	Data        map[string][]list.ListItem
	List        *widget.List
	SelectEntry *inquiryEntry
	ShowData    list.ListData
	Modified    bool
}
type Inquiry struct {
	FilterList        string
	SearchMap         map[string]int
	LinkageMap        map[int]int
	ExpandL1          *widget.Label
	ExpandL2          *widget.Label
	ExpandL3          *widget.Label
	InquiryTabs       *container.AppTabs
	InquiryScrollStop bool
	InqTitle          *widget.Label
	InqIntro          *widget.Label
}

// application state structures
type AppState struct {
	currentList       string
	currentMenuItem   string
	noList            bool
	alphaSort         alphaSort
	currentThemeAlias string
}
type alphaSort struct {
	enabled bool
	order   int // 0 asc, 1 desc
}
