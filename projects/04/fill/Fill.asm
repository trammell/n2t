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


// initialize symbol R0 (black) to 65535 (0xffff)
@32767
D=A
D=D+A
D=D+1
@R0
M=D

// initialize symbol R1 (white) to 0
@R1
M=0

// initialize symbol R2 (color) to white (R1)
@R1
D=M
@R2
M=D

// initialize symbol R3 (address) to @SCREEN
@SCREEN
D=A
@R3
M=D

// outer infinite loop
(OUTERLOOP)

// testing with just first two rows

// if address > 24575, then @addr=@SCREEN
@R3
D=M
@24575
D=D-A
@ADDR_RANGE_OK
D;JLE   // jump if @addr <= 24575
@SCREEN
D=A
@R3
M=D
(ADDR_RANGE_OK)

// if @KBD is true then set @color to @black
// else set @color to @white
@KBD
D=M
@KEYPRESS_TRUE
D; JGT
@WHITE
D=M
@color
M=D
@KEYPRESS_END
0; JMP
(KEYPRESS_TRUE)
@black
D=M
@color
M=D
(KEYPRESS_END)

// set @addr to @color
@color
D=M
@R3
M=D

// increment @addr
@R3
M=M+1

// infinite loop
@OUTERLOOP
0;JMP

