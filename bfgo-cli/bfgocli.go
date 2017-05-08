package main

import (
    "flag"
    "fmt"
    "github.com/deanveloper/bfgo"
    "io"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    noChange := flag.Bool("eofnc", false, "Decides if the cell should change on input if input is EOF. True means no change.")
    eofDefault := flag.Uint("eof", 10, "What should be considered as input on EOF. Overridden if eofnc=true")
    keepCr := flag.Bool("keepcr", false, "Whether the CR in CRLF input lines should be kept.")
    initialArrSize := flag.Uint("init", 30, "Initial size of the tape used")
    inputS := flag.String("in", "stdin", "Input source. \"stdin\" for cli input, \"!\" for BF input, otherwise name of a file.")
    outputS := flag.String("out", "stdout", "Output destination. \"stdout\" for cli output, otherwise a filename.")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s <bf | file.bf> [flags...]\n", os.Args[0])
        fmt.Fprintln(os.Stderr, "  First arg may be brainfuck code, or a file if it ends with .b or .bf")
        flag.PrintDefaults()
    }

    flag.Parse()

    if flag.NArg() != 1 {
        flag.Usage()
        return
    }

    var bf []byte
    if strings.HasSuffix(flag.Arg(0), ".b") || strings.HasSuffix(flag.Arg(0), ".bf") {
        var err error
        bf, err = ioutil.ReadFile(flag.Arg(0))
        if err != nil {
            fmt.Println("Error: " + err.Error())
            return
        }
    } else {
        bf = []byte(flag.Arg(0))
    }

    var input io.Reader
    var output io.Writer

    if *inputS == "stdin" {
        input = os.Stdin
    } else {
        file, err := os.Open(*inputS)
        if err != nil {
            fmt.Println("Error: " + err.Error())
            return
        }
        input = file
    }

    if *outputS == "stdout" {
        output = os.Stdout
    } else {
        file, err := os.Open(*outputS)
        if err != nil {
            fmt.Println("Error: " + err.Error())
            return
        }
        output = file
    }
    settings := &bfgo.Settings{
        EOFNoChange: *noChange,
        EOFDefault: byte(*eofDefault),
        KeepCR: *keepCr,
        InitialArraySize: uint64(*initialArrSize),
        Input: input,
        Output: output,
    }

    bfgo.RunWithSettings(bf, settings)
}
