package go_tc_homework

import (
    "context"
    "database/sql"
    "github.com/pkg/errors"
    "time"
)

var pool *sql.DB

// 检索数据：测试 wrap error 用例
func Query(ctx context.Context, id int64) (string, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    var name string
    err := pool.QueryRowContext(ctx, "select p.name from people as p where p.id = :id;",
        sql.Named("id", id)).Scan(&name)
    if err != nil {
        return "", errors.Wrap(err, "unable to execute search query")
    }

    return "", nil
}
