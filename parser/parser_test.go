package parser

import "testing"


func TestBase(t *testing.T) {
  name := "Jake"

  if name != "Jake" {
    t.Fatal("WRONG NAME??????")
  }
}
