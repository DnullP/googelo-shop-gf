// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"googelo-shop-gf/app/auth/internal/dao/internal"
)

// internalTesttableDao is an internal type for wrapping the internal DAO implementation.
type internalTesttableDao = *internal.TesttableDao

// testtableDao is the data access object for the table testtable.
// You can define custom methods on it to extend its functionality as needed.
type testtableDao struct {
	internalTesttableDao
}

var (
	// Testtable is a globally accessible object for table testtable operations.
	Testtable = testtableDao{
		internal.NewTesttableDao(),
	}
)

// Add your custom methods and functionality below.
