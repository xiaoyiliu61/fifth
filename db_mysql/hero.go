package db_mysql

import (
	"11lol/entity"
	"database/sql"

)
import  _"github.com/go-sql-driver/mysql"

var LolDB *sql.DB
func OpenDatabase()(error) {
	if LolDB!=nil {
		return nil
	}
	database,err:=sql.Open("mysql","root:409216@tcp(127.0.0.1:3306)/herolist?charset=utf8")
	if err!=nil {
		return err
	}
	LolDB=database
	return nil
}
func CloseDatabase()error  {
	if LolDB!=nil {
		err:=LolDB.Close()
		if err!=nil {
			return err
		}
	}
	return nil
}
func SaveLols2Db(lol entity.Herolist,i int)( int64,error)  {
		result,err := LolDB.Exec("insert into " +
			"lol(" +
			"HeroId,Name,Alias,Title,Attack,Defense,Magic,Difficulty)" +
			"values(" +
			"?,?,?,?,?,?,?,?)",
			 lol.Hero[i].HeroId,lol.Hero[i].Name,lol.Hero[i].Alias, lol.Hero[i].Title, lol.Hero[i].Attack, lol.Hero[i].Defense, lol.Hero[i].Magic, lol.Hero[i].Difficulty)
		if err != nil {

			return  0,err
		}
		rowId, err := result.RowsAffected()
		if err != nil {
			return 0, err
		}
		return rowId, nil


}

func QuerylolsNum() (int,error){
	rows :=LolDB.QueryRow("select count(HeroId) recordnum from lol")
	var recordNum int
	err :=rows.Scan(&recordNum)
	if err != nil {
		return 0,err
	}
	return recordNum,nil
}


func QueryAllLols(i int) ([]entity.Herolist,error){
	rows,err :=LolDB.Query("select * from lol")
	if err != nil {
		return nil,err
	}
	lols := make([]entity.Herolist,0)
	for  i=0;i<len(lols);i++{
		var lol entity.Herolist

		err =rows.Scan(&lol.Hero[i].HeroId,&lol.Hero[i].Name,&lol.Hero[i].Alias,&lol.Hero[i].Title,&lol.Hero[i].Attack,&lol.Hero[i].Defense,&lol.Hero[i].Magic,&lol.Hero[i].Difficulty)
		if err != nil {
			return nil,err
		}
		lols = append(lols,lol)
	}
	return lols,nil
}
