package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/pterm/pterm"
)

type asset struct {
	URL         string
	Pages       int
	Images      []string
	Lang        string
	Orientation string
	PDFSize     string
}

func getImageBytesFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func assetGeneration(asset asset, pdfname string) error {
	pterm.Println("Processing: ", pdfname)
	//clean images
	err := os.RemoveAll("./images")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.MkdirAll("./images", os.FileMode(0777))
	if err != nil {
		log.Println(err)
		return err
	}

	for i := 1; i < asset.Pages+1; i++ {
		asset.Images = append(asset.Images, fmt.Sprintf(asset.URL, asset.Lang, i))
	}

	p, _ := pterm.DefaultProgressbar.WithTotal(asset.Pages).WithTitle("Downloading Images").Start()
	for i, url := range asset.Images {
		imgBytes, err := getImageBytesFromURL(url)
		if err != nil {
			log.Fatal("GetImageBytesFromURL: ", err)
			p.Stop()
			return err
		}

		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			log.Println("Decode: ", err)
			p.Stop()
			return err
		}

		out, err := os.Create(fmt.Sprintf("./images/page-%d.jpeg", i+1))
		if err != nil {
			log.Println(err)
			p.Stop()
			return err
		}
		defer out.Close()

		err = jpeg.Encode(out, img, nil)
		if err != nil {
			log.Println("Encode: ", err)
			p.Stop()
			return err
		}

		pterm.Success.Println("Downloading " + url)
		p.Increment()
	}
	p.Stop()

	//pdf generation
	pdf := gofpdf.New(asset.Orientation, "mm", asset.PDFSize, "")
	w, h := pdf.GetPageSize()

	p, _ = pterm.DefaultProgressbar.WithTotal(asset.Pages).WithTitle("PDF Generation").Start()
	for i := 1; i < asset.Pages+1; i++ {
		pdf.AddPage()
		pdf.Image(fmt.Sprintf("./images/page-%d.jpeg", i), 0, 0, w, h, false, "", 0, "")
		pterm.Success.Println(fmt.Sprintf("Page %d", i))
		p.Increment()
	}
	p.Stop()

	err = pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", pdfname))
	if err != nil {
		log.Fatal("OutputFileAndClose: ", err)
		os.Exit(1)
	}

	pterm.Println("Done!!!")
	return nil
}

const second = time.Second

func introScreen() {
	pterm.DefaultHeader. // Use DefaultHeader as base
				WithMargin(15).
				WithBackgroundStyle(pterm.NewStyle(pterm.BgLightYellow)).
				WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
				Println("cyberpunk-goodies")

	pterm.Info.Println("Generate comic and artbook pdf via the cyberpunk cdn" +
		"\nWill use your system locale." +
		"\nFallback locale is en." +
		"\n" +
		"\nMade by bay (https://github.com/bay0/cyberpunk-goodies)" +
		"\n" + pterm.Green(time.Now().Format("02 Jan 2006 - 15:04:05 MST")))
	pterm.Println()
}

func clear() {
	print("\033[H\033[2J")
}

func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes", "Y":
		return true
	case "n", "no", "N":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}
