package domain

import (
	"time"
)

type IndexerCheckpoint struct {
	ID                     int64     `gorm:"column:id;primaryKey;autoIncrement"`
	StreamName             string    `gorm:"column:stream_name;not null"`
	ChainID                int64     `gorm:"column:chain_id;not null"`
	LastProcessedBlock     int64     `gorm:"column:last_processed_block;not null"`
	LastProcessedLogIndex  int32     `gorm:"column:last_processed_log_index;not null"`
	LastProcessedBlockHash *string   `gorm:"column:last_processed_block_hash"`
	UpdatedAt              time.Time `gorm:"column:updated_at;not null"`
}

func (IndexerCheckpoint) TableName() string {
	return "indexer_checkpoints"
}

type IndexedBlock struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ChainID     int64     `gorm:"column:chain_id;not null"`
	BlockNumber int64     `gorm:"column:block_number;not null"`
	BlockHash   string    `gorm:"column:block_hash;not null"`
	ParentHash  string    `gorm:"column:parent_hash;not null"`
	BlockTime   time.Time `gorm:"column:block_time;not null"`
	IsFinalized bool      `gorm:"column:is_finalized;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
}

func (IndexedBlock) TableName() string {
	return "indexed_blocks"
}

type ReorgEvent struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement"`
	ChainID      int64     `gorm:"column:chain_id;not null"`
	DetectedAt   time.Time `gorm:"column:detected_at;not null"`
	FromBlock    int64     `gorm:"column:from_block;not null"`
	ToBlock      int64     `gorm:"column:to_block;not null"`
	RowsReverted int32     `gorm:"column:rows_reverted;not null"`
	Notes        *string   `gorm:"column:notes"`
	CreatedAt    time.Time `gorm:"column:created_at;not null"`
}

func (ReorgEvent) TableName() string {
	return "reorg_events"
}
