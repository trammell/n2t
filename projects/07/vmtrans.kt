/**
 * VMX - Virtual Machine Xlator for nand2tetris
 *
 * See textbook chapters 7 and 8 for details.
 */

import java.io.File

enum class CType {
  C_NONE, C_ARITHMETIC, C_PUSH, C_POP, C_LABEL, C_GOTO,
  C_IF, C_FUNCTION, C_RETURN, C_CALL
}

fun main(args: Array<String>) {

  if (args.isEmpty()) {
    println("usage: vmtrans [<file>.vm | directory/]")
    return
  }

  // we only create one CodeWriter object
  val cw = CodeWriter()

  // create a parser for each argument, handling directories correctly
  for (path in args) {
    var file = File(path);

    if (file.isDirectory()) {
      file
        .walk()
        .filter { it.isFile }
        .filter { it.extension == "vm" }
        .forEach { parseFile(it.canonicalPath, cw) }
    } else {
      parseFile(path, cw)
    }
  }

}

fun parseFile(path: String, cw: CodeWriter) {
  val parser = Parser(path)
  cw.setFileName(path)
  File(path).forEachLine {
    parser.command = it
    when (parser.commandType()) {
      C_PUSH -> cw.writePushPop( C_PUSH, parser.arg1())
      else -> {
        cw.writeComment(it)
      }
    }
  }
}

class Parser {

  command: String = ""

  fun commandType() {
  }

  fun arg1() {
  }

  fun arg2() {
  }
}

class CodeWriter() {

  fun setFileName(name: String) {
    writeComment("Parsing file: ${name}")
  }

  fun writeComment(comment: String = "") {
    println(if (comment.isBlank()) "" else "// ${comment}")
  }


  fun writeArithmetic() {
  }

  fun writePushPop() {
  }

  fun close() {
  }
}
