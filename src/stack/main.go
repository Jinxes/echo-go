package stack

import "github.com/go-martini/martini"
import "Index"

func Start() {
  m := martini.Classic()

  m.Get("/", Index)

  m.Run()
}
