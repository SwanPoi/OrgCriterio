package usecase

type Repository interface {

}

type Usecase interface {

}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
