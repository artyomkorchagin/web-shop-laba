package category

type Service struct {
	repo ReadWriter
}

func NewService(repo ReadWriter) *Service {
	return &Service{
		repo: repo,
	}
}
