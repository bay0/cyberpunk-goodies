package main

import (
	"fmt"
	"log"

	"github.com/Xuanwo/go-locale"
	"github.com/pterm/pterm"
)

func main() {
	introScreen()

	osLang, err := locale.Detect()
	if err != nil {
		log.Fatal(err)
	}

	lang, _ := osLang.Base()

	pterm.Println(fmt.Sprintf("Found locale: %s", lang.String()))
	pterm.Println()

	pterm.Println("Press 'Enter' to start...")
	fmt.Scanln()

	clear()

	comic := asset{
		URL:         `https://cdn-l-cyberpunk.cdprojektred.com/comicbook/assets-99pDTTxHXq2GD68d/%s/pages/page-%d.jpg`,
		Pages:       82,
		Lang:        "kk",
		PDFSize:     "A5",
		Orientation: "P",
	}

	artbook := asset{
		URL:         `https://cdn-l-cyberpunk.cdprojektred.com/artbook/assets-99pDTTxHXq2GD68d/%s/pages/page-%d.jpg`,
		Pages:       60,
		Lang:        "kk",
		PDFSize:     "A4",
		Orientation: "L",
	}

	err = assetGeneration(comic, fmt.Sprintf("comic_%s", comic.Lang))
	if err != nil {
		clear()
		pterm.Println(fmt.Sprintf("There is no comic for the %s version\nPress 'y' to Download the english version or 'n' to skip it...", lang.String()))
		if askForConfirmation() {
			comic.Lang = "en"
			assetGeneration(comic, fmt.Sprintf("comic_%s", comic.Lang))
		}
	}

	err = assetGeneration(artbook, fmt.Sprintf("artbook_%s", artbook.Lang))
	if err != nil {
		clear()
		pterm.Println(fmt.Sprintf("There is no artbook for the %s version\nPress 'y' to Download the english version or 'n' to skip it...", lang.String()))
		if askForConfirmation() {
			artbook.Lang = "en"
			assetGeneration(artbook, fmt.Sprintf("artbook_%s", artbook.Lang))
		}
	}

	pterm.Println("Press 'Enter' to exit...")
	fmt.Scanln()
}
