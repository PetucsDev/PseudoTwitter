package comentarios

import(
	"context"
	"database/sql"
	"PseudoTwitter/internal/domain"
)

type Repository interface{
	GetAllByUsers(ctx context.Context, id int) ([]domain.Comments, error)
	GetAllByPublications(ctx context.Context, id int) ([]domain.Comments, error)
	Save(ctx context.Context, u domain.Comments) (int, error)
	Update(ctx context.Context, u domain.Comments) error
	Delete(ctx context.Context, id int) error
}

type repository struct{
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}


func (r *repository) GetAllByUsers(ctx context.Context, id int) ([]domain.Comments, error) {
	query := "SELECT * FROM comentarios WHERE usuarios_id=?;"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var users []domain.Comments

	for rows.Next() {
		s := domain.Comments{}
		_ = rows.Scan(&s.ID, &s.Descripcion, &s.UsuariosId, &s.PublicacionesId)
		users = append(users, s)
	}

	return users, nil
}

func (r *repository) GetAllByPublications(ctx context.Context, id int) ([]domain.Comments, error) {
	query := "SELECT * FROM comentarios WHERE publicaciones_id=?;"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var users []domain.Comments

	for rows.Next() {
		s := domain.Comments{}
		_ = rows.Scan(&s.ID, &s.Descripcion, &s.UsuariosId, &s.PublicacionesId)
		users = append(users, s)
	}

	return users, nil
}



func (r *repository) Save(ctx context.Context, s domain.Comments) (int, error) {
	query := "INSERT INTO comentarios (descripcion, usuarios_id, publicaciones_id) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.Descripcion, s.PublicacionesId, s.UsuariosId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Comments) error {
	query := "UPDATE comentarios SET descripcion=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.Descripcion, s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM comentarios WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return err
	}

	return nil
}