package main

import (
    "github.com/deanveloper/bfgo"
    flags "github.com/jessevdk/go-flags"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

type options struct {
    EOFNoChange    bool `short:"n" long:"eofnochange" description:"Decides if \",\" should change a cell on EOF. Overrides -d"`
    EOFDefault     byte `short:"d" long:"eofdefault" default:"10" description:"Decides what \",\" should set a cell to on EOF."`
    KeepCR         bool `short:"c" long:"keepcr" description:"Whether to keep the CR (\"\r\") in CRLF (\"\r\n\") line breaks"`
    InitialArrSize uint64 `short:"s" long:"initialarrsize" default:"30" description:"The initial size of the array (tape), which expands to the right when needed"`
    Input          string `short:"i" long:"input" default:"stdin" description:"Input source. \"stdin\" for cli input, \"!\" for \"<bf>!input\" input, otherwise a filename"`
    Output         string `short:"o" long:"output" default:"stdout" description:"Output destination. \"stdout\" for cli output, otherwise a filename."`
}

func main() {
    log.SetFlags(0)
    opts := options{}
    parser := flags.NewParser(&opts, flags.Default)
    parser.ArgsRequired = true
    parser.Usage = "bfgo-cli <bf | file.bf> [OPTIONS]"
    args, err := parser.Parse()
    if err != nil {
        if err.(*flags.Error).Type == flags.ErrHelp {
            os.Exit(1)
        }
        log.Fatalln("Error:", err)
    }

    if len(args) != 1 {
        log.Fatalln("Error: Not enough arguments! Use bfgo-cli -h for help.")
    }

    var bf []byte
    if strings.HasSuffix(args[0], ".b") || strings.HasSuffix(args[0], ".bf") {
        var err error
        bf, err = ioutil.ReadFile(args[0])
        if err != nil {
            log.Fatalln("Error:", err)
        }
    } else {
        bf = []byte(args[0])
    }

    var input io.Reader
    var output io.Writer

    if opts.Input == "stdin" {
        input = os.Stdin
    } else if opts.Input == "!" {
        temp := strings.SplitN(string(bf), "!", 2)[1]
        input = strings.NewReader(temp)
    } else {
        file, err := os.Open(opts.Input)
        if err != nil {
            log.Fatalln("Error:", err)
        }
        input = file
    }

    if opts.Output == "stdout" {
        output = os.Stdout
    } else {
        file, err := os.Open(opts.Output)
        if err != nil {
            log.Fatalln("Error:", err)
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
