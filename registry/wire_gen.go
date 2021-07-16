// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package registry

import (
	"github.com/QMDAKA/comment-mock/app/api"
	"github.com/QMDAKA/comment-mock/auth"
	comment2 "github.com/QMDAKA/comment-mock/handler/rest/comment"
	"github.com/QMDAKA/comment-mock/infrastructure/store/mysql"
	"github.com/QMDAKA/comment-mock/middleware"
	"github.com/QMDAKA/comment-mock/service/comment"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeServer(db *gorm.DB) (api.Server, error) {
	mysqlComment := mysql.ProvideCommentRepo(db)
	transaction := mysql.ProvideTransaction(db)
	user := mysql.ProvideUserRepo(db)
	authAuth := auth.NewAuth(user)
	commentComment := comment.NewComment(mysqlComment, transaction, authAuth)
	index := comment2.NewCommentIndex(commentComment)
	create := comment2.NewCommentCreate(commentComment)
	handlerCollection := api.NewHandlerCollection(index, create)
	middlewareAuth := middleware.ProvideAuth(authAuth)
	server := api.NewServer(db, handlerCollection, middlewareAuth)
	return server, nil
}
