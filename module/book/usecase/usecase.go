package usecase

type bookUseCase struct {
}

func NewBookUseCase() BookUseCase {
	return &bookUseCase{}
}
