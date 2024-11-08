package main


	// construct output filehandle
	out, err := os.Create(outfname)
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer out.Close()

	// translate all lines in all source files
	for _, srcFileName := range getSourceFiles(file) {
		log.Info().Msgf(`Translating source file: "%s"`, file)
		srcFile, err := os.Open(srcFileName)
		if err != nil {
			fmt.Println(err)
		}

		scanner := bufio.NewScanner(srcFile)
		for scanner.Scan() {
			//fmt.Println(`// ` + scanner.Text())
			cmd := NewCommand(scanner.Text())
			out.WriteString(cmd.GetAsm())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		// p := NewParser(fn)
		// for p.Scan() {
		// 	//cw.writeComment(p.Text())
		// 	fmt.Println(p.Scanner.Text())
		// 	switch p.commandType() {
		// 	case C_ARITHMETIC:
		// 		cw.writeComment("arithmetic: " + p.Text())
		// 		log.Info().Msgf(`looks like arithmetic`)
		// 		cw.writeArithmetic(p.Text())
		// 	case C_POP, C_PUSH:
		// 		cw.writeComment("pushpop: " + p.Text())
		// 		log.Info().Msgf(`looks like pushpop`)
		// 		cw.writePushPop(p.commandType(), p.arg1(), p.arg2())
		// 	default:
		// 		cw.writeComment("unknown: " + p.Text())
		// 		log.Info().Msgf(`unknown type`)
		// 	}
		// }
	}
}
