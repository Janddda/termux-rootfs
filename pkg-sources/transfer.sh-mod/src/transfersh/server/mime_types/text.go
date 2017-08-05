/*
The MIT License (MIT)

Copyright (c) 2017 Leonid Plyushch <leonid.plyushch@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package mime_types

import (
	"mime"
)

func initTextMime() {
	mime.AddExtensionType(".323", "text/h323")
	mime.AddExtensionType(".asc", "text/plain")
	mime.AddExtensionType(".asm", "text/x-asm")
	mime.AddExtensionType(".bas", "text/plain")
	mime.AddExtensionType(".cbl", "text/x-cobol")
	mime.AddExtensionType(".cc", "text/x-c++src")
	mime.AddExtensionType(".cfg", "text/plain")
	mime.AddExtensionType(".cmake", "text/x-cmake")
	mime.AddExtensionType(".cob", "text/x-cobol")
	mime.AddExtensionType(".config", "text/plain")
	mime.AddExtensionType(".conf", "text/plain")
	mime.AddExtensionType(".cpp", "text/x-c++src")
	mime.AddExtensionType(".cp", "text/x-c++src")
	mime.AddExtensionType(".csh", "text/x-shellscript")
	mime.AddExtensionType(".cs", "text/x-csharp")
	mime.AddExtensionType(".csv", "text/csv")
	mime.AddExtensionType(".c--", "text/plain")
	mime.AddExtensionType(".c", "text/x-csrc")
	mime.AddExtensionType(".cxx", "text/x-c++src")
	mime.AddExtensionType(".diff", "text/x-patch")
	mime.AddExtensionType(".erl", "text/x-erlang")
	mime.AddExtensionType(".etx", "text/x-setext")
	mime.AddExtensionType(".f90", "text/x-fortran")
	mime.AddExtensionType(".f95", "text/x-fortran")
	mime.AddExtensionType(".for", "text/x-fortran")
	mime.AddExtensionType(".f", "text/x-fortran")
	mime.AddExtensionType(".go", "text/x-go")
	mime.AddExtensionType(".hh", "text/x-c++hdr")
	mime.AddExtensionType(".hpp", "text/x-c++hdr")
	mime.AddExtensionType(".hp", "text/x-c++hdr")
	mime.AddExtensionType(".htc", "text/x-component")
	mime.AddExtensionType(".h", "text/x-chdr")
	mime.AddExtensionType(".h++", "text/x-c++hdr")
	mime.AddExtensionType(".htt", "text/webviewhtml")
	mime.AddExtensionType(".hxx", "text/x-c++hdr")
	mime.AddExtensionType(".ics", "text/calendar")
	mime.AddExtensionType(".inf", "text/plain")
	mime.AddExtensionType(".ini", "text/plain")
	mime.AddExtensionType(".jad", "text/vnd.sun.j2me.app-descriptor")
	mime.AddExtensionType(".java", "text/x-java")
	mime.AddExtensionType(".log", "text/x-log")
	mime.AddExtensionType(".lua", "text/x-lua")
	mime.AddExtensionType(".markdown", "text/markdown")
	mime.AddExtensionType(".md", "text/markdown")
	mime.AddExtensionType(".mkd", "text/markdown")
	mime.AddExtensionType(".mk", "text/x-makefile")
	mime.AddExtensionType(".mli", "text/x-ocaml")
	mime.AddExtensionType(".ml", "text/x-ocaml")
	mime.AddExtensionType(".mrl", "text/x-mrml")
	mime.AddExtensionType(".m", "text/x-objcsrc")
	mime.AddExtensionType(".pas", "text/x-pascal")
	mime.AddExtensionType(".patch", "text/x-patch")
	mime.AddExtensionType(".pl", "text/x-perl")
	mime.AddExtensionType(".pm", "text/x-perl")
	mime.AddExtensionType(".pod", "text/x-pod")
	mime.AddExtensionType(".po", "text/x-gettext-translation")
	mime.AddExtensionType(".pot", "text/x-gettext-translation-template")
	mime.AddExtensionType(".p", "text/x-pascal")
	mime.AddExtensionType(".py", "text/x-python")
	mime.AddExtensionType(".pyx", "text/x-python")
	mime.AddExtensionType(".qmlproject", "text/x-qml")
	mime.AddExtensionType(".qml", "text/x-qml")
	mime.AddExtensionType(".qmltypes", "text/x-qml")
	mime.AddExtensionType(".readme", "text/x-readme")
	mime.AddExtensionType(".reg", "text/x-ms-regedit")
	mime.AddExtensionType(".rej", "text/x-reject")
	mime.AddExtensionType(".rs", "text/rust")
	mime.AddExtensionType(".scala", "text/x-scala")
	mime.AddExtensionType(".scm", "text/x-script.scheme")
	mime.AddExtensionType(".sct", "text/scriptlet")
	mime.AddExtensionType(".sgml", "text/sgml")
	mime.AddExtensionType(".sgm", "text/sgml")
	mime.AddExtensionType(".sh", "text/x-shellscript")
	mime.AddExtensionType(".shtml", "text/html")
	mime.AddExtensionType(".smali", "text/plain")
	mime.AddExtensionType(".sql", "text/x-sql")
	mime.AddExtensionType(".s", "text/x-asm")
	mime.AddExtensionType(".stm", "text/html")
	mime.AddExtensionType(".tk", "text/x-tcl")
	mime.AddExtensionType(".tsv", "text/tsv")
	mime.AddExtensionType(".txt", "text/plain")
	mime.AddExtensionType(".uls", "text/iuls")
	mime.AddExtensionType(".uue", "text/x-uuencode")
	mime.AddExtensionType(".vala", "text/x-vala")
	mime.AddExtensionType(".vapi", "text/x-vala")
	mime.AddExtensionType(".vcf", "text/x-vcard")
	mime.AddExtensionType(".vhdl", "text/x-vhdl")
	mime.AddExtensionType(".vhd", "text/x-vhdl")
	mime.AddExtensionType(".v", "text/x-verilog")
	mime.AddExtensionType(".wsgi", "text/x-python")
	mime.AddExtensionType(".xhtml", "text/html")
	mime.AddExtensionType(".xmi", "text/x-xmi")
	mime.AddExtensionType(".xml", "text/xml")
	mime.AddExtensionType(".xsl", "text/xsl")
	mime.AddExtensionType(".zsh", "text/x-shellscript")
}
