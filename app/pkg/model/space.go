package model

type SpaceVisibility string

const (
	SpaceVisibilityPrivate  SpaceVisibility = "PRIVATE"
	SpaceVisibilityInternal SpaceVisibility = "INTERNAL"
	SpaceVisibilityPublic   SpaceVisibility = "PUBLIC"
)

type Space struct {
	BaseModel
	ShortId     string          `db:"short_id" json:"shortId"`
	Name        string          `db:"name" json:"name"`
	Description *string         `db:"description" json:"description"`
	Visibility  SpaceVisibility `db:"visibility" json:"visibility"`
}

type CreateSpace struct {
	Name        string          `json:"name"`
	Description *string         `json:"description"`
	Visibility  SpaceVisibility `json:"visibility"`
}

type SpaceService interface {
	FindOne(id int64) (*Space, error)
	FindOneByShortId(shortId string) (*Space, error)
	FindAllForUser(user int64) ([]Space, error)
	Create(owner int64, params CreateSpace) (*Space, error)
}
