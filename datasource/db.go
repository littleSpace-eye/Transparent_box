package datasource

import (
	"database/sql"
)

// 初始化数据库
func Init() (err error) {
	db, _ := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/demo")


	// sql语句，如果没存在库表todo，则新建一个
	var sqlStr = `create table make_up
(
    id       int auto_increment
        primary key,
    name     varchar(192)                 not null,
    tel      varchar(192)                 not null,
    e_mail varchar(192)                 null,
    password varchar(92) default '123456' not null,
    sex      tinyint(1)  default 0        not null,
    birth    date                         null,
    constraint make_up_id_uindex
        unique (id)
);`

	// 执行建表语句
	_, err = db.Exec(sqlStr)
	if err != nil {
		return err
	}
	defer db.Close()
	return
}
