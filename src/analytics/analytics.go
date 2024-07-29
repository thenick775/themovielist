package analytics

import (
	"image"
	"image/color"
	"strings"
	"the-list/list"

	"github.com/psykhi/wordclouds"
)

var defaultConfImg = confImg{
	FontMaxSize:     600,
	FontMinSize:     15,
	RandomPlacement: false,
	FontFile:        "Roboto-Regular.ttf", // prepended with fontLoc once initialized
	Colors:          darkModeColors,       // dark is default
	BackgroundColor: darkModeBackground,   // dark is default
	Width:           2048,
	Height:          2048,
	Mask: maskConf{
		"",
		color.RGBA{ // no masking by default
			R: 0,
			G: 0,
			B: 0,
			A: 0,
		}},
}

// generate image object containing word cloud
func GenWordCloudImg(list []list.ListItem, fontDir string, isLightMode bool) (image.Image, map[string]int) {
	confImg := defaultConfImg
	confImg.FontFile = fontDir + "/" + confImg.FontFile
	// exclusion zones if present
	var boxes []*wordclouds.Box
	if confImg.Mask.File != "" {
		boxes = wordclouds.Mask(
			confImg.Mask.File,
			confImg.Width,
			confImg.Height,
			confImg.Mask.Color,
		)
	}
	// word colors
	if isLightMode {
		confImg.Colors = lightModeColors
		confImg.BackgroundColor = lightModeBackground
	}

	colors := make([]color.Color, 0)
	for _, c := range confImg.Colors {
		colors = append(colors, c)
	}
	// data processing
	wordCounts := make(map[string]int)
	for _, item := range list {
		splitTags := strings.Fields(item.Tags)

		for _, tag := range splitTags {
			wordCounts[tag] += 1
		}
	}

	cloud := wordclouds.NewWordcloud(
		wordCounts,
		wordclouds.FontFile(confImg.FontFile),
		wordclouds.FontMaxSize(confImg.FontMaxSize),
		wordclouds.FontMinSize(confImg.FontMinSize),
		wordclouds.Colors(colors),
		wordclouds.BackgroundColor(confImg.BackgroundColor),
		wordclouds.MaskBoxes(boxes),
		wordclouds.Height(confImg.Height),
		wordclouds.Width(confImg.Width),
		wordclouds.RandomPlacement(confImg.RandomPlacement),
	)
	// image generation
	img := cloud.Draw()
	return img, wordCounts
}

func GenStats(lists map[string][]list.ListItem) []listsSummary {
	var ret []listsSummary
	for name, list := range lists {
		listSum := listsSummary{
			Name:                  name,                 // name of list
			TotalContentCount:     len(list),            // total number of items in list
			ContentCountPerRating: getRatingCount(list), // count of items in list per rating
		}
		ret = append(ret, listSum)
	}

	return ret
}

func getRatingCount(list []list.ListItem) map[int]int {
	var ret = make(map[int]int)
	for _, listItem := range list {
		ret[listItem.Rating] += 1
	}
	return ret
}
