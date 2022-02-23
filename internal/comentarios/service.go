package comentarios

import(
	"PseudoTwitter/internal/domain"
	"context"
	"errors"
)


var(
	ErrNotFound = errors.New("comment not found")

)


type Service interface{
	GetAllByUsers(ctx context.Context, id int) ([]domain.Comments, error)
	GetAllByPublications(ctx context.Context, id int) ([]domain.Comments, error)
	Save(ctx context.Context, u domain.Comments) (domain.Comments, error)
	Update(ctx context.Context, u domain.Comments) error
	Delete(ctx context.Context, id int) error
}


type service struct {
	repo Repository
}


func NewService(s Repository) Service {
	return &service {repo: s,
	}
}


func (s *service) GetAllByUsers(ctx context.Context, id int) ([]domain.Comments, error) {
	ps, err := s.repo.GetAllByUsers(ctx, id)
	if err != nil {
		return nil, err
	}
	 return ps,nil
}

func (s *service) GetAllByPublications(ctx context.Context, id int) ([]domain.Comments, error) {
	ps, err := s.repo.GetAllByPublications(ctx, id)
	if err != nil {
		return nil, err
	}
	 return ps,nil
}

func (s *service) Save(ctx context.Context, se domain.Comments) (domain.Comments, error) {
	
	p, err := s.repo.Save(ctx,se)

	if err != nil {
		return domain.Comments{}, err
	}

	se.ID = p

	return se, nil
}


func (s *service) Update(ctx context.Context, se domain.Comments)  error {

	return s.repo.Update(ctx, se)

 }

 func (s *service) Delete(ctx context.Context, id int) error {
	
	return s.repo.Delete(ctx, id)
	
}
