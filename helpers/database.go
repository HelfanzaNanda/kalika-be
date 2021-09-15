package helpers

import (
	"fmt"
	"gorm.io/gorm"
)

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		fmt.Println(errorRollback)
		fmt.Println(err)
	} else {
		errorCommit := tx.Commit()
		fmt.Println(errorCommit)
	}
}