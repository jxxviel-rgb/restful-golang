package simple

type BarRepository struct {
}
type BarService struct {
	*BarRepository
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}
func NewBarService(barRepository *BarRepository) *BarService {
	return &BarService{BarRepository: barRepository}
}
