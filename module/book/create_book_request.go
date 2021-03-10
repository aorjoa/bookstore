package book

type CreateBookRequest struct {
	Name     *string `json:"name"`
	ISBN     *string `json:"isbn"`
	Language *string `json:"language"`
	Status   *bool   `json:"status"`
}
