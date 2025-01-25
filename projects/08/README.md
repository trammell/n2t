# README for n2t Chapter 8



```
// This is a slick bit of assembly that:
//  1. decrements the stack pointer
//  2. sets D to the top value in the stack
@SP         // A=0
AM=M-1      // (A,M[0]) = (M[0]-1, M[0]-1)
D=M         // D = previous top stack value
```




