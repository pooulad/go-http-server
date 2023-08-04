package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var migrateOnce sync.Once = sync.Once{}

func MigrateOrPanic(ctx context.Context,db *sql.DB) (err error) {
	migrateOnce.Do(func() {
		fpath,fpathErr := filepath.Abs(filepath.Join("assets","track_table.sql"))
		if fpathErr != nil {
			fmt.Println(fpath)
			return
		}
	
		fbyte,fbyteErr := os.ReadFile(fpath)
		if fbyteErr != nil {
			return
		}
		
		_,exeErr := db.ExecContext(ctx,string(fbyte))
		if exeErr != nil {
			return
		}
	})

	return
}