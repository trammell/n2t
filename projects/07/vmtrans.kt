/**
 * VMX - Virtual Machine Xlator for nand2tetris
 *
 * See textbook chapters 7 and 8 for details.
 */

import java.io.InputStreamReader
import java.io.OutputStreamWriter

fun main(args: Array<String>) {
  println("Hello VMTranslator World " + args[0])
  if (args.isEmpty()) {
    print("Please add some command line arguments")
    return
  }
  
  // we only create one CodeWriter object
  val cw = CodeWriter(System.out.writer())

  // create a parser for each argument
  for (arg in args) {
    val parser = Parser(arg)
    cw.writeComment("starting parse of file: ${arg}")

    cw.writeComment()
  }
// print("Hello ${args[0]}")
  // if it's a file, read from the file
  // if it's a directory, run on all files within
}

enum class CType {
  C_NONE, C_ARITHMETIC, C_PUSH, C_POP, C_LABEL, C_GOTO,
  C_IF, C_FUNCTION, C_RETURN, C_CALL
}

class Parser(val filename: String) {

  fun hasMoreCommands() {
  }

  fun advance() {
  }

  fun commandType() {
  }

  fun arg1() {
  }

  fun arg2() {
  }
}

class CodeWriter(val ostream: OutputStreamWriter) {

  fun setFileName() {
  }

  fun writeComment() {
    println()
  }

  fun writeComment(comment: String) {
    println("// " + comment)
  }

  fun writeArithmetic() {
  }

  fun writePushPop() {
  }

  fun close() {

  }
}
