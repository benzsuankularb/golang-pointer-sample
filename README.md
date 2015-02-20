#Example show how pointer works in golang
## How to
~~~
go get github.com/benzsuankularb/golang-pointer
~~~
To run,
~~~
go test
~~~
## Method explanation
- Method1 : Create struct and passing to function with a pointer type will no any copy of variable.
- Method2 : Create struct and passing to function with a non pointer type, struct will be copied
- Method3 : Create struct and passing to function with a pointer but return without pointer type, sturct will be copied.
- Method4 : Pointer type have it's own address. Pointer value hold an address.
- Method5 : Copy a struct will copy all the variable of struct, pointer also been copied but value of these pointer was the same.
- Method6 : Copy struct which have sub-struct, All variable in struct and sub-struct also be cloned.