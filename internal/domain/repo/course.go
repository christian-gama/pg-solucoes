package repo

import (
	"context"

	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	"github.com/christian-gama/pd-solucoes/internal/domain/querying"
)

type (
	CreateCourseParams struct {
		Course *model.Course
	}

	UpdateCourseParams struct {
		Course *model.Course
	}

	DeleteCourseParams struct {
		ID uint
	}

	FindAllCourseParams struct {
		Filterer  querying.Filterer
		Paginator querying.Paginator
		Sorter    querying.Sorter
	}

	FindCourseByIDParams struct {
		ID uint

		Filterer querying.Filterer
	}
)

type Course interface {
	// Create creates a new course.
	Create(
		ctx context.Context,
		params *CreateCourseParams,
	) (*model.Course, error)

	// Delete deletes a course by its ID.
	Delete(
		ctx context.Context,
		params *DeleteCourseParams,
	) error

	// FindAll returns a list of courses.
	FindAll(
		ctx context.Context,
		params *FindAllCourseParams,
		preload ...string,
	) (*querying.PaginationOutput[*model.Course], error)

	// FindByID returns a course by its ID.
	FindByID(
		ctx context.Context,
		params *FindCourseByIDParams,
		preload ...string,
	) (*model.Course, error)

	// Update updates a course.
	Update(
		ctx context.Context,
		params *UpdateCourseParams,
	) (*model.Course, error)
}