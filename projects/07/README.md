


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

1. First thing: false = 0, true = -1. This makes sense as in 2's complement -1
   is all bits "on".

2. Second thing: the `%[1]d` indicates where a distinguishing integer is
   interpolated into the generated assembly code, this gives our labels unique
   names and prevents collisions.

3. Third thing: comparison operators operate in this "sense" on stack values.
   If the stack contains `(2, 3)` the operation `LT` results in `true`, and
   leaves the stack containing `(-1)` (which is "true").

```
// eq
@SP             // A=0          // address the stack
M=M-1           // M[0]--       // shorten the stack
A=M             // A=M[0]       // dereference SP into A
D=M             // D=M[A]       // save the top stack value in D
A=A-1           // A--          // address the previous stack value
D=D-M           // D=D-M[A]     // compare

@EQ_TRUE_123    // A=EQ_TRUE_123
D; JEQ          // if D=0, jump to EQ_TRUE_123

// if we get here, D != 0 (we're in the FALSE branch)
@SP             // A=0
A=M             // A=M[0]       // dereference SP, save into A
A=A-1           // A--          // reference the data value
M=0             // M[A]=0       // M[A]=false

@EQ_CONT_123    // unconditional jump to CONTINUE
0; JEQ

// if we get here, D == 0 (we're in the TRUE branch)
(EQ_TRUE_123)   // label
@SP             // A=0
A=M             // A=M[0]
A=A-1           // A--          // reference the data value
M=-1            // M[A]=-1      // M[A]=true

(EQ_CONTINUE_123)
```

========================================================

```
// lt
@SP             // A=0          // address the stack
M=M-1           // M[0]--       // shorten the stack by 1
A=M             // A=M[0]       // dereference
D=M             // D=M[A]       // save the top stack value in D
A=A-1           // A--          // address the next stack value
D=M-D           // D=M[A]-D     // compare

// If the stack contains (3, 5), we want the stored value to be TRUE,
// since 3<5. At this point, D=3-5=-2, so we jump to the TRUE branch if D<0.

@LT_TRUE_123    // address branch LT_TRUE_123
D; JLT          // if D<0, branch to LT_TRUE_123

// if we get here, D>=0 (we're in the FALSE branch)
@SP             // A=0          // address the stack
A=M             // A=M[0]       // dereference
A=A-1           // A--          // reference the data value
M=0             // M[A]=0       // 

@EQ_CONT_123
0; JEQ

(EQ_TRUE_123
@SP
A=M
M=-1

(EQ_CONTINUE_123)
```


```
// gt
@SP             // A=0          // address the stack
M=M-1           // M[0]--       // shorten the stack by 1
A=M             // A=M[0]       // dereference
D=M             // D=M[A]       // save the top stack value in D
A=A-1           // A--          // address the next stack value
D=M-D           // D=M[A]-D     // compare

// If the stack contains (3, 5), we want the stored value to be TRUE,
// since 3<5. At this point, D=3-5=-2, so we jump to the TRUE branch if D<0.

@LT_TRUE_123    // address branch LT_TRUE_123
D; JLT          // if D<0, branch to LT_TRUE_123

// FIXME this looks wrong
// if we get here, D>=0 (we're in the FALSE branch)
@SP             // A=0          // address the stack
A=M             // A=M[0]       // dereference
M=0             // M[A]=0       //-
@EQ_CONT_123
0; JEQ

(EQ_TRUE_123
@SP
A=M
M=-1

(EQ_CONTINUE_123)
```


=============================================================

## Assembler for `push constant 7`

1. D=1
2. A=M[SP]      // ==M[0]
3. M=D
4. increment SP

```
// push constant 7:
//   1. D = 7
//   2. A = M[0]
//   3. M[A] = D
//   4. M[0]++
@7
D=A
@SP
A=M
M=D
@SP
M=M+1
```

