-- +goose Up

CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  title TEXT,
  description TEXT, 
  url TEXT,
  published_at TIMESTAMP,
  feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  CONSTRAINT url UNIQUE (url)
);

-- +goose Down 
DROP TABLE posts;