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

	"github.com/cheggaaa/pb"
	"github.com/phpdave11/gofpdf"
)

type Asset struct {
	URL         string
	Pages       int
	Images      []string
	Lang        string
	Orientation string
	PDFSize     string
}

func GetImageBytesFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}

func AssetGeneration(asset Asset, pdfname string) {
	log.Println("Processing: ", pdfname)
	//clean images
	err := os.RemoveAll("./images")
	if err != nil {
		log.Println(err)
	}
	os.Mkdir("./images", os.FileMode(0522))

	for i := 1; i < asset.Pages+1; i++ {
		asset.Images = append(asset.Images, fmt.Sprintf(asset.URL, asset.Lang, i))
	}

	log.Println("Downloading Images!!!")
	dlBar := pb.StartNew(len(asset.Images))
	for i, url := range asset.Images {
		imgBytes, err := GetImageBytesFromURL(url)
		if err != nil {
			log.Fatal("GetImageBytesFromURL: ", err)
		}

		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			log.Fatalln("Decode: ", err)
		}

		out, _ := os.Create(fmt.Sprintf("./images/page-%d.jpeg", i+1))
		defer out.Close()

		err = jpeg.Encode(out, img, nil)
		if err != nil {
			log.Println("Encode: ", err)
		}
		dlBar.Increment()
	}
	dlBar.Finish()

	//pdf generation
	log.Println("PDF Generation!!!")
	pdf := gofpdf.New(asset.Orientation, "mm", asset.PDFSize, "")
	w, h := pdf.GetPageSize()

	pdfBar := pb.StartNew(len(asset.Images))
	for i := 1; i < asset.Pages+1; i++ {
		pdf.AddPage()
		pdf.Image(fmt.Sprintf("./images/page-%d.jpeg", i), 0, 0, w, h, false, "", 0, "")
		pdfBar.Increment()
	}
	pdfBar.Finish()

	err = pdf.OutputFileAndClose(fmt.Sprintf("%s.pdf", pdfname))
	if err != nil {
		log.Fatal("OutputFileAndClose: ", err)
		os.Exit(1)
	}

	log.Println("Done!!!")
}

func main() {
	comic := Asset{
		URL:         `https://cdn-l-cyberpunk.cdprojektred.com/comicbook/assets-99pDTTxHXq2GD68d/%s/pages/page-%d.jpg`,
		Pages:       82,
		Lang:        "de",
		PDFSize:     "A5",
		Orientation: "P",
	}

	artbook := Asset{
		URL:         `https://cdn-l-cyberpunk.cdprojektred.com/artbook/assets-99pDTTxHXq2GD68d/%s/pages/page-%d.jpg`,
		Pages:       60,
		Lang:        "de",
		PDFSize:     "A4",
		Orientation: "L",
	}

	AssetGeneration(comic, "comic")
	AssetGeneration(artbook, "artbook")
}
