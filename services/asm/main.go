package main

import (
	"log"
	"os"
)

// main function: takes no arguments, reads source file name from os.Args
// Read the instructions, resolve symbols, and emit the assembled code
func main() {

	// read source file into `lines`
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("os.Open: %s", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err() != nil {
		log.Fatalf("scanner.Scan: %s", err)
	}

	// 2. compile instructions one at a time
	// 3. ???
	// 4. profit

	a := NewAssembler()
	for _, instruction in = a.Assemble(lines) {


	}
	a.Emit()



        out = a.SymbolTable.String()
        for _, inst := range a.Instructions {
                out += fmt.Sprintln(inst)
        }
        return
}



}
