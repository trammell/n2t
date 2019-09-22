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

// Notes on program speed
// ======================
// I am having a hard time getting this to run fast enough to show the desired
// behavior, so I modified the requirements to instead scan through the SCREEN
// buffer, putting black pixels when any key is pressed, and white pixels
// otherwise.

// notes on SCREEN layout
// ======================
// There are 256 rows.
// Each row has 512 pixels (32 16-bit words)
// 256 rows x 32 words / row = 8192 words in RAM (8192 = 0x2000)
//
// SCREEN RAM looks like this:
//         16384 (0x4000)   start of top row
//         16385 (0x4001)   second word in top row
//         16386 (0x4002)
//         ...
//         16415 (0x401f)   end of top row
//         16416 (0x4020)   start of second row
//         16417 (0x4021)   second word in second row
//         ...
//         24574 (0x5ffe)
//         24575 (0x5fff)
//
// KBD RAM is just one word at 24576 (0x6000)

// Variables used:
//    @R0 is a constant, 65535 (0xffff), denoting the color "black"
//    @R1 is a constant, 0, denoting the color "white"
//    @R2 contains the current selected color, either white or black
//    @R3 contains the current screen RAM address being updated

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

// initialize symbol R3 (addr) to @SCREEN
@SCREEN
D=A
@R3
M=D

// start of infinite loop
(INFINITE_LOOP)

// make sure we don't pass our loop boundary
// if @addr >= 24576, then @addr := @SCREEN (production code)
// if @addr >= 16416, then @addr := @SCREEN (test code)
@R3
D=M
@24576
// @16416   // uncomment this line for shorter test loop
D=D-A
@ADDR_RANGE_OK
D;JLT   // jump if @addr is in bounds (does not change @addr)
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
@R1
D=M
@R2
M=D
@KEYPRESS_END
0; JMP
(KEYPRESS_TRUE)
@R0
D=M
@R2
M=D
(KEYPRESS_END)

// Apply @color to the current SCREEN RAM address. Need to dereference the
// contents of R3, which is an address somewhere in the screen buffer.
@R2
D=M
@R3     // A := 3
A=M     // A := M[3]
M=D     // M[3] := D

// increment @addr
@R3
M=M+1

// go back to top of infinite loop
@INFINITE_LOOP
0;JMP

