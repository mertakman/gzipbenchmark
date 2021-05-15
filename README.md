Benchmark results 
Compression:
BenchmarkGZipWithStdLib-12                	      34	  33124084 ns/op	40885242 B/op	     865 allocs/op
BenchmarkGZipWithCompress-12              	      62	  20770934 ns/op	70122738 B/op	     760 allocs/op
BenchmarkGZipWithCompressPool-12          	      88	  13676932 ns/op	   73864 B/op	       5 allocs/op

Decompression:
BenchmarkDecompressWithStdLib-12          	      24	  49165638 ns/op	493592968 B/op	   74999 allocs/op
BenchmarkDecompressWithCompress-12        	      61	  18535577 ns/op	 3572307 B/op	   50556 allocs/op
BenchmarkDecompressWithCompressPool-12    	      55	  18233096 ns/op	 1679221 B/op	   47905 allocs/op