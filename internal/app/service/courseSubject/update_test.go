package service_test

import (
	"context"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/courseSubject"
	"github.com/christian-gama/pd-solucoes/internal/domain/model"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/courseSubject"
	fakeModel "github.com/christian-gama/pd-solucoes/testutils/fake/domain/model"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/domain/repo"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UpdateCourseSubjectSuite struct {
	suite.Suite
}

func TestUpdateCourseSubjectSuite(t *testing.T) {
	suite.RunUnitTest(t, new(UpdateCourseSubjectSuite))
}

func (s *UpdateCourseSubjectSuite) TestHandle() {
	type Sut struct {
		Sut               service.UpdateCourseSubject
		CourseSubjectRepo *mocks.CourseSubject
		Input             *service.UpdateInput
		CourseSubject     *model.CourseSubject
	}

	makeSut := func() *Sut {
		courseSubjectRepo := mocks.NewCourseSubject(s.T())
		sut := service.NewUpdateCourseSubject(courseSubjectRepo)

		return &Sut{
			Sut:               sut,
			CourseSubjectRepo: courseSubjectRepo,
			Input:             fake.UpdateCourseSubjectInput(),
			CourseSubject:     fakeModel.CourseSubject(),
		}
	}

	s.Run("should add update a courseSubject", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("Update", mock.Anything, mock.Anything).
			Return(sut.CourseSubject, nil)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.NoError(err)
		s.Equal(sut.CourseSubject.ID, result.ID)
		s.Equal(sut.CourseSubject.SubjectID, result.SubjectID)
	})

	s.Run("courseSubjectRepo.Update returns an error", func() {
		sut := makeSut()

		sut.CourseSubjectRepo.
			On("Update", mock.Anything, mock.Anything).
			Return(nil, assert.AnError)

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.ErrorIs(err, assert.AnError)
		s.Nil(result)
	})

	s.Run("should return an error if domain validation fails", func() {
		sut := makeSut()

		sut.Input.SubjectID = 0

		result, err := sut.Sut.Handle(context.Background(), sut.Input)

		s.Error(err)
		s.Nil(result)
	})
}
