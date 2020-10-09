package db_mysql

func QueryAdmin(name string ,pwd string)(int,error){
	row :=LolDB.QueryRow("select count(admin_name) admin_num from lol_admin where admin_name = ? and admin_pwd = ?",
		name,pwd)
	var admin_num int
	err := row.Scan(&admin_num)
	if err != nil {
		return 0,err
	}
	return admin_num,nil
}
