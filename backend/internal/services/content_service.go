package services

/*
type ContentService struct {
	spaceRepo    *db.SpaceRepository
	pageRepo     *db.PageRepository
	bookmarkRepo *db.BookmarkRepository
}

func NewContentService(
	spaceRepo *db.SpaceRepository,
	pageRepo *db.PageRepository,
	bookmarkRepo *db.BookmarkRepository) *ContentService {
	return &ContentService{
		spaceRepo:    spaceRepo,
		pageRepo:     pageRepo,
		bookmarkRepo: bookmarkRepo,
	}
}

func (s *ContentService) CreateSpace(ctx context.Context, space *internal.Space) error {
	return s.spaceRepo.Save(ctx, space)
}

func (s *ContentService) GetSpaceById(ctx context.Context, id bson.ObjectID) (internal.Space, error) {
	space, err := s.spaceRepo.FindByID(ctx, id)
	if err != nil {
		return internal.Space{}, err
	}

	pages, err := s.pageRepo.FindBySpaceId(ctx, space.ID)
	if err != nil {
		return internal.Space{}, err
	}

	space.Pages = make([]bson.ObjectID, len(pages))
	for i, page := range pages {
		space.Pages[i] = page.ID
	}
	return *space, nil
}

func (s *ContentService) GetPagesBySpaceId(ctx context.Context, spaceId bson.ObjectID) ([]internal.Page, error) {
	return s.pageRepo.FindBySpaceId(ctx, spaceId)
}

func (s *ContentService) BuildPageTree(pages []*internal.Page) []*internal.Page {
	return s.pageRepo.BuildPageTree(pages)
}

func (s *ContentService) GetSpacesForUser(userId bson.ObjectID) ([]*internal.Space, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"ownerId": userId},
			{"shared.userId": userId},
		},
	}

	founds, err := s.spaceRepo.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return founds, nil
}

func (s *ContentService) DeleteSpace(ctx context.Context, id bson.ObjectID) error {
	return s.spaceRepo.Delete(ctx, id)
}
*/
