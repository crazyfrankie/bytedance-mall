package util

import "context"

func GetUserID(ctx context.Context) uint32 {
	userId := ctx.Value("user_id")
	if userId == nil {
		return 0
	}

	return userId.(uint32)
}
