package util

import "auth/src/kind"

func PipeUserOps(ops []kind.UserOp, user *kind.User) (*kind.User, error) {
	for _, op := range ops {
		user, err := op(user)
		if err != nil {
			return user, err
		}
	}
	return user, nil
}
