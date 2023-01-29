# uniques

## Install
```shell
go get github.com/dvvedz/uniques@latest
```

## Usage 
```go
// Remove duplicates
myIntSlice := []int{1,1,2,5,4}
strSliceUnique := uniques.IntSlice(myIntSlice)
// [1, 2, 5, 4]


myStrSlice := []string{"AA", "a", "a", "c", "b"}
strSliceUnique := uniques.StringSlice(myStrSlice)
// ["AA", "a", "c", "b"]

```
