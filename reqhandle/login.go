package reqhandle

import (
	"11lol/db_mysql"
	"11lol/entity"
	"html/template"
	"net/http"
)

func Login(writer http.ResponseWriter,request *http.Request){

	err :=request.ParseForm()
	if err != nil {
		tmpt, _ :=template.ParseFiles("./views/error.html")
		tmpt.Execute(writer,err.Error())
		return
	}

	adminName := request.FormValue("user_name")
	adminPwd := request.FormValue("user_pwd")

	if adminName == "" || adminPwd ==  ""{
		tmpt, _ :=template.ParseFiles("./views/error.html")
		tmpt.Execute(writer,"用户名或者密码为空，请检查后重新尝试")
		return
	}

	//根据这个两个值，到数据库中进行匹配
	admin_num,err :=db_mysql.QueryAdmin(adminName,adminPwd)
	if err != nil {
		tmpt, _ :=template.ParseFiles("./views/error.html")
		tmpt.Execute(writer,"用户名或者密码为空，请检查后重新尝试")
		return
	}

	if admin_num > 0 {
		lols,err :=db_mysql.QueryAllLols(148)
		if err != nil {
			tmpt, _ :=template.ParseFiles("./views/error.html")
			tmpt.Execute(writer,err.Error())
			return
		}

		tempt,_ := template.ParseFiles("./views/home.html")

		showData := entity.HomeShowData{AdminName:adminName,AllLols:lols}
		tempt.Execute(writer,showData)
	}

}
