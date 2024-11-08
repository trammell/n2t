


## Assembler for unary operations like `neg`

```
// SP holds the address of the current top of the stack,
// which is *one past* the top value on the stack.
@SP         // address the stack pointer
A=M         // get the current address SP points to
A=A-1       // point A to the value we want to negate
M=-M        // negate it
```

## Assembler for binary operations like `add`

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

* `sub` is the same except the final operation is M=M-D

## Assembler for comparison operations like `eq`

> Note that the `%[1]d` indicates where a distinguishing integer is injected
> into the generated assembly code, this gives our labels unique names and
> prevents collisions.


```
// eq
@SP             // address the stack
M=M-1           // shorten the stack by 1
A=M             // fetch the address of the top stack value
D=M             // save the top stack value in D
A=A-1           // address the underlying stack value
D=D-M           // subtract / test for equality

@EQ_TRUE_123  // if D=0, branch to EQ_TRUE
D; JEQ

// if we get here, D != 0 (we're in the FALSE branch)
@SP
A=M
M=0
@EQ_CONT_123
0; JEQ

(EQ_TRUE_123
@SP
A=M
M=-1

(EQ_CONTINUE_123)
```


## Assembler for `push constant 1`

1. D=1
2. A=M[SP]      // ==M[0]
3. M=D
4. increment SP

```
// push constant 1
@1
D=A
@SP        // *SP = D
A=M
M=D
@SP        // SP++
M=M+1
```

