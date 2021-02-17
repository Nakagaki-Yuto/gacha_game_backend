package handler

type Handler struct {
	db *DB
}

func New(db *DB) *Handler {
	return &Handler{
		db: db,
	}
}