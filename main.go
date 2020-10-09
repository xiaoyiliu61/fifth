package main

import (
	"11lol/db_mysql"
	"11lol/entity"
	"11lol/reqhandle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	err := db_mysql.OpenDatabase()
	defer db_mysql.CloseDatabase()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("欢迎来到英雄联盟，敌军还有5秒到达战场！")

	url := "https://game.gtimg.cn/images/lol/act/img/js/heroList/hero_list.js"

	client := http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	/*fmt.Println(string(dataBytes))*/

	var herolist entity.Herolist
	err = json.Unmarshal(dataBytes, &herolist)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	/*fmt.Println(herolist.Hero[0].Name)*/

	/*fmt.Println(len(herolist.Hero))*/

	/*for i:=0;i<len(herolist.Hero);i++ {
		fmt.Println(herolist.Hero[i])
	}*/
	/*fmt.Println(herolist.Hero[0].Name)*/


	lolNum,err := db_mysql.QuerylolsNum()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if lolNum <=0 {
		for i := 0; i < len(herolist.Hero); i++ {
			_, err := db_mysql.SaveLols2Db(herolist,i)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
	fmt.Printf("数据库中已经存在数据%d条\n",lolNum)



	http.HandleFunc("/",reqhandle.Index)

	http.HandleFunc("/login",reqhandle.Login)


	fmt.Println("当前网络已开启监听，可以通过http://127.0.0.1:9090进行访问")

	err =http.ListenAndServe(":9090",nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
