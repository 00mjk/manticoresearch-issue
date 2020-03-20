# manticoresearch-issue

## Requirements
| dependency                                                   | mode   | version                                                            |
| :----------------------------------------------------------- | :----: | :----------------------------------------------------------------- |
| [`Go`](https://golang.org/doc/install)                       | Local  | `v1.13` _([go modules](https://blog.golang.org/using-go-modules))_ |
| [`Manticoresearch`](https://manticoresearch.com/downloads/)  | Local  | `3.3.0`															 |

## Steps
```
$ export GO111MODULE=on
$ searchd --config ./config/manticore.conf
$ go run golang/go-sdk/main.go --index
$ go run golang/go-sdk/main.go --search --query Fulda      # here is happen the bug, cannot search the index
$ go run golang/go-sdk/main.go --search --query AllClimate # here is happen the bug, cannot search the index
$ php php/test.php                                         # works
```

## Bug Description
```bash
panic: runtime error: index out of range [-2]

goroutine 1 [running]:
github.com/manticoresoftware/go-sdk/manticore.(*QueryResult).parseMatch(0xc0000ee000, 0xc0000a8e00, 0xc0000b01e0)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/search.go:901 +0xd48
github.com/manticoresoftware/go-sdk/manticore.(*QueryResult).parseResult(0xc0000ee000, 0xc0000b01e0, 0x1, 0xc0000ee000)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/search.go:819 +0x2a1
github.com/manticoresoftware/go-sdk/manticore.parseSearchAnswer.func1(0xc0000b01e0, 0xc0000b0121, 0xc0000e0000)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/search.go:946 +0x74
github.com/manticoresoftware/go-sdk/manticore.(*Client).netQuery(0xc0000d0000, 0xc0000b0000, 0xc0000b01c0, 0xc0000ba2f0, 0xc0000d97ac, 0xc0000d97c0, 0x108cf54, 0xc000110000)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/client.go:225 +0x1ca
github.com/manticoresoftware/go-sdk/manticore.(*Client).RunQuery(0xc0000d0000, 0x1400000000, 0x3e8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x116dda8, ...)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/manticore.go:320 +0xc6
github.com/manticoresoftware/go-sdk/manticore.(*Client).Query(0xc0000d0000, 0x116e15d, 0x5, 0xc0000d9dc0, 0x1, 0x1, 0x3, 0x2460, 0x11a07a0)
	/Users/lucmichalski/go/pkg/mod/github.com/manticoresoftware/go-sdk@v0.0.0-20191205035816-0e8dbffac2c9/manticore/manticore.go:267 +0x141
main.main()
	/Users/lucmichalski/go/src/github.com/lucmichalski/manticoresearch-issue/golang/go-sdk/main.go:65 +0x96f
exit status 2
```