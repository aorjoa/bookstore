package book

type BookResponse struct {
	ID       *uint   `json:"id"`
	Name     *string `json:"name"`
	ISBN     *string `json:"isbn"`
	Language *string `json:"language"`
	Status   *bool   `json:"status"`
}
