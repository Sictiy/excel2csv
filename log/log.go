package log

import "fmt"

func CheckError(err error) bool {
	if err != nil{
		fmt.Print(err)
		return true
	}
	return false
}

func LogFormat(format string, a ... interface{}){
	if format == "" {
		Log(a...)
		return
	}
	format = format + "\n"
	if a != nil{
		_,e := fmt.Printf(format, a ...)
		CheckError(e)
	}else{
		fmt.Print(format)
	}
}

func Log(a ... interface{})  {
	fmt.Print(a ...)
	fmt.Print("\n")
}
