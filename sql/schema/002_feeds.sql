-- +goose Up
CREATE TABLE feeds(
  id UUID PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
  name TEXT NOT NULL,
  url TEXT NOT NULL,
  user_id UUID,
  CONSTRAINT unique_url UNIQUE (url),
  FOREIGN KEY(user_id) references users(id) ON DELETE CASCADE
);

-- +goose Down 
DROP TABLE feeds;