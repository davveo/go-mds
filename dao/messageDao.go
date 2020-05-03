package dao

import (
	"time"

	"github.com/zbrechave/go-mds/models"

	mydb "github.com/zbrechave/go-mds/db/mysql"
)

type MessageDao struct{}

func (md *MessageDao) Insert(message *models.Message) error {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`," +
			"`file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		return err
	}
	return nil

}
