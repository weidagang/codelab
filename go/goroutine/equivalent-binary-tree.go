// https://tour.golang.org/concurrency/8

package main

import (
  "fmt"
  "golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  _walk(t, ch)
  close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
  if t == nil {
      return
  }
  _walk(t.Left, ch)
  ch <- t.Value
    _walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for v1 := range ch1 {
    v2, ok := <-ch2
    if !ok {
      return false
    }
    if v1 != v2 {
        return false
    }
  }
  
  // ch1 has finished, but ch2 has more.
  _, ok := <-ch2
  if ok {
    return false
  }

  return true
}

func main() {
  isSame := Same(tree.New(3), tree.New(3))
  fmt.Println("Is the same: ", isSame)
}
