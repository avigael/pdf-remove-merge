package main

import (
    "fmt"
    pdf "github.com/pdfcpu/pdfcpu/pkg/api"
    "os"
)

func main() {
    var cmd string
    len := len(os.Args)
    if len > 1 {
        cmd = os.Args[1]
    }
    switch cmd {
    case "merge":
        // You need at least 2 PDFs to perform a merge
        if len < 5 {
            fmt.Print("Usage: ./editpdf merge [output.pdf] [input-1.pdf ",
                "input-2.pdf...]\n")
            os.Exit(0)
        }
        // Merge PDFs
        err := pdf.MergeCreateFile(os.Args[3:], os.Args[2], nil)
        if err != nil {
            fmt.Print("MERGE FAILED: ", err)
            os.Exit(1)
        }
        // Merge Successful
        fmt.Print("Successfully merged PDFs!\nOutput -> ", os.Args[2])
    case "remove":
        // You need to specify at least 1 page for removal
        if len < 5 {
            fmt.Print("Usage: ./editpdf remove [output.pdf] [input.pdf] [page",
                "-num...]\n")
            os.Exit(0)
        }
        // Remove pages
        err := pdf.RemovePagesFile(os.Args[2], os.Args[1], os.Args[3:], nil)
        if err != nil {
            fmt.Print("REMOVE FAILED: ", err)
            os.Exit(1)
        }
        // Removal Successful
        fmt.Print("Successfully remove page(s)!\nOutput -> ", os.Args[2])
    case "help":
        // Info on using Merge
        fmt.Print("Merging PDFs?\nTo merge any number of PDFs type the ",
            "following command.\n\n./editpdf merge [output.pdf] [input-1.pdf ",
            " input-2.pdf ...]\n\nExample: ./editpdf merge full-menu.pdf ",
            "day-menu.pdf night-menu.pdf drink-menu.pdf\n(NOTE: PDFs are ",
            "combined in the order they are typed.)\n\n")
        // Info on using remove
        fmt.Print("Removing pages from PDF?\nTo remove pages from a PDF type ",
            "the following command.\n\n./editpdf remove [output.pdf] ",
            "[input.pdf] [page-num ...]\n\nExample: ./editpdf remove ",
            "no-cover.pdf book.pdf 1 2 3\n(The following will remove the 1st, ",
            "2nd, and 3rd pages)\n")
        os.Exit(0)
    default:
        fmt.Print("Please specify \"merge\" or \"remove\" or type \"./editpdf",
            " help\" for help\n")
        os.Exit(0)
    }
}
