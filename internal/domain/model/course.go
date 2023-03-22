package model

import (
	"errors"

	"github.com/christian-gama/pd-solucoes/pkg/errutil"
)

// Course is a model that represents a course of a college.
type Course struct {
	ID        uint   `faker:"uint"`
	Name      string `faker:"len=50"`
	CollegeID uint   `faker:"uint"`
	College   *College
}

// NewCourse creates a new Course.
func NewCourse(id uint, name string, collegeID uint, college *College) (*Course, error) {
	m := &Course{
		ID:        id,
		Name:      name,
		CollegeID: collegeID,
		College:   college,
	}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	return m, nil
}

// Validate validates the Course.
func (m *Course) Validate() error {
	var errs *errutil.Error

	if m.Name == "" {
		errs = errutil.Append(errs, errors.New("name is required"))
	}

	if m.CollegeID == 0 {
		errs = errutil.Append(errs, errors.New("collegeID is required"))
	}

	if m.College == nil {
		errs = errutil.Append(errs, errors.New("college is required"))
	} else if err := m.College.Validate(); err != nil {
		errs = errutil.Append(errs, errors.New("college is invalid"))
	}

	if errs.HasErrors() {
		return errs
	}

	return nil
}
