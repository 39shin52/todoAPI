package response

import "github.com/39shin52/todoAPI/app/domain/entity"

type SelectUserResponse struct {
	ID   string `json:"user_id"`
	Mail string `json:"mail"`
	Work string `json:"work"`
}

type SelectUsersResponse struct {
	Users []entity.User `json:"users"`
}

func UsersResponse(users []entity.User) SelectUsersResponse {
	sur := SelectUsersResponse{
		Users: users,
	}

	return sur
}
