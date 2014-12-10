package go_koans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	__string__       string  = "impossibly lame value"
	__int__          int     = -1
	__positive_int__ int     = 42
	__byte__         byte    = 255
	__bool__         bool    = false // ugh
	__boolean__      bool    = true  // oh well
	__float32__      float32 = -1.0
	__delete_me__    bool    = false
)

var __runner__ runner = nil

func TestKoans(t *testing.T) {
	aboutBasics()
	aboutStrings()
	aboutArrays()
	aboutSlices()
	aboutTypes()
	aboutControlFlow()
	aboutEnumeration()
	aboutAnonymousFunctions()
	aboutVariadicFunctions()
	aboutFiles()
	aboutInterfaces()
	aboutCommonInterfaces()
	aboutMaps()
	aboutPointers()
	aboutStructs()
	aboutAllocation()
	aboutChannels()
	aboutConcurrency()
	aboutPanics()

	fmt.Printf("\n%c[32;1mYou won life. Good job.\n\n", 27)
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
