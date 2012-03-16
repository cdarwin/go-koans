package go_koans

import (
  "testing"
  "os"
  "fmt"
  "runtime"
  "io/ioutil"
  "path"
  "strings"
)

var __string__ string = "impossibly lame value"
var __int__ int = -1
var __byte__ byte = 255
var __bool__ bool = false // ugh
var __float32__ float32 = -1.0

func TestKoans(t *testing.T) {
  aboutNumbers()
  aboutStrings()
  aboutArrays()
  aboutSlices()
  aboutControlFlow()
  aboutEnumeration()
  aboutAnonymousFunctions()
  aboutVariadicFunctions()
  aboutFiles()
  aboutInterfaces()
  aboutMaps()
  aboutPointers()
  aboutStructs()
  aboutAllocation()

  // TODO: ie, gameplan
  //aboutGoroutines()
  //aboutChannels()
  //aboutPanics()

  fmt.Printf("\n%c[32;1mYou won life. Good job.\n\n", 27)
}

func assert(o bool) {
  if !o {
    fmt.Printf("\n%c[35m%s\n\n", 27, __getRecentLine())
    os.Exit(1)
  }
}

func __getRecentLine() string {
  _, file, line, _ := runtime.Caller(2)
  buf, _ := ioutil.ReadFile(file)
  code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
