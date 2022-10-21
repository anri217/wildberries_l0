package migrate

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TryInitTables(pool *pgxpool.Pool) error {
	path := "migrate/init.sql"

	c, io_err := ioutil.ReadFile(path)
	if io_err != nil {
		return fmt.Errorf("ERROR: can't find imit sql file: %v", io_err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("ERROR: can't return connection from the Pool: %v", err)
	}
	defer conn.Release()

	sql := string(c)
	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("ERROR: can't run init script: %v", err)
	}

	return nil
}
