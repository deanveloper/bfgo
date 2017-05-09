package bfgo

import (
    "io"
    "os"
)

// Settings for RunWithSettings
type Settings struct {
    // Whether EOF on input should change the current cell or not
    EOFNoChange bool
    // Default byte when EOF is run. Overridden by EOFChange.
    EOFDefault byte

    // Whether to keep carriage returns ('\r') in inputs
    KeepCR bool

    // Initial array size
    InitialArraySize uint64

    // Where to take input from
    Input io.Reader

    // Where to send output to
    Output io.Writer
}

// Initializes Settings with the following defaults:
// EOF -> '\n'
// KeepCR -> false
// InitialArraySize -> 30,000
// Input -> os.Stdin
// Output -> os.Stdout
func DefaultSettings() *Settings {
    return &Settings{
        EOFNoChange:      false,
        EOFDefault:       10,
        KeepCR:           false,
        InitialArraySize: 30,
        Input:            os.Stdin,
        Output:           os.Stdout,
    }
}

// Runs given BF code with default settings defined by DefaultSettings()
func Run(code []byte) {
    RunWithSettings(code, DefaultSettings())
}

// Runs given BF code with the given settings
func RunWithSettings(code []byte, settings *Settings) {
    tape := make([]byte, settings.InitialArraySize)
    tapeIndex := 0
    for i := 0; i < len(code); i++ {
        switch code[i] {
        case byte('+'):
            tape[tapeIndex]++
        case byte('-'):
            tape[tapeIndex]--
        case byte('<'):
            tapeIndex--
        case byte('>'):
            tapeIndex++
            if len(tape) <= tapeIndex {
                tape = append(tape, 0)
            }
        case byte('['):
            if tape[tapeIndex] == 0 {
                for bracks := 1; bracks > 0; {
                    i++
                    switch code[i] {
                    case byte('['):
                        bracks++
                    case byte(']'):
                        bracks--
                    }
                }
            }
        case byte(']'):
            if tape[tapeIndex] != 0 {
                for bracks := -1; bracks < 0; {
                    i--
                    switch code[i] {
                    case byte('['):
                        bracks++
                    case byte(']'):
                        bracks--
                    }
                }
            }
        case byte('.'):
            slice := tape[tapeIndex:tapeIndex+1]
            settings.Output.Write(slice)
        case byte(','):
            slice := make([]byte, 1)
            _, err := settings.Input.Read(slice)
            if err == io.EOF {
                if settings.EOFNoChange {
                    slice[0] = tape[tapeIndex]
                } else {
                    slice[0] = settings.EOFDefault
                }
            }
            tape[tapeIndex] = slice[0]
        }
    }
}
