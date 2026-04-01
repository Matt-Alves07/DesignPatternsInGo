package main

import "fmt"

type Person struct {
  FirstName, MiddleName, LastName string
}

func (p *Person) Names() []string {
  return []string{p.FirstName, p.MiddleName, p.LastName}
}

func (p *Person) NamesGenerator() <-chan string {
  out := make(chan string)
  go func() {
    defer close(out)
    out <- p.FirstName
    if (len(p.MiddleName) > 0) {
      out <- p.MiddleName
    }
    out <- p.LastName
  }()
  return out
}

type PersonNameIterator struct {
  person *Person
  current  int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
  return &PersonNameIterator{person, -1}
}

func (p *PersonNameIterator) MoveNext() bool {
  p.current++
  return p.current < len(p.person.Names())
}

func (p *PersonNameIterator) Current() string {
  switch p.current {
  case 0:
    return p.person.FirstName
  case 1:
    return p.person.MiddleName
  case 2:
    return p.person.LastName
  }
  panic("Invalid state")
}

func main() {
  // p := Person{"Alexander", "Graham", "Bell"}
  p := Person{"Alexander", "", "Bell"}
  for it := NewPersonNameIterator(&p); it.MoveNext(); {
    fmt.Println(it.Current())
  }
}
