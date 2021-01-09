/**
 * VMX - Virtual Machine Xlator for nand2tetris
 *
 * See textbook chapters 7 and 8 for details.
 */

import java.io.File

// define symbols for command types
enum class CType {
  C_NONE, C_ARITHMETIC, C_PUSH, C_POP, C_LABEL, C_GOTO,
  C_IF, C_FUNCTION, C_RETURN, C_CALL
}

//data class Command(var type: CType, var arg0: String, var arg1: String)


class Hello : CliktCommand() {
  val count: Int by option(help="Number of greetings").int().default(1)
  val name: String by option(help="The person to greet").prompt("Your name")

  override fun run() {
      repeat(count) {
          echo("Hello $name!")
      }
  }
}

fun main(args: Array<String>) = Hello().main(args)

fun main2(args: Array<String>) {

  if (args.isEmpty()) {
    println("usage: vmtrans [<file>.vm | directory/]")
    return
  }

  // // we only create one CodeWriter object
  // val cw = CodeWriter()

  // // create a parser for each argument, handling directories correctly
  // for (path in args) {
  //   var file = File(path);

  //   if (file.isDirectory()) {
  //     file
  //       .walk()
  //       .filter { it.isFile }
  //       .filter { it.extension == "vm" }
  //       .forEach { parseFile(it.canonicalPath, cw) }
  //   } else {
  //     parseFile(path, cw)
  //   }
  // }

}

// fun parseFile(path: String, cw: CodeWriter) {
//   //val parser = Parser(path)
//   cw.setFileName(path)
//   File(path).forEachLine {
//     cw.writeComment(it)
//     // val cmd = VMCommand(it)
//     // parser.command = it
//     // val (ctype, arg1, arg2) = parseCommand()
//     // when (ctype) {
//     //   C_PUSH, C_POP -> cw.writePushPop(ctype, arg1, arg2)
//     //   else -> {
//     //     cw.writeComment(it)
//     //   }
//     // }
//   }
// }


// class VMCommand(ctype: CType = C_NONE, arg0: String, arg1: String) {
//   constructor(line: String) {
//     // strip comments from end of line
//     ctype = C
//   }
// }


// class Parser(path: String) {

//   command: String = ""

//   fun commandType() {
//   }

//   fun arg1() {
//   }

//   fun arg2() {
//   }
// }

// class CodeWriter() {

//   fun setFileName(name: String) {
//     writeComment("Parsing file: ${name}")
//   }

//   fun writeComment(comment: String = "") {
//     println(if (comment.isBlank()) "" else "// ${comment}")
//   }


//   fun writeArithmetic() {
//   }

//   fun writePushPop(command: String, segment: String, index: String) {
// // push puts the value at @SP then increments it
// // pop removes the value at @SP then decrements it


//   }

//   fun close() {
//   }
// }
