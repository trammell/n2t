package lib


// Remove all comments from an instruction
func stripComments(inst string) string {
	return regexp.MustCompile(`//.*`).ReplaceAllString(inst, "")
}



