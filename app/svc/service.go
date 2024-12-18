package svc

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"grimoire/app/model"
)

type Account interface {
	GetByID(context.Context, int) (*model.Account, error)
	// Auth
	SignIn(context.Context, *model.Account) (*model.Account, error)
	SignUp(context.Context, *model.Account) (*model.Account, error)
}

type Characters interface {
	GetCharByID(context.Context, int) (*model.Characters, error)
	GetCharByName(context.Context, string) (*model.Characters, error)
	GetCharTop10Kill(ctx context.Context) ([]map[string]interface{}, error)
}

type Sait interface {
	// Profile
	UpdateAvatar(context.Context, *model.Profile) (*model.Profile, error)

	// Session
	CreateCookie(string) *fiber.Cookie
	GenerateAccessToken(int) (string, error)
	GetSession(context.Context, *model.Session) (*model.Session, error)
	DeleteSession(context.Context, string) (*fiber.Cookie, error)
	ExistSession(context.Context, *model.Session) (bool, error)
	CreateSession(context.Context, *model.Session) error
	UpdateSession(context.Context, *model.Session) error
	UpdateOrCreateSession(context.Context, *model.Session, string) error
	HandleSession(context.Context, *model.Session) (string, error)

	// News
	GetNewsSlice(context.Context) (*model.NewsSlice, error)
	GetNewsByID(context.Context, *model.News) (*model.News, error)

	// Comments
	AddComment(context.Context, *model.Comment) (*model.Comment, error)
	DeleteComment(context.Context, int) error
	UpdateComment(context.Context, *model.Comment) error
	GetCommentByID(context.Context, int) (*model.Comment, error)
	GetCommentByNewsID(context.Context, int) (*model.CommentSlice, error)
}
