package go_koans

import "io/ioutil"
import "strings"

func testFiles() {
  filename := "about_files.go"
  contents, _ := ioutil.ReadFile(filename)
  lines := strings.Split(string(contents), "\n")
  assert(lines[5] == __string__) // handling files is too trivial
}
