package user


import(
	"context"
	"database/sql"
	"PseudoTwitter/internal/domain"
)


type Repository interface{
	GetAll(ctx context.Context) ([]domain.Users, error)
	Get(ctx context.Context, id int) (domain.Users, error)
	Exists(ctx context.Context, userName string) bool
	Save(ctx context.Context, u domain.Users) (int, error)
	Update(ctx context.Context, u domain.Users) error
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

func (r *repository) GetAll(ctx context.Context) ([]domain.Users, error) {
	query := "SELECT * FROM usuarios"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []domain.Users

	for rows.Next() {
		s := domain.Users{}
		_ = rows.Scan(&s.ID, &s.UserName, &s.Password, &s.Mail)
		users = append(users, s)
	}

	return users, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Users, error) {
	query := "SELECT * FROM usuarios WHERE id=?;"
	row := r.db.QueryRow(query, id)
	s := domain.Users{}
	err := row.Scan(&s.ID, &s.UserName, &s.Password, &s.Mail)
	if err != nil {
		return domain.Users{}, err
	}

	return s, nil
}

func (r *repository) Exists(ctx context.Context, userName string) bool {
	query := "SELECT username FROM usuarios WHERE username=?;"
	row := r.db.QueryRow(query, userName)
	err := row.Scan(&userName)
	return err == nil
}


func (r *repository) Save(ctx context.Context, s domain.Users) (int, error) {
	query := "INSERT INTO usuarios (username, password, mail) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.UserName, s.Password, s.Mail)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Users) error {
	query := "UPDATE usuarios SET username=?, password=?, mail=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.Mail, s.Password, s.UserName, s.ID)
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
	query := "DELETE FROM usuarios WHERE id=?"
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