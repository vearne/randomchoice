## randomchoice
从长度为N的数组中，随机挑出M个元素


### Install:
```
go get -u github.com/vearne/randomchoice
```

### Import:
```
import "github.com/vearne/randomchoice"
```

### Quick Start
```
package main

import (
    rc "github.com/vearne/randomchoice"
    "fmt"
)

type Car struct{
    color string
    name  string
}

func (c *Car) String() string{
    return c.color + "-" + c.name
}

func main(){
    // example 1
    var children []string
    children = []string{"lily", "rose","lisa"}
    // randomly select 2 kids from children
    idxSlice := rc.RandomChoice(len(children), 2)
    result := make([]string, 0, 2)
    for _, v := range idxSlice{
        result = append(result, children[v])
    }
    // result: selected kids
    fmt.Println(result)

    // example 2
    var carSlice []*Car= make([]*Car, 0, 3)
    bmw := Car{color:"black", name:"bmw"}
    carSlice = append(carSlice, &bmw)
    buick := Car{color:"silvery", name:"buick"}
    carSlice = append(carSlice, &buick )
    skoda := Car{color:"white", name:"skoda"}
    carSlice = append(carSlice, &skoda )
    // random select 2 kinds from carSlice
    idxSlice = rc.RandomChoice(len(carSlice), 2)
    cars := make([]*Car, 0, 2)
    for _, v := range idxSlice{
        cars = append(cars, carSlice[v])
    }
    fmt.Println(cars)
}
```
