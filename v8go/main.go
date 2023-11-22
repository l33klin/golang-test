package main

import (
	"fmt"
	"runtime"
	"time"
)
import v8 "rogchap.com/v8go"

func main() {
	fetchData()
}

func timeElapsed() {
	printVMUsage()
	start := time.Now()
	ctx := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	elapsed := time.Since(start)
	fmt.Printf("NewContext：%s\n", elapsed)
	printVMUsage()
	ctx.RunScript("const add = (a, b) => a + b", "math.js") // executes a script on the global context
	ctx.RunScript(`const result = add(3, 4)`, "main.js")    // any functions previously added to the context can be called
	_, err := ctx.RunScript(`let jsonStr = {a:1, b:2}`, "main.js")
	if err != nil {
		fmt.Println(err)
	} // any functions previously added to the context can be called
	val, err := ctx.RunScript("result", "value.js") // return a value in JavaScript back to Go
	if err != nil {
		fmt.Println(err)
	}
	js, err := ctx.RunScript("JSON.stringify(jsonStr)", "value2.js") // return a value in JavaScript back to Go
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("addition result: %s\n", val)
	fmt.Printf("jsonStr: %s\n", js)

	elapsed = time.Since(start)
	fmt.Printf("RunScript：%s\n", elapsed)
	printVMUsage()

	ctx2 := v8.NewContext(ctx.Isolate())
	globalObj := ctx2.Global()
	err = globalObj.Set("jsonStr", `{"c":3, "d":4}`)
	if err != nil {
		fmt.Println(err)
	}
	js2, err := ctx2.RunScript("JSON.stringify(JSON.parse(jsonStr))", "ctx2_get_value.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("jsonStr in ctx2: %s\n", js2)

	ctx3 := v8.NewContext(ctx2.Isolate())
	js3, err := ctx3.RunScript("JSON.stringify(JSON.parse(jsonStr))", "ctx3_get_value.js")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("jsonStr in ctx3: %s\n", js3)

	ctx3.RunScript("console.log(12345)", "ctx3_log.js")

	elapsed = time.Since(start)
	fmt.Printf("NewContext with VM：%s\n", elapsed)
	printVMUsage()
}

func callbackDemo() {
	iso := v8.NewIsolate() // create a new VM
	// a template that represents a JS function
	printfn := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		fmt.Printf("%v", info.Args()) // when the JS function is called this Go callback will execute
		return nil                    // you can return a value back to the JS caller if required
	})
	global := v8.NewObjectTemplate(iso)       // a template that represents a JS Object
	global.Set("print", printfn)              // sets the "print" property of the Object to our function
	ctx := v8.NewContext(iso, global)         // new Context with the global Object set to our object template
	ctx.RunScript("print('foo')", "print.js") // will execute the Go callback with a single argunent 'foo'
}

func printVMUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("当前进程使用的内存大小：%d 字节\n", m.Sys)
}

func fetchData() {
	jsStr := `
const options = {
  hostname: 'icanhazdadjoke.com',
  path: '/',
  method: 'GET'
};

const req = https.request(options, (res) => {
  let data = '';

  res.on('data', (chunk) => {
    data += chunk;
  });

  res.on('end', () => {
    joke=data;
  });
});

req.on('error', (error) => {
  console.error(error);
});

req.end();
`
	ctx := v8.NewContext()
	_, err := ctx.RunScript(jsStr, "main.js")
	if err != nil {
		fmt.Println(err)
	}
	value, err := ctx.RunScript("joke", "main.js")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}
