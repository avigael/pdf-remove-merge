# PDF Page Remover & Merger

This is a command line tool written in Golang for merging multiple PDFs into a single PDF as well as removing specific pages from any PDF.

## Build
```
$ git clone https://github.com/avigael/pdf-remove-merge.git
$ cd pdf-remove-merge
$ go build editpdf.go
```

### To Merge PDFs
```
$ ./editpdf merge [output.pdf] [input-1.pdf input-2.pdf ...]
```
You can merge any number of PDFs, but a merge requires at least 2 PDFs to execute.

### To Remove pages from PDF
```
$ ./editpdf remove [output.pdf] [input.pdf] [page-num...]
```
Currently you have to type the individual pages that you wish to remove, but I might add the ability to provide a range later.

Feel free to use this code in anyway you wish.

Credit to [pdfcpu](https://github.com/pdfcpu/pdfcpu "pdfcpu") for the PDF processing package

