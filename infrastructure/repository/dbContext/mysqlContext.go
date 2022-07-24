package dbContext

import (
	"fops/infrastructure/repository/admin"
	"fs/data"
)

type MysqlContext struct {
	*data.DbContext
	Admin data.TableSet[admin.PO]
}

func NewContext() MysqlContext {
	context := data.NewDbContext("mysql")
	return MysqlContext{
		DbContext: context,
		Admin:     data.NewTableSet(context, "admin", admin.PO{}),
	}
}
