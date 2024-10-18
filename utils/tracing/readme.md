## Tracing

Tracing allows you to trace part of the code and view the
execution path then the code run completes

To trace eg a function import this package then add at the top of the function as below:

```
import (
    	"github.com/judewood/bakery/utils/tracing"
)	

func MyFunc() {
	f := tracing.StartTrace()
	defer tracing.StopTrace(f)
    ...
```

When the program completes a trace.out file will be created
To view this file run the following in the command line
 go tool trace trace.out this will open the trace [output file](http://127.0.0.1:63843/) in your browser

 From here you can select options to see more trace detail.
 There is a guide [here](https://blog.gopheracademy.com/advent-2017/go-execution-tracer/)
