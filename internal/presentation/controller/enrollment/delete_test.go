package controller_test

import (
	"fmt"
	"net/http"
	"testing"

	service "github.com/christian-gama/pd-solucoes/internal/app/service/enrollment"
	controller "github.com/christian-gama/pd-solucoes/internal/presentation/controller/enrollment"
	fake "github.com/christian-gama/pd-solucoes/testutils/fake/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/gintest"
	mocks "github.com/christian-gama/pd-solucoes/testutils/mocks/app/service/enrollment"
	"github.com/christian-gama/pd-solucoes/testutils/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DeleteCourseEnrollmentSuite struct {
	suite.Suite
}

func TestDeleteCourseEnrollmentSuite(t *testing.T) {
	suite.RunUnitTest(t, new(DeleteCourseEnrollmentSuite))
}

func (s *DeleteCourseEnrollmentSuite) TestHandle() {
	type Sut struct {
		Sut                    controller.DeleteCourseEnrollment
		Input                  *service.DeleteCourseEnrollmentInput
		DeleteCourseEnrollment *mocks.DeleteCourseEnrollment
	}

	makeSut := func() *Sut {
		input := fake.DeleteCourseEnrollmentInput()
		deleteCourseEnrollment := new(mocks.DeleteCourseEnrollment)
		sut := controller.NewDeleteCourseEnrollment(deleteCourseEnrollment)
		return &Sut{Sut: sut, DeleteCourseEnrollment: deleteCourseEnrollment, Input: input}
	}

	s.Run("should find one courseEnrollment", func() {
		sut := makeSut()

		sut.DeleteCourseEnrollment.On("Handle", mock.Anything, sut.Input).Return(nil)

		ctx := gintest.MustRequest(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusNoContent, ctx.Writer.Status())
		sut.DeleteCourseEnrollment.AssertCalled(s.T(), "Handle", mock.Anything, sut.Input)
	})

	s.Run("invalid ID: it's required", func() {
		sut := makeSut()

		sut.Input.ID = 0

		ctx, _ := gintest.MustRequestWithBody(sut.Sut, gintest.Option{
			Params: []string{fmt.Sprint(sut.Input.ID)},
		})

		s.Equal(http.StatusBadRequest, ctx.Writer.Status())
	})

	s.Run("panics when DeleteCourseEnrollment.Handle returns error", func() {
		sut := makeSut()

		sut.DeleteCourseEnrollment.On("Handle", mock.Anything, sut.Input).Return(assert.AnError)

		s.Panics(func() {
			gintest.MustRequest(sut.Sut, gintest.Option{
				Params: []string{fmt.Sprint(sut.Input.ID)},
			})
		})
	})
}
