package repository

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"sort"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	// "github.com/Al3XS0n-sempai/distributed_systems/internal/service"
)

type SharededPostgresql struct {
	shards   map[int]*pgxpool.Pool
	replicas int
	keys     []int                    // отсортированный список хешей
	nodeIDs  map[*pgxpool.Pool]string // вспомогательная карта для генерации виртуальных узлов
}

// hash функция из интернета
func hashKey(key string) int {
	h := sha1.New()
	h.Write([]byte(key))
	hash := h.Sum(nil)

	return int((uint32(hash[0]) << 24) | (uint32(hash[1]) << 16) | (uint32(hash[2]) << 8) | uint32(hash[3]))
}

func (db SharededPostgresql) getShard(key string) *pgxpool.Pool {
	if len(db.keys) == 0 {
		return nil
	}
	hash := hashKey(key)
	idx := sort.Search(len(db.keys), func(i int) bool {
		return db.keys[i] >= hash
	})
	if idx == len(db.keys) {
		idx = 0
	}
	// CHECK THAT KEYS ARE DIFFERENT
	// fmt.Println(idx)
	return db.shards[db.keys[idx]]
}

func (db *SharededPostgresql) addShard(nodeID string, pool *pgxpool.Pool) {
	db.nodeIDs[pool] = nodeID
	for i := 0; i < db.replicas; i++ {
		virtualKey := nodeID + "#" + strconv.Itoa(i)
		hash := hashKey(virtualKey)
		db.keys = append(db.keys, hash)
		db.shards[hash] = pool
	}
	sort.Ints(db.keys)
}

func (db SharededPostgresql) Get(key string) (string, error) {
	shard := db.getShard(key)
	var value string
	err := shard.QueryRow(context.Background(), `SELECT value FROM kv_store WHERE key = $1`, key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("key not found")
	}
	return value, err
}

func (db SharededPostgresql) Set(key, value string) error {
	shard := db.getShard(key)
	_, err := shard.Exec(context.Background(), `
        INSERT INTO kv_store (key, value) VALUES ($1, $2)
        ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value
    `, key, value)
	return err
}

func initTable(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS kv_store (
			key TEXT PRIMARY KEY,
			value TEXT
		)
	`)
	if err != nil {
		return fmt.Errorf("create table error: %w", err)
	}
	return nil
}

// NewShardedPostgresql create new Sharded Postgresql repository
func NewShardedPostgresql(connStrings []string) (*SharededPostgresql, error) {
	SharededDB := &SharededPostgresql{
		shards:   make(map[int]*pgxpool.Pool),
		replicas: 1, // пока так а то вообше тяжело столько БД поднимать
		nodeIDs:  make(map[*pgxpool.Pool]string),
	}
	for i, connStr := range connStrings {
		config, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			return nil, err
		}
		config.MaxConns = 200
		config.MinConns = 100
		// db, err := pgxpool.New(context.Background(), connStr)
		db, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			return nil, err
		}
		if err = initTable(db); err != nil {
			return nil, err
		}
		SharededDB.addShard(strconv.Itoa(i), db)
	}
	return SharededDB, nil
}

// check for Interface
// var _ service.SimpleServiceRepository = SharededPostgresql{}
