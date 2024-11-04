package main

import "fmt"

// CComp lists the possible C-instruction computations.
var CComp = map[string]uint8{
	"0":   0b0101010,
        "1":   0b0111111,
        "-1":  0b0111010,
        "D":   0b0001100,
        "A":   0b0110000,
        "!D":  0b0001101,
        "!A":  0b0110001,
        "-D":  0b0001111,
        "-A":  0b0110011,
        "D+1": 0b0011111,
        "A+1": 0b0110111,
        "D-1": 0b0001110,
        "A-1": 0b0110010,
        "D+A": 0b0000010,
        "D-A": 0b0010011,
        "A-D": 0b0000111,
        "D&A": 0b0000000,
        "D|A": 0b0010101,
        "M":   0b1110000,
        "!M":  0b1110001,
        "-M":  0b1110011,
        "M+1": 0b1110111,
        "M-1": 0b1110010,
        "D+M": 0b1000010,
        "D-M": 0b1010011,
        "M-D": 0b1000111,
        "D&M": 0b1000000,
        "D|M": 0b1010101,
}

// CJump lists the C-instruction jump encodings
var CJump = map[string]uint8{
        "":    0, // 000
        "JGT": 1, // 001
        "JEQ": 2, // 010
        "JGE": 3, // 011
        "JLT": 4, // 100
        "JNE": 5, // 101
        "JLE": 6, // 110
        "JMP": 7, // 111
}

// calculate the destination bits from the `dest` part of the C instruction
func dest(str string) string {
	var dbits uint8 = 0
        if strings.Contains(str, "M") {
                dbits |= 1
        }
        if strings.Contains(str, "D") {
                dbits |= 2
        }
        if strings.Contains(str, "A") {
                dbits |= 4
        }
	return fmt.Sprintf("%03b", dbits)
}

// look up the compute bits from the `comp` part of the C instruction
func comp(str string) string {
        // calculate computation bits
        if cbits, ok := CComp[str]; ok {
	    return fmt.Sprintf("%07b", cbits)
	}
        log.Fatal().Msgf("error finding comp bits for %v", comp)
}

// look up the jump bits from the `jump` part of the c instruction
func jump(str string) string {
	if jbits, ok := CJump[jump]; ok {
	    return fmt.Sprintf("%03b", jbits)
	}
        log.Fatal().Msgf("error finding jump bits for %v", jump)
}
