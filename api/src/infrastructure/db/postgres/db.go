package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"

	"github.com/kmifa/QuizWhiz/config"
)

var txKey = struct{}{}

type DB struct {
	client *bun.DB
}

func NewDB() *DB {
	dbConfig := config.Env().Db

	var sqldb *sql.DB
	if dbConfig.TcpHost != "" {
		// TCP connection for local
		// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Tokyo", dbConfig.TcpHost, dbConfig.UserName, dbConfig.Password, dbConfig.Name)
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.UserName, dbConfig.Password, dbConfig.TcpHost, "5432", dbConfig.Name)
		config := pgdriver.NewConnector(pgdriver.WithDSN(dsn))
		sqldb = sql.OpenDB(config)
	} else {
		// Unix socket connection for cloud sql
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s TimeZone=Asia/Tokyo", dbConfig.InstanceConnectName, dbConfig.UserName, dbConfig.Password, dbConfig.Name)
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			log.Fatalf("db.NewDB: failed to parse config. %w", err)
		}
		config.PreferSimpleProtocol = true
		sqldb = stdlib.OpenDB(*config)
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	configureConnectionPool(db)

	return &DB{
		client: db,
	}
}

func configureConnectionPool(db *bun.DB) {
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(1800 * time.Second)
}

// GetDB は `bun.IDB` インターフェイスを返します。
// 実際には `*bun.DB` または `bun.Tx` のどちらかが返ります。
func (s *DB) GetDB(ctx context.Context) bun.IDB {
	tx, ok := ctx.Value(txKey).(bun.Tx)
	if ok {
		return tx
	}
	return s.client
}

// Tx はトランザクションを開始しつつ、引数のfn関数を実行します。
// fn関数の中では `GetDB()` を使って `bun.Tx` を取得し、クエリを発行してください。
// もしもfn関数がエラーを返したらロールバックしつつエラーを返します。
func (s *DB) Tx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := s.client.Begin()
	if err != nil {
		return fmt.Errorf("db.DB.Tx: failed to begin. %w", err)
	}

	newCtx := context.WithValue(ctx, txKey, tx)

	if err := fn(newCtx); err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("db.DB.Tx: failed to rollback. %w", err)
		}
		return fmt.Errorf("db.DB.Tx: transaction rolled back. %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("db.DB.Tx: failed to commit. %w", err)
	}

	return nil
}

// SetDB は `*sql.DB` から `*bun.DB` を生成します
func SetDB(sqldb *sql.DB) *DB {
	db := bun.NewDB(sqldb, pgdialect.New())
	return &DB{client: db}
}
