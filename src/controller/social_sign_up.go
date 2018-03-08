package controller

import "auth/src/pojo"

type Result struct {
	User *pojo.User `json:"user"`
	Token string `json:"token"`
}

// Return type based on Formatter Template
// TODO Create custom error type
func (ctrl ProviderPersistence) SocialSignUp(token string) (string) {
	user, err := ctrl.prov.GetUserFromProvider(token)
	if err != nil {
		return ctrl.format.Format(err)
	}
	user, err = ctrl.persist.SignUp(user)
	// Erases refreshToken after it is already saved
	user.RefreshToken = ""
	if err != nil {
		return ctrl.format.Format(err)
	}
	// TODO Use JWT
	result := &Result{
		User:user,
		Token:"",
	}
	return ctrl.format.Format(result)
}