package mysql

import (
	"github.com/go-sql-driver/mysql"
)

func IsAlreadyExistErr(err error) bool {
	if err == nil {
		return false
	}
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1050: // MySQL code for already exists
			return true
		default:
			// Handle other errors
		}
	}
	return false
}
