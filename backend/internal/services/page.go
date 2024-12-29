package services

import (
	"backend/internal/db/sqlc"
	"context"
	"time"
)

type PageService struct {
	queries *sqlc.Queries
}

// Hierarchisches Struct für die Darstellung
type HierarchicalPage struct {
	ID          int64
	SpaceID     int64
	Name        string
	Description *string
	Visibility  string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	SubPages    []*HierarchicalPage // Für Unterseiten
}

func NewPageService(queries *sqlc.Queries) *PageService {
	return &PageService{queries: queries}
}

func (s *PageService) GetPagesBySpaceId(spaceId int64, rootOnly bool) ([]sqlc.Page, error) {
	if rootOnly {
		return s.queries.FindRootPagesBySpaceId(context.TODO(), spaceId)
	} else {
		return s.queries.FindPagesBySpaceId(context.TODO(), spaceId)
	}
}

// MapPagesToHierarchy wandelt eine flache Liste von Pages in eine hierarchische Struktur um
func MapPagesToHierarchy(pages []sqlc.Page) []*HierarchicalPage {
	// Map für schnelle Lookups
	pageMap := make(map[int64]*HierarchicalPage)

	// Erzeuge HierarchicalPages aus Pages
	var roots []*HierarchicalPage
	for _, page := range pages {
		hPage := &HierarchicalPage{
			ID:          page.ID,
			SpaceID:     page.SpaceID,
			Name:        page.Name,
			Description: page.Description,
			Visibility:  page.Visibility,
			CreatedAt:   page.CreatedAt,
			UpdatedAt:   page.UpdatedAt,
			SubPages:    []*HierarchicalPage{},
		}
		pageMap[page.ID] = hPage

		// Falls keine ParentID vorhanden ist, ist es eine Root-Seite
		if page.ParentID == nil {
			roots = append(roots, hPage)
		}
	}

	// Ordne die Seiten den Eltern zu
	for _, page := range pages {
		if page.ParentID != nil {
			parent, exists := pageMap[*page.ParentID]
			if exists {
				parent.SubPages = append(parent.SubPages, pageMap[page.ID])
			}
		}
	}

	return roots
}
