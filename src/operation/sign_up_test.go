package operation

//import (
//	"testing"
//	"github.com/golang/mock/gomock"
//	"auth/src/data"
//)


//func TestModel_SignInShouldFailWhenCantAuthenticate(t *testing.T) {
//	// Setup
//	mockCtrl, _, method, model := setup(t)
//	defer mockCtrl.Finish()
//	// Test
//	method.EXPECT().Auth(gomock.Any()).Return("", "", errors.New("error"))
//	user, token, err := model.SignUp(method, "1")
//	// Assert
//	if err == nil {
//		t.Fatal("Must throw an error")
//	}
//	if user != nil || token != "" {
//		t.Fatal("Without Exchange we should  return nothing")
//	}
//}

//func setup(t *testing.T) (*gomock.Controller, *data.MockDB, *MockMethod, *Persistent) {
//	mockCtrl := gomock.NewController(t)
//	db := data.NewMockDB(mockCtrl)
//	method := NewMockMethod(mockCtrl)
//	model := NewPersistent(db)
//	return mockCtrl, db, method, model
//}