package signin

import (
	"testing"
	"github.com/golang/mock/gomock"
	"auth/src/persistance"
	"errors"
)


func TestModel_SignInShouldFailWhenCantAuthenticate(t *testing.T) {
	// Setup
	mockCtrl, _, method, model := setup(t)
	defer mockCtrl.Finish()
	// Test
	method.EXPECT().Auth(gomock.Any()).Return("", "", errors.New("error"))
	user, token, err := model.SignIn(method, "1")
	// Assert
	if err == nil {
		t.Fatal("Must throw an error")
	}
	if user != nil || token != "" {
		t.Fatal("Without Auth we should  return nothing")
	}
}

func setup(t *testing.T) (*gomock.Controller, *persistance.MockDB, *MockMethod, *Model) {
	mockCtrl := gomock.NewController(t)
	db := persistance.NewMockDB(mockCtrl)
	method := NewMockMethod(mockCtrl)
	model := NewModel(db)
	return mockCtrl, db, method, model
}