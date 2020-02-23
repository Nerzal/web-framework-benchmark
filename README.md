# Web Framework Benchmark

## Requirement

- `dep`

## Installation

- `go get github.com/Nerzal/web-framework-benchmark`

## Running Benchmark

```sh
go test -bench=.
```


## Results
Comparison between atreugo, echo and gin. 

### Hello
Path /hello returns statusCode 200 and string "hello"

### Static
Path /static returns statusCode 200 and this README.md file


```c#
Benchmark_Atreugo_Hello-8    	   75698	     15847 ns/op	       8 B/op	       1 allocs/op
Benchmark_Echo_Hello-8       	   38376	     31598 ns/op	    2286 B/op	      21 allocs/op
Benchmark_Gin_Hello-8        	   36249	     33155 ns/op	    2269 B/op	      20 allocs/op
Benchmark_Atreugo_Static-8   	   41910	     29420 ns/op	    1374 B/op	       6 allocs/op
Benchmark_Echo_Static-8      	   23901	     52619 ns/op	    4817 B/op	      32 allocs/op
Benchmark_Gin_Static-8       	   23750	     49489 ns/op	    4887 B/op	      36 allocs/op
```