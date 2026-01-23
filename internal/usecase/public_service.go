package usecase

type PublicService struct {

}

type HealthCheck struct {
	Status string `json:"status"`
}

func NewPublicService() *PublicService {
	return &PublicService{}
}

func (p *PublicService) HealthCheck() HealthCheck {
	data := HealthCheck{
			Status: "ok",
	}

	return data
}