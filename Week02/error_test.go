package Week02


//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

import (
	"database/sql"
	"github.com/pkg/errors"
	"testing"
)

var (
	DB *sql.DB
)

type IDao interface {
	getTableName() string
}

type UserDao struct {
	Dao
}
type Dao struct {
}


func (d *Dao) GetRecordById(ID int) (ret interface{}, err error) {
	//err = DB.QueryRow("SELECT name FROM user WHERE id = ?",userID).Scan()
	err = sql.ErrNoRows
	return ret, errors.Wrap(err, "get db error" )
}

func TestError(t *testing.T){
	ret,err := new(UserDao).GetRecordById(1)
	t.Log(ret,err)
	//todo  err 吞掉或是返回
}
