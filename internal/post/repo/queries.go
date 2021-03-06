package repo

const (
	SELECT_QUERY = "SELECT id, title, slug,body,description,is_draft,author_id,created_at,updated_at,deleted_at FROM post LIMIT $1 OFFSET $2"
	INSERT_QUERY = `
		INSERT INTO post (title, slug,body,description,is_draft,author_id) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, title, slug,body,description,is_draft,author_id,created_at,updated_at,deleted_at`
)
