// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
M=M-1
A=M
D=M
A=A-1
D=D-M
@EQ_TRUE_1
D; JEQ
@SP
A=M
A=A-1
M=0
@EQ_CONT_1
0; JEQ
(EQ_TRUE_1)
@SP
A=M
A=A-1
M=-1
(EQ_CONT_1)

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
M=M-1
A=M
D=M
A=A-1
D=D-M
@EQ_TRUE_2
D; JEQ
@SP
A=M
A=A-1
M=0
@EQ_CONT_2
0; JEQ
(EQ_TRUE_2)
@SP
A=M
A=A-1
M=-1
(EQ_CONT_2)

// push constant 16
@16
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 17
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
M=M-1
A=M
D=M
A=A-1
D=D-M
@EQ_TRUE_3
D; JEQ
@SP
A=M
A=A-1
M=0
@EQ_CONT_3
0; JEQ
(EQ_TRUE_3)
@SP
A=M
A=A-1
M=-1
(EQ_CONT_3)

// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@LT_TRUE_4
D; JLT
@SP
A=M
A=A-1
M=0
@LT_CONT_4
0; JEQ
(LT_TRUE_4)
@SP
A=M
A=A-1
M=-1
(LT_CONT_4)

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 892
@892
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@LT_TRUE_5
D; JLT
@SP
A=M
A=A-1
M=0
@LT_CONT_5
0; JEQ
(LT_TRUE_5)
@SP
A=M
A=A-1
M=-1
(LT_CONT_5)

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 891
@891
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@LT_TRUE_6
D; JLT
@SP
A=M
A=A-1
M=0
@LT_CONT_6
0; JEQ
(LT_TRUE_6)
@SP
A=M
A=A-1
M=-1
(LT_CONT_6)

// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@GT_TRUE_7
D; JGT
@SP
A=M
A=A-1
M=0
@GT_CONT_7
0; JEQ
(GT_TRUE_7)
@SP
A=M
A=A-1
M=-1
(GT_CONT_7)

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32767
@32767
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@GT_TRUE_8
D; JGT
@SP
A=M
A=A-1
M=0
@GT_CONT_8
0; JEQ
(GT_TRUE_8)
@SP
A=M
A=A-1
M=-1
(GT_CONT_8)

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 32766
@32766
D=A
@SP
A=M
M=D
@SP
M=M+1

// gt
@SP
M=M-1
A=M
D=M
A=A-1
D=M-D
@GT_TRUE_9
D; JGT
@SP
A=M
A=A-1
M=0
@GT_CONT_9
0; JEQ
(GT_TRUE_9)
@SP
A=M
A=A-1
M=-1
(GT_CONT_9)

// push constant 57
@57
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 31
@31
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 53
@53
D=A
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

// push constant 112
@112
D=A
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

// neg
@SP
A=M
A=A-1
M=-M

// and
@SP
M=M-1
A=M
D=M
A=A-1
M=M&D

// push constant 82
@82
D=A
@SP
A=M
M=D
@SP
M=M+1

// or
@SP
M=M-1
A=M
D=M
A=A-1
M=M|D

// not
@SP
A=M
A=A-1
M=!M

