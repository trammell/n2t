


## Assembly for `add` function

```
// SP holds the address of the current top of the stack,
// which is *one past* the top value on the stack.
@SP         // address the stack pointer
M=M-1       // decrement the stack pointer
A=M         // get the current address SP points to
D=M         // fetch the top of the stack into D
A=A-1       // address the new stack top
M=D+M       // save the new value
```



