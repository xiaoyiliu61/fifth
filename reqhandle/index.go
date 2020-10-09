package reqhandle

import (
	"html/template"
	"net/http"
)

//处理请求的函数
func Index(writer http.ResponseWriter,request *http.Request){
	//writer.Write([]byte("hello every one"))
	temp,err :=template.ParseFiles("./views/index.html")
	if err !=nil {
		//writer.Write([]byte(err.Error()))
		errorTmp,_ :=template.ParseFiles("./views/error.html")
		errorTmp.Execute(writer,err.Error())
		return
	}
	//解析模板正常
	//exe:Execute的缩写
	temp.Execute(writer,nil)
}
