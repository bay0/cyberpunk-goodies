<h1 align="center">Welcome to cyberpunk-goodies 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

> Generate pdf's from the cyberpunk cdn

```go 
type Asset struct {
	URL         string //url to the cdn
	Pages       int //max amount of pages
	Images      []string
	Lang        string //de, en, us and so on
	Orientation string //"P" or "Portrait", "L" or "Landscape"
	PDFSize     string //"A3", "A4", "A5", "Letter", "Legal", "Tabloid"
}
```

Define your asset and let it grab all the images and create a pdf.

## Author

👤 **bay**

* Github: [@bay0](https://github.com/bay0)