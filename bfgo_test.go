package bfgo

import (
    "bytes"
    "strings"
    "testing"
)

var(
    // A solution I created, handles both EOF and \n case
    REVERSE1 = []byte(">+[[-]++++++++++,[->+>+<<]>>[-<<+>>]<----------]<<[.<]")
    // A solution I found online, does not handle EOF case
    REVERSE2 = []byte("+[->,----------]<[+++++++++++.<]")
    // A solution I found online which does not handle \n case
    REVERSE_EOF_NOCHANGE = []byte(">,[>,]<[.<]")

    // A cool BF interpreter written in BF!
    BF_INTERPRETER = []byte(`
    >>>+[[-]>>[-]++>+>+++++++[<++++>>++<-]++>>+>+>+++++[>++>++++++<<-]+>>>,<++[[>[
    ->>]<[>>]<<-]<[<]<+>>[>]>[<+>-[[<+>-]>]<[[[-]<]++<-[<+++++++++>[<->-]>>]>>]]<<
    ]<]<[[<]>[[>]>>[>>]+[<<]<[<]<+>>-]>[>]+[->>]<<<<[[<<]<[<]+<<[+>+<<-[>-->+<<-[>
    +<[>>+<<-]]]>[<+>-]<]++>>-->[>]>>[>>]]<<[>>+<[[<]<]>[[<<]<[<]+[-<+>>-[<<+>++>-
    [<->[<<+>>-]]]<[>+<-]>]>[>]>]>[>>]>>]<<[>>+>>+>>]<<[->>>>>>>>]<<[>.>>>>>>>]<<[
    >->>>>>]<<[>,>>>]<<[>+>]<<[+<<]<]
    [input a brainfuck program and its input, separated by an exclamation point.
    Daniel B Cristofani (cristofdathevanetdotcom)
    http://www.hevanet.com/cristofd/brainfuck/]`)
)


func TestRun(t *testing.T) {

    t.Run("Reverse String 1", func(t *testing.T) {
        in := strings.NewReader("cool stuff dude")
        out := bytes.NewBufferString("")
        RunWithSettings(REVERSE1, &Settings{
            EOFDefault:       10,
            InitialArraySize: 50,
            Input:            in,
            Output:           out,
        })
        expected := "edud ffuts looc"
        actual := out.String()
        if expected != actual {
            t.Errorf("Expected %q, actual %q", expected, actual)
        }
    })

    t.Run("Reverse String 2", func(t *testing.T) {
        in := strings.NewReader("cool stuff dude")
        out := bytes.NewBufferString("")
        RunWithSettings(REVERSE2, &Settings{
            EOFDefault:10,
            InitialArraySize:50,
            Input:in,
            Output:out,
        })
        expected := "edud ffuts looc"
        actual := out.String()
        if expected != actual {
            t.Errorf("Expected %q, actual %q", expected, actual)
        }
    })

    t.Run("Reverse String with nochange EOF", func(t *testing.T) {
        in := strings.NewReader("cool stuff dude")
        out := bytes.NewBufferString("")
        RunWithSettings(REVERSE_EOF_NOCHANGE, &Settings{
            EOFNoChange:      true,
            InitialArraySize: 50,
            Input:            in,
            Output:           out,
        })
        expected := "edud ffuts looc"
        actual := out.String()
        if expected != actual {
            t.Errorf("Expected %q, actual %q", expected, actual)
        }
    })

    // BF interprets a BF interpreter that interprets a string reverser
    t.Run("Brainfuck Interpreter Interpreter", func(t *testing.T) {
        in := strings.NewReader(string(BF_INTERPRETER) + "!" + string(REVERSE_EOF_NOCHANGE) + "!cool stuff dude")
        out := bytes.NewBufferString("")
        RunWithSettings(BF_INTERPRETER, &Settings{
            EOFNoChange:true,
            InitialArraySize:30000,
            Input:in,
            Output:out,
        })
        expected := "edud ffuts looc"
        actual := out.String()
        if expected != actual {
            t.Errorf("Expected %q, actual %q", expected, actual)
        }
    })
}