package user

import(
	"PseudoTwitter/internal/domain"
	"context"
	"errors"
)


var(
	ErrNotFound = errors.New("user not found")

)

type Service interface{
	GetAll(ctx context.Context) ([]domain.Users, error)
	Get(ctx context.Context ,id int ) (domain.Users, error)
	Save(ctx context.Context,us domain.Users) (domain.Users, error)
	Update(ctx context.Context,  us domain.Users) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, userName string) bool
}


type service struct {
	repo Repository
}


func NewService(s Repository) Service {
	return &service {repo: s,
	}
}


func (s *service) GetAll(ctx context.Context) ([]domain.Users, error) {
	ps, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	 return ps,nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Users, error) {
	p, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Users{}, err
	}

	return p, nil
}

func (s *service) Save(ctx context.Context, se domain.Users) (domain.Users, error) {
	exist := s.repo.Exists(ctx, se.UserName)
	if exist {
			return domain.Users{}, errors.New("el user ya existe")
	}
	p, err := s.repo.Save(ctx,se)

	if err != nil {
		return domain.Users{}, err
	}

	se.ID = p

	return se, nil
}

func (s *service) Update(ctx context.Context, se domain.Users)  error {

	return s.repo.Update(ctx, se)

 }

func (s *service) Delete(ctx context.Context, id int) error {
	
	return s.repo.Delete(ctx, id)
	
}

func (s *service) Exists(ctx context.Context, userName string) bool {
	return s.repo.Exists(ctx, userName)
}