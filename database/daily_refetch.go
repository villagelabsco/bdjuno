package database

import "fmt"

// GetTotalBlocks implements database.Database
func (db *Db) GetTotalBlocks() int64 {
	var blockCount int64
	err := db.SQL.QueryRow(`SELECT count(*) FROM block;`).Scan(&blockCount)
	if err != nil {
		panic(fmt.Errorf("failed getting total blocks: %d", blockCount))
	}
	return blockCount
}

// GetMissingBlocks returns an array of missing blocks from one day ago
func (db *Db) GetMissingBlocks(startHeight, endHeight int64) []int64 {
	var result []int64
	stmt := `SELECT generate_series($1::int,$2::int) EXCEPT SELECT height FROM block ORDER BY 1;`
	err := db.Sqlx.Select(&result, stmt, startHeight, endHeight)
	if err != nil {
		return nil
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
