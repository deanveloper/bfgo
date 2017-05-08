package main

import (
    "flag"
    "fmt"
    "github.com/deanveloper/bfgo"
    flags "github.com/jessevdk/go-flags"
    "io"
    "io/ioutil"
    "os"
    "strings"
)

type options struct {
    EOFNoChange    bool `short:"n" long:"eofnochange" default:"true" description:"Decides if \",\" should change a cell on EOF. Overrides -d"`
    EOFDefault     byte `short:"d" long:"eofdefault" default:"10" description:"Decides what \",\" should set a cell to on EOF."`
    KeepCR         bool `short:"c" long:"keepcr" default:"false" description:"Decides if CR should be kept in CRLF linebreaks"`
    InitialArrSize uint64 `short:"s" long:"initialarrsize" default:"30" description:"Initial size for the tape"`
    Input          string `short:"i" long:"input" default:"stdin" description:"Input source. \"stdin\" for cli input, \"!\" for BF input, otherwise name of a file."`
    Output         string `short:"o" long:"output" default:"stdout" description:"Output destination. \"stdout\" for cli output, otherwise a filename."`
}

func main() {
    opts := options{}
    args, err := flags.Parse(&opts)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %s\n", err)
    }

    var bf []byte
    if strings.HasSuffix(args[0], ".b") || strings.HasSuffix(args[0], ".bf") {
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

    if opts.Input == "stdin" {
        input = os.Stdin
    } else if opts.Input == "!" {
        temp := strings.SplitN(string(bf), "!", 2)[0]
        input = strings.NewReader(temp)
    } else {
        file, err := os.Open(opts.Input)
        if err != nil {
            fmt.Println("Error: " + err.Error())
            return
        }
        input = file
    }

    if opts.Output == "" || opts.Output == "stdout" {
        output = os.Stdout
    } else {
        file, err := os.Open(opts.Output)
        if err != nil {
            fmt.Println("Error: " + err.Error())
            return
        }
        output = file
    }
    settings := &bfgo.Settings{
        EOFNoChange:      opts.EOFNoChange,
        EOFDefault:       opts.EOFDefault,
        KeepCR:           opts.KeepCR,
        InitialArraySize: opts.InitialArrSize,
        Input:            input,
        Output:           output,
    }

    bfgo.RunWithSettings(bf, settings)
}
