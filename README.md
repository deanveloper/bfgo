# BFGO
A Brainfuck interpreter written in [Go].

## What is Brainfuck?
Brainfuck is an esoteric language, meaning that it is not meant for things like readability, but
really only exist *because it can*.

By concept, brainfuck uses an "infinite tape" as its memory space. Brainfuck can only point to
a single spot on the tape (which we will call "the byte"), and change the value at that spot by +/- 1. Brainfuck also has some
mediocre input and output functionality.

There are only 8 characters that make up Brainfuck programs. They are as follows:

| Char | Description |
| ---- | ---- |
| `<` | Moves the pointer on the tape a single space to the left |
| `>` | Moves the pointer on the tape a single space to the right |
| `+` | Adds one to "the byte" |
| `-` | Subtracts one from "the byte" |
| `[` | Skips to the corresponding `]` if "the byte" is zero |
| `]` | Backs up to the corresponding `[` if "the byte" is not zero |
| `,` | Reads a single byte of input and stores it in the byte |
| `.` | Writes the byte as ASCII to the output |

I have exposed this as both a library and a command line interface.

----
# Library Usage
Usage for the library is simple. For basic brainfuck programs, import
`github.com/deanveloper/bfgo` (remember to `go get` first) and call `bfgo.Run(codeBytes)`.
There are some nifty settings you can include though, which can be used
with `bfgo.RunWithSettings(codeBytes, settings)`.

## Settings
| Setting | Description | Default |
| ------- | ----------- | ------- |
| EOFNoChange | Decides if `,` should change a cell or not on EOF. Overrides `EOFDefault` | `false` |
| EOFDefault | Decides what `,` should set a cell to on EOF. | `10 (\n)` |
| KeepCR | Whether to keep the CR (`\r`) in CRLF (`\r\n`) line breaks | `false` |
| InitialArraySize | The initial size of the array (tape), which expands to the right when needed | `false` |
| Input | Where to take input from | `os.Stdin` |
| Output | Where to send output to | `os.Stdout` |

# Command Line Usage
Installation: `go get github.com/deanveloper/bfgo/bfgo-cli`

Now that you have it installed, you can run any code with `bfgo-cli <code>`.
On the other hand, you can also use some very nifty options as well!

## Options
| Flag | Description | Default |
| ---- | ----------- | ------- |
| `-n` `-eofnochange` | Decides if `,` should change a cell or not on EOF, overrides `-d` | `false` |
| `-d` `-eofdefault` | Decides what `,` should set a cell to on EOF | `10 (\n)` |
| `-c` `-keepcr` | Whether to keep the CR (`\r`) in CRLF (`\r\n`) line breaks | `false` |
| `-s` `-initialarrsize` | The initial size of the array (tape), which expands to the right when needed | `30` |
| `-i` `-input` | Input source. "stdin" for cli input, `!` for `<bf>!input` input, otherwise a filename | `stdin` |
| `-o` `-output` | Output destination. "stdout" for cli output, otherwise a filename. | `stdout` |

[Go]: https://golang.org/
