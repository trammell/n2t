// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/05/Control.tst

load Control.hdl,
output-file Control.out,
compare-to Control.cmp,
output-list inst%X1.4.1 iBit%B3.1.2 vBus%X1.4.1 aBit%B3.1.2
            cBus%B1.6.1 dBus%B2.3.1 jBus%B2.3.1;

set inst %B0111111111111111,
eval,
output;

set inst %B1110000000000000,
eval,
output;

set inst %B1110100000100100,
eval,
output;

set inst %B1110010000010010,
eval,
output;

set inst %B1110001000001001,
eval,
output;

