package api

import (
	// "database/sql"
	// "log"
	"net/http"
	"time"

	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"
	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// new struct to store the create account request
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type  createUSerResponse struct {
	Username          string    `json:"username"`
	// HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}


func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorReponse(err))
		return
	}

	// hashed password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorReponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorReponse(err))
		return
	}

	rsp := createUSerResponse {
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
