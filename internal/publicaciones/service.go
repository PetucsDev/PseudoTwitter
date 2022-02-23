package publicaciones

import(
	"PseudoTwitter/internal/domain"
	"context"
	"errors"
)


var(
	ErrNotFound = errors.New("publication not found")

)

type Service interface{
	GetAll(ctx context.Context, id int) ([]domain.Publications, error)
	Get(ctx context.Context, id int) (domain.Publications, error)
	Save(ctx context.Context, u domain.Publications) (domain.Publications, error)
	Update(ctx context.Context, u domain.Publications) error
	Delete(ctx context.Context, id int) error
}


type service struct {
	repo Repository
}


func NewService(s Repository) Service {
	return &service {repo: s,
	}
}

func (s *service) GetAll(ctx context.Context, id int) ([]domain.Publications, error) {
	ps, err := s.repo.GetAll(ctx, id)
	if err != nil {
		return nil, err
	}
	 return ps,nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Publications, error) {
	p, err := s.repo.Get(ctx, id)
	if err != nil {
		return domain.Publications{}, err
	}

	return p, nil
}


func (s *service) Save(ctx context.Context, se domain.Publications) (domain.Publications, error) {
	
	p, err := s.repo.Save(ctx,se)

	if err != nil {
		return domain.Publications{}, err
	}

	se.ID = p

	return se, nil
}


func (s *service) Update(ctx context.Context, se domain.Publications)  error {

	return s.repo.Update(ctx, se)

 }

 func (s *service) Delete(ctx context.Context, id int) error {
	
	return s.repo.Delete(ctx, id)
	
}