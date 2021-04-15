package main

import (
	"flag"
	"fmt"
)

type celsiusFlag struct {
	value float32
	unit string
}

func (c *celsiusFlag) String() string {
	return fmt.Sprintf("%f%s", c.value, c.unit)
}

func (c *celsiusFlag) Set(str string) error {
	var unit string
	var value float32
	_, err := fmt.Sscanf(str, "%f%s", &value, &unit)
	if err == nil {
		c.unit = unit
		c.value = value
		return nil
	} else {
		return err
	}
}

func CelsiusFlag(name string, value string, usage string) *celsiusFlag {
	f := celsiusFlag{}
	err := f.Set(value)
	fmt.Println(f)
	if err == nil {
		flag.CommandLine.Var(&f, name, usage)
		return &f
	} else {
		return nil
	}

}

func main()  {
	var temp = CelsiusFlag("temp", "20.0C" , "temperature")
	flag.Parse()
	fmt.Println("temp: ", temp.value, temp.unit)
}