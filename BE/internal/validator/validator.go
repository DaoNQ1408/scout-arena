package validator

import (
	arenaModel "scout-arena/internal/arena/model"
	userModel "scout-arena/internal/user/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterValidation("arena_status", func(fl validator.FieldLevel) bool {
			status := arenaModel.ArenaStatus(fl.Field().String())
			return status.IsValid()
		})

		v.RegisterValidation("team_status", func(fl validator.FieldLevel) bool {
			status := userModel.TeamStatus(fl.Field().String())
			return status.IsValid()
		})

		v.RegisterValidation("user_status", func(fl validator.FieldLevel) bool {
			status := userModel.UserStatus(fl.Field().String())
			return status.IsValid()
		})
	}
}
