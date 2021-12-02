package main

import (
  "fmt"
)

type Person struct {
  FirstName string
  Age       int
}

func main() {
  // 初期値を与えて作成
  a := Person{
    FirstName: "John",
    Age:       20,
  }
  fmt.Println(a)
  fmt.Printf("%#v\n", a)

  // 実態なのでaの値は変更されない
  aa := a
  aa.Age = 25
  fmt.Printf("a.Age = %v\n", a.Age)
  fmt.Printf("aa.Age = %v\n", aa.Age)

  // 参照型でコピーするとaの値が変更される
  aaa := &a
  aaa.Age = 27
  fmt.Printf("a.Age = %v\n", a.Age)
  fmt.Printf("aaa.Age = %v\n", aaa.Age)

  // new関数で初期化
  var b = new(Person)
  b.FirstName = "Paul"
  b.Age = 30
  fmt.Println(b)
  fmt.Printf("%#v\n", b)

  // 参照型のためbの値が変更される
  bb := b
  bb.Age = 35
  fmt.Printf("b.Age = %v\n", b.Age)
  fmt.Printf("bb.Age = %v\n", bb.Age)

  // 値のコピーなのでbの値は変更されない
  bbb := *b
  bbb.Age = 20
  fmt.Printf("b.Age = %v\n", b.Age)
  fmt.Printf("bbb.Age = %v\n", bbb.Age)

  // make関数で初期化
  // 配列のスライス、map、チャネル専用
  c := make(map[string]int)
  c["John"] = 20
  c["Paul"] = 30
  fmt.Println(c)
  fmt.Printf("%#v\n", c)

  // structの別の作り方
  var d struct {
    FirstName string
    Age       int
  }
  d.FirstName = "John"
  d.Age = 20
  fmt.Println(d)
  fmt.Printf("%#v\n", d)
}