// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/05/Control.tst

load Control.hdl,
output-file Control.out,
compare-to Control.cmp,
output-list inst%B1.16.1 isAddress%B5.1.5 isCompute%B5.1.5;

set inst %X0000,
eval,
output;

set inst %Xe000,
eval,
output;

