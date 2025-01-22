// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TesttableDao is the data access object for the table testtable.
type TesttableDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of the current DAO.
	columns TesttableColumns // columns contains all the column names of Table for convenient usage.
}

// TesttableColumns defines and stores column names for the table testtable.
type TesttableColumns struct {
	Name string //
	Age  string //
}

// testtableColumns holds the columns for the table testtable.
var testtableColumns = TesttableColumns{
	Name: "name",
	Age:  "age",
}

// NewTesttableDao creates and returns a new DAO object for table data access.
func NewTesttableDao() *TesttableDao {
	return &TesttableDao{
		group:   "default",
		table:   "testtable",
		columns: testtableColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TesttableDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TesttableDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TesttableDao) Columns() TesttableColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TesttableDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TesttableDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TesttableDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
