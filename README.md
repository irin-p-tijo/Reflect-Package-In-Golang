#Golang Using Reflect Package


Reflection is an advanced topic in Golang.This will teach you about Go's reflection in detail

What is reflection?
Reflection is a way to inspect and manipulate the structure and values of variables at runtime. It allows you to analyze the structure of a variable, including its type, fields, methods, and other metadata.

Mainly there two purposes for this:

1.Helps to obtain the type and value of a variable.
2.Modifying the value of a variable

The package contains two main types: reflect.Type and reflect.Value. reflect.Type represents the type of a variable, and reflect.Value represents the value of a variable.

```

package main

import (
    "fmt"
    "reflect"
)

func main(){
    var x int =5
    y:=reflect.TypeOf(x)
    z:=reflect.ValueOf(x)

    fmt.Println("the type of x is : ",y,the value of x is :",z)
}

```

the program output is 
```
the type of x is int the value of x is 5 

```


if we want to modify the value of the variable we want to use reflect.Value .


```

package main


import(
    "fmt"
    "reflect"
)

func main(){
    var x int = 5

    v:=reflect.ValueOf(&x).Elem()
    v.SetInt(20)

    fmt.Println("new value is :",x)
}
```
out put is :
```
new value is :20

```


type conversion using reflect package

```
package main

impoet (
    "fmt"
    "reflect"
)


func TypeConversion(value interface{})float64{
    v:=reflect.ValueOf(value)

    if v.Kind()!=reflect.Int{
        panic("there is diffferenc in datatype)
    }

    i:=v.Int()
    return float64(i)
}

func main(){
var x int = 29

f:=Typeconversion(x)
fmt.Println("Integer is :",x)
fmt.Prinln("converted to :",f)

}

```


Conversion of struct to Map

``` 
package main

import (
    "fmt"
    "reflection"
)


type Person struct{
    Name string
    Age int
}

func StructToMap(s interface{})map[string]interface{}{
    result:=make(map[string]interface{})

    v:=reflect.ValueOf(s)
    if v.Kind()!= reflect.Struct{
        panic("input is not struct")
    }

    t:=v.Type()


    for i:=0;i<v.NumField();i++{
        field:=t.Field(i)
        value:=v.Field(i)

        result[field.Name]=value.Interface()
    }
    return result
}

func main(){
    p:=Person{Name:"Alice",Age:18}
    m:=StructToMap(p)

    fmt.Println(p)
    fmt.Println(m)
}

```