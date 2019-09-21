// vim: set ai et ts=4 :
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

// the screen has 256 rows and 512 columns.
// The address of the first 16-bit word is @SCREEN (16384)
// Each row has 32 16-bit words.
// There are 8192 words in total for the screen.

// pseudocode:
//    initalize @black=65535, @white=0, @row=0, @col=0
//    initalize @color=@white
//    while true:
//        if @col > 32, set @col=0 and @row=@row+1
//        if @row > 256, set @row=0
//        read @kbd
//        if @kbd is true, set @color=@black
//            else set @color=@white
//        calculate @addr = @screen + 32 * @row + @col
//        set M[@addr] to @color
//        @col = @col + 1

// set symbol @white to 0
@white
M=0

// set symbol @black to 65535 (0xffff)
@32767
D=A
D=D+A
D=D+1
@black
M=D

// set symbols @row and @col to 0
@row
M=0
@col
M=0

// set symbol @word to 0
@word
M=0

// increment @word
@word
M=M+1

// set @SCREEN[@word] = @color
@SCREEN
M=D

@SCREEN
A=A+1
M=D


// KBD is 0 when no key is pressed
// if KBD is 0, then set R0 to R1
//     else set R0 to R2

@KBD

//D=D+1
//A=D
//M=1


//    @R2
//    M=0
//
//
//    @SCREEN
//    D=A
//    @R0
//    M=0
//
//(LOOP)
//
//    // jump to end if R0 is at the end of the screen (@SCREEN + 8192)
//    @R0
//    D=M
//    @END
//    D;JEQ
//
//    // R2 = R2 + R1
//    @R1
//    D=M
//    @R2
//    M=D+M
//
//    // R0 = R0 - 1
//    @R0
//    M=M-1
//
//    // jump to top of loop
//    @LOOP
//    0;JMP
//

(END)
    @END
    0;JMP
