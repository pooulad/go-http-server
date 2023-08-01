package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
)

func MigrateOrPanic(ctx context.Context,db *sql.DB) error {
	fpath,fpathErr := filepath.Abs(filepath.Join("assets","track_table.sql"))
	if fpathErr != nil {
		return fpathErr
	}

	fbyte,fbyteErr := os.ReadFile(fpath)
	if fbyteErr != nil {
		return fbyteErr
	}
	
	_,exeErr := db.ExecContext(ctx,string(fbyte))
	if exeErr != nil {
		return exeErr
	}

	return nil
}