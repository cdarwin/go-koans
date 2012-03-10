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

var __ interface{}

func TestKoans(t *testing.T) {
  testBasics()
  fmt.Printf("%c[32m", 27)
}

func assert(o bool) {
  if !o {
    fmt.Printf("%c[31m", 27)
    fmt.Printf("%s\n\n", __getRecentLine())
    os.Exit(1)
  }
}

func __getRecentLine() string {
  _, file, line, _ := runtime.Caller(2)
  buf, _ := ioutil.ReadFile(file)
  code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
  return fmt.Sprintf("\n%v:%d\n%s", path.Base(file), line, code)
}
