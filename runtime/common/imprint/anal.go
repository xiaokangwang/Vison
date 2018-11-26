package imprint

import (
	"bytes"
	"strconv"
)

func IsKInput(imp []byte)bool{
	sign:=[]byte("func.kinput.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetKInputVal(imp []byte)string{
	sign:=[]byte("func.kinput.")
	return string(imp[0:len(sign):len(imp)-1])
}

func IsQInput(imp []byte)bool{
	sign:=[]byte("func.qinput.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetQInputVal(imp []byte)int{
	sign:=[]byte("func.qinput.")
	i,_:=strconv.Atoi(string(imp[0:len(sign):len(imp)-1]))
	return i
}


func IsKOutput(imp []byte)bool{
	sign:=[]byte("func.koutput.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetKOutputVal(imp []byte)string{
	sign:=[]byte("func.koutput.")
	return string(imp[0:len(sign):len(imp)-1])
}

func IsQOutput(imp []byte)bool{
	sign:=[]byte("func.qoutput.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetQOutputVal(imp []byte)int{
	sign:=[]byte("func.qinput.")
	i,_:=strconv.Atoi(string(imp[0:len(sign):len(imp)-1]))
	return i
}

func IsIntConst(imp []byte)bool{
	sign:=[]byte("const.int.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetIntConst(imp []byte)int{
	sign:=[]byte("const.int.")
	i,_:=strconv.Atoi(string(imp[0:len(sign):len(imp)]))
	return i
}

func IsStringConst(imp []byte)bool{
	sign:=[]byte("const.string.")

	return bytes.Equal(imp[0:len(sign)],sign)
}

func GetStringConst(imp []byte)string{
	sign:=[]byte("const.string.")
	return string(imp[0:len(sign):len(imp)])
}
