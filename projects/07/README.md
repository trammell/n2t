


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

1. Store value "7" in D
2. Store D in the current stack pointer address
3. increment the stack pointer

```asm
@7          // A=7
D=A         // D=7
@SP         // A=0
A=M         // A=M[0]
M=D         // M[A]=7
@SP         // A=0
M=M+1       // M[0]++
```

=============================================================

## Assembler for `push local 1`

This pushes the second value in the local segment onto the stack.

```
// push local 1
@1          // A=1
D=A         // D=1
@LCL        // A=1
A=M         // A=M[1]   :: A is now the memory address of the local stack
A=A+D       // A is now the memory address of local 1
D=M         // D = local 1 (finally!)
@SP         // A=0
A=M         // A=M[0]   :: A is now the memory address of the stack pointer
M=D         // M[A]=D   :: push local 1 on to the stack
@SP         // A=0
M=M+1       // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `push argument 1`

This pushes the second value in the argument segment on to the stack. Almost
identical to `push local 1`.

```
// push argument 1
@1          // A=1
D=A         // D=1
@ARG        // A=2
A=M         // A=M[2]   :: A is now the memory address of the argument stack
A=A+D       // A is now the memory address of argument 1
D=M         // D = ARG[1]
@SP         // A=0
A=M         // A=M[0]   :: A is now the memory address of the stack pointer
M=D         // M[A]=D   :: push argument 1 on to the stack
@SP         // A=0
M=M+1       // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `push this 1`

```
// push this 1
@1          // A=1
D=A         // D=1
@THIS       // A=3
A=M         // A=M[3]   :: A is now the memory address of the THIS segment
A=A+D       // A is now the memory address of this[1]
D=M         // D = this[1] (finally!)
@SP         // A=0
A=M         // A=M[0]   :: A is now the memory address of the stack pointer
M=D         // M[A]=D   :: push this[1] on to the stack
@SP         // A=0
M=M+1       // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `push that 1`

```
// push that 1
@1          // A=1
D=A         // D=1
@THAT       // A=4
A=M         // A=M[1]   :: A is now the memory address of the THAT segment
A=A+D       // A is now the memory address of THAT[1]
D=M         // D = THAT[1]
@SP         // A=0
A=M         // A=M[0]   :: A is now the memory address of the stack pointer
M=D         // M[A]=D   :: push THAT[1] on to the stack
@SP         // A=0
M=M+1       // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `push pointer 1`

Push the value of `pointer 1` on to the stack

From TEoCS Chapter 7, page 142:

> The `pointer` segment is mapped on RAM locations 3-4 (also called `THIS` and
> `THAT`) ... Thus access to `pointer i` should be translated to assembly code
> that addresses RAM location `3 + i` ...

Another example is on page 138, VM code for



Also, it appears that valid `pointer`s are only 0 and 1, so I've added that
restriction to the code also.

```
// push pointer 1
@4          // A=4      // calculated in code (index + 3)
D=M         // D=M[4]   // D=POINTER[1]
@SP         // A=0
A=M         // A=M[0]   // A = stack pointer address
M=D         // M[A]=D   // push POINTER[1] on to the stack
@SP         // A=0
M=M+1       // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `push static 1`

Push the 1-th value in the static segment to the stack.
See documentation on TEoCS page 143.

* When a new symbol is encountered for the first time by the assembler, the
  assembler allocates a new RAM address for it, starting at address 16.
* Exploit this by representing variable J in file F as symbol `F.J`. 
* Depends on the filename!

```
// push static 5
@filename.5         // address the previously defined value
D=M                 // D = RAM[filename.5]
@SP                 // A=0
A=M                 // A=M[0]   // A = stack pointer address
M=D                 // M[A]=D   // push RAM[filename.5] onto the stack
@SP                 // A=0
M=M+1               // M[0]++   :: increment the stack pointer
```

=============================================================

## Assembler for `pop local 1`




=============================================================

## Assembler for `pop argument 1`

=============================================================

## Assembler for `pop this 1`

=============================================================

## Assembler for `pop that 1`

=============================================================

## Assembler for `pop local 1`




=============================================================

## Assembler for `pop argument 1`

=============================================================

## Assembler for `pop this 1`

=============================================================

## Assembler for `pop that 1`
