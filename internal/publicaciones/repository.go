package publicaciones

import(
	"context"
	"database/sql"
	"PseudoTwitter/internal/domain"
)


type Repository interface{
	GetAll(ctx context.Context, id int) ([]domain.Publications, error)
	Get(ctx context.Context, id int) (domain.Publications, error)
	Save(ctx context.Context, u domain.Publications) (int, error)
	Update(ctx context.Context, u domain.Publications) error
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

func (r *repository) GetAll(ctx context.Context, id int) ([]domain.Publications, error) {
	query := "SELECT * FROM publicaciones WHERE usuarios_id=?;"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	var users []domain.Publications

	for rows.Next() {
		s := domain.Publications{}
		_ = rows.Scan(&s.ID, &s.Titulo, &s.Fecha, &s.UsuariosId)
		users = append(users, s)
	}

	return users, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Publications, error) {
	query := "SELECT * FROM publicaciones WHERE usuarios_id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Publications{}
	err := row.Scan(&s.ID, &s.Titulo, &s.Fecha, &s.UsuariosId)
	if err != nil {
		return domain.Publications{}, err
	}

	return s, nil
}

func (r *repository) Save(ctx context.Context, s domain.Publications) (int, error) {
	query := "INSERT INTO publicaciones (titulo, fecha, usuarios_id) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.Titulo, s.Fecha, s.UsuariosId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Publications) error {
	query := "UPDATE publicaciones SET titulo=?, fecha=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.Titulo, s.Fecha, s.ID)
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
	query := "DELETE FROM publicaciones WHERE usuarios_id=?"
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