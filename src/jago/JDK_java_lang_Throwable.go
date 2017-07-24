package jago

import (
	"strings"
	"strconv"
)

func register_java_lang_Throwable() {
	register("java/lang/Throwable.fillInStackTrace(I)Ljava/lang/Throwable;", Java_jang_lang_Throwable_fillInStackTrace)
}

func Java_jang_lang_Throwable_fillInStackTrace(this Reference, dummy Int) Reference {
	thread := THREAD_MANAGER.currentThread

	size := len(thread.vmStack) - ___exceptionHierarchy(this.Class()) // skip how many frames
	stacktrace := make([]string, size)


	for i, frame := range thread.vmStack[:size] {
		javaClassName := strings.Replace(frame.method.class.name, "/", ".", -1)
		stacktrace[size-1-i] = javaClassName + "." + frame.method.name + __getSourceFileAndLineNumber(frame)
	}

	this.SetExtra(stacktrace)
	return this
}

func ___exceptionHierarchy(class *Class) int {
	distance := 1
	for c := class.superClass; c != nil; c = c.superClass {
		distance++
	}
	return distance
}

func __getSourceFileAndLineNumber(frame *Frame) string {
	sourceFile := frame.method.class.sourceFile
	if sourceFile == "" {
		sourceFile = "<Unknow>"
	}
	lineNumber := -1
	lineNumbers := frame.method.lineNumbers
	for i := len(lineNumbers)-1; i >= 0; i-- {
		entry := lineNumbers[i]
		if frame.pc >= entry.startPc {
			lineNumber = int(entry.lineNumber)
		}
	}

	linenum := ""
	if lineNumber >= 0 {
		linenum = strconv.FormatInt(int64(lineNumber), 10)
	}
	return "(" + sourceFile + ":" + linenum + ")"
}
