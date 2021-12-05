package CreateDBConnect

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)
var config = mysql.Config{
	User: "hekapoo",
	Passwd: "1234",
	Net: "tcp",
	Addr: "localhost:3306",
	DBName: "libtest",
	Collation: "",
}
func CreateClient(retry int)(*sql.DB,error) {
	db,err:=sql.Open("mysql",config.FormatDSN())
	if err:=db.Ping();err!=nil {
		if retry>=3 {
			return nil,err
		}
		retry++
		time.Sleep(time.Second)
		return CreateClient(retry)
	}
	return db,err
}
