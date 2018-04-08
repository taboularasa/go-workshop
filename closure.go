package main

func counterFactory(start, increment int) (() -> (int)) {
  counter := start
  increment := increment
  return func() int {
    counter = counter + increment
    return counter
  }
}

countBy10 := counterFactory(0, 10)

fmt.Println(countBy10())
fmt.Println(countBy10())
fmt.Println(countBy10())
fmt.Println(countBy10())
