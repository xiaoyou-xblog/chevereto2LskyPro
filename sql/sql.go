// @Title  mysql驱动简单封装1.0版本
// @Description  该驱动仅适用于mysql
// @Author 小游
// @Date 2021/01/25
package sql

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql" //这里我们导入驱动(注册匿名包)
	"img/common"
	"strings"
)



// 数据库连接池(第一个是cheverto 第二个是lsky)
var DB1 *sql.DB
var DB2 *sql.DB


// 初始化数据库连接池
func InitDb1() bool {
	//先关闭之前的数据库
	if DB1 != nil {
		_ = DB1.Close()
	}
	//读取配置
	data := common.GetConfig("cheverto")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{data["username"], ":", data["password"], "@tcp(", data["ip"], ":", data["port"], ")/", data["Dbname"], "?charset=utf8mb4&collation=utf8mb4_unicode_ci"}, "")
	////打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sqltool-driver/mysql"
	DB1, _ = sql.Open("mysql", path)
	////设置数据库最大连接数
	DB1.SetConnMaxLifetime(100)
	////设置数据库最大闲置连接数
	DB1.SetMaxIdleConns(10)
	////验证连接
	if DB1.Ping() != nil {
		return false
	}
	return true
}
// 数据库重连函数
func db1Reconnect() {
	if DB1 != nil {
		if DB1.Ping() != nil {
			return
		}
	}
	InitDb1()
}
// 关闭数据库
func Db1Close() error {
	err := DB1.Close()
	if err != nil {
		return err
	}
	return nil
}

// 初始化数据库连接池
func InitDb2() bool {
	//先关闭之前的数据库
	if DB2 != nil {
		_ = DB2.Close()
	}
	//读取配置
	data := common.GetConfig("lsky")
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{data["username"], ":", data["password"], "@tcp(", data["ip"], ":", data["port"], ")/", data["Dbname"], "?charset=utf8mb4&collation=utf8mb4_unicode_ci"}, "")
	////打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sqltool-driver/mysql"
	DB2, _ = sql.Open("mysql", path)
	////设置数据库最大连接数
	DB2.SetConnMaxLifetime(100)
	////设置数据库最大闲置连接数
	DB2.SetMaxIdleConns(10)
	////验证连接
	if DB2.Ping() != nil {
		return false
	}
	return true
}
// 数据库重连函数
func db2Reconnect() {
	if DB2 != nil {
		if DB2.Ping() != nil {
			return
		}
	}
	InitDb2()
}
// 关闭数据库
func Db2Close() error {
	err := DB2.Close()
	if err != nil {
		return err
	}
	return nil
}


// 数据库1查询语句
func Db1Dql(sql string, args ...interface{}) ([][]string, error) {
	//fmt.Println(sql)
	//fmt.Println(sql) -->这里通过传递一个可变长的参数，我们用这个来实现参数化查询
	//数据库自动重连
	db1Reconnect()
	var result [][]string
	//这里通过传递一个可变长的参数，我们用这个来实现参数化查询
	rows, err := DB1.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//这边是利用反射还有接口来获取所有内容
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.New("查询行数失败")
	}
	pointers := make([]interface{}, len(cols))
	_ = rows.Scan(pointers...)
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	result = append(result, container)
	for rows.Next() {
		err = rows.Scan(pointers...)
		container := make([]string, len(cols))
		for i, _ := range pointers {
			pointers[i] = &container[i]
		}
		result = append(result, container)
	}
	if len(result) != 1 {
		result = result[:len(result)-1]
	}
	return result, nil
}

// 数据库1查询语句
func Db2Dql(sql string, args ...interface{}) ([][]string, error) {
	//fmt.Println(sql)
	//fmt.Println(sql) -->这里通过传递一个可变长的参数，我们用这个来实现参数化查询
	//数据库自动重连
	db2Reconnect()
	var result [][]string
	//这里通过传递一个可变长的参数，我们用这个来实现参数化查询
	rows, err := DB2.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//这边是利用反射还有接口来获取所有内容
	cols, err := rows.Columns()
	if err != nil {
		return nil, errors.New("查询行数失败")
	}
	pointers := make([]interface{}, len(cols))
	_ = rows.Scan(pointers...)
	container := make([]string, len(cols))
	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	result = append(result, container)
	for rows.Next() {
		err = rows.Scan(pointers...)
		container := make([]string, len(cols))
		for i, _ := range pointers {
			pointers[i] = &container[i]
		}
		result = append(result, container)
	}
	if len(result) != 1 {
		result = result[:len(result)-1]
	}
	return result, nil
}

// 数据库2插入语句
func Db2Dml(sql string, args ...interface{}) bool {
	//fmt.Println(sql)
	//数据库自动重连
	db2Reconnect()
	tx, err := DB2.Begin()
	if err != nil {
		return false
	}
	_, err = tx.Exec(sql, args...)
	if err == nil {
		err = tx.Commit()
		if err == nil {
			return true
		}
	}
	return false
}


