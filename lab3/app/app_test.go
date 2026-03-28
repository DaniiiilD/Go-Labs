package app

import (
	"errors"
	"testing"

	"lab3/app/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	type Test struct {
		Name        string
		SetupMock   func(db *mocks.MockDB)
		ExpectedMsg string
		ExpectedErr error
	}

	errDB := errors.New("database error")

	tests := []Test{
		{
			Name: "Success",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil).Times(1)
				db.EXPECT().Insert("hello fyne from lab3!").Return(nil).Times(1)
				db.EXPECT().GetFirst().Return("hello fyne from lab3!", nil).Times(1)
			},
			ExpectedMsg: "hello fyne from lab3!",
			ExpectedErr: nil,
		},
		{
			Name: "Error at CreateTable",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(errDB).Times(1)
			},
			ExpectedMsg: "",
			ExpectedErr: errDB,
		},
		{
			Name: "Error at Insert",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil).Times(1)
				db.EXPECT().Insert("hello fyne from lab3!").Return(errDB).Times(1)
			},
			ExpectedMsg: "",
			ExpectedErr: errDB,
		},
		{
			Name: "Error at GetFirst",
			SetupMock: func(db *mocks.MockDB) {
				db.EXPECT().CreateTable().Return(nil).Times(1)
				db.EXPECT().Insert("hello fyne from lab3!").Return(nil).Times(1)
				db.EXPECT().GetFirst().Return("", errDB).Times(1)
			},
			ExpectedMsg: "",
			ExpectedErr: errDB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDB := mocks.NewMockDB(ctrl)
			tt.SetupMock(mockDB)

			a := New(mockDB)
			msg, err := a.Run()

			require.ErrorIs(t, err, tt.ExpectedErr)
			require.Equal(t, tt.ExpectedMsg, msg)
		})
	}
}