[
    I am very proud of this brainfuck program. It reverses any string that is inputted.
    I did some online research, there are a couple other shorter string reversal BF algorithms.
    They are both listed in bfgo_test.go.
]
>+[                       loop (add 1 so the loop executes; keep #0 = 0)
  [-]++++++++++           set cell to \n (to handle EOF case)
  ,                       take input
  [->+>+<<]               copy into next two cells
  >>                      go to 2nd copy
  [-<<+>>]                move value to original cell; makes "origin" & "copy"
  <----- -----            subtract 10 from copy

  note: copy's value is not relevant; all that matters is that if it is 0
  (aka if the char was a newline) that it will break out of the loop!
  this means that we can overwrite the copy with the next character; allowing
  for the string to be stored neatly on the tape
]

<<        go to cell before newline (twice because we are currently on the newlines copy)
[.<]      keep outputting until the 0th cell is reached
