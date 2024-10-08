-- name: GetNextFetched :one 
SELECT * 
FROM feeds 
ORDER BY last_fetched_at NULLS first 
LIMIT 1; 