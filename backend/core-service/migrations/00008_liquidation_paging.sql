-- +goose Up
CREATE INDEX liquidations_time_idx ON liquidations (block_time DESC, id DESC);

DROP INDEX IF EXISTS liquidations_market_time_idx;

CREATE INDEX liquidations_market_time_idx ON liquidations (market_id, block_time DESC, id DESC);

-- +goose Down
DROP INDEX IF EXISTS liquidations_market_time_idx;

CREATE INDEX liquidations_market_time_idx ON liquidations (market_id, block_time DESC);

DROP INDEX IF EXISTS liquidations_time_idx;
