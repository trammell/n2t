// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

    // initialize screen to white
    // loop forever over screen addresses?
    // in loop, if any key is pressed put black pixel
    // otherwise put white pixel


    // need to use array syntax


    @R2
    M=0


    @SCREEN
    D=A
    @R0
    M=0

(LOOP)

    // jump to end if R0 is at the end of the screen (@SCREEN + 8192)
    @R0
    D=M
    @END
    D;JEQ

    // R2 = R2 + R1
    @R1
    D=M
    @R2
    M=D+M

    // R0 = R0 - 1
    @R0
    M=M-1

    // jump to top of loop
    @LOOP
    0;JMP

(END)
    @END
    0;JMP



