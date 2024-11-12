// push constant 3030
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1

// pop pointer 0
@SP
AM=M-1
D=M
@3
M=D

// push constant 3040
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1

// pop pointer 1
@SP
AM=M-1
D=M
@4
M=D

// push constant 32
@32
D=A
@SP
A=M
M=D
@SP
M=M+1

// pop this 2
@2
D=A
@THIS
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D

// push constant 46
@46
D=A
@SP
A=M
M=D
@SP
M=M+1

// pop that 6
@6
D=A
@THAT
D=D+M
@R15
M=D
@SP
AM=M-1
D=M
@R15
A=M
M=D

// push pointer 0
@3
D=M
@SP
A=M
M=D
@SP
M=M+1

// push pointer 1
@4
D=M
@SP
A=M
M=D
@SP
M=M+1

// add
@SP
M=M-1
A=M
D=M
A=A-1
M=D+M

// push this 2
@2
D=A
@THIS
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

// sub
@SP
M=M-1
A=M
D=M
A=A-1
M=M-D

// push that 6
@6
D=A
@THAT
A=M
A=A+D
D=M
@SP
A=M
M=D
@SP
M=M+1

// add
@SP
M=M-1
A=M
D=M
A=A-1
M=D+M

