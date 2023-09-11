package service


import (
	"time"
	"context"
	"database/sql"
	"github.com/msalbrain/authSphere/internals/database"
)

type UserService interface {
	// Returns User by id
	GetUserById(id int64) (database.User, error)
 	// Same as idan ☝🏻, retrieves single user by
	FindUserByEmail(email string) (database.User, error)
	CreateNewEmailUser(name, email, bio, passport string) (database.User, error)
	ValidatePassword(hashedpassword, unhashedpassword string)  bool

}




type UserSqlService struct {
	*database.Queries
	Config
}



func NewUserService(q *database.Queries, env Config) *UserSqlService {
	return &UserSqlService{q, env}
}


func (q *UserSqlService) GetPasswordHashCost() int64 {
	return q.Config.PasswordCost
}

func (q *UserSqlService) ValidatePassword(hashedpassword, unhashedpassword string)  bool {
	v := verifyPassword(hashedpassword, unhashedpassword)
	if v!= nil {
		return false
	}else {
		return true
	}
}

func (q *UserSqlService) GetUserById(id int64) (database.User, error){
	ctx := context.Background()
	return q.GetUser(ctx, id)
}

func (q *UserSqlService) FindUserByEmail(email string) (database.User, error){
	ctx := context.Background()
	emailString := sql.NullString{email, true}

	return q.GetUserByEmail(ctx, emailString)
}

func (q *UserSqlService) CreateNewEmailUser(name, email, bio, password string) (database.User, error) {
	ctx := context.Background()

    bioValid := true
    if bio == "" {
    } else {
        bioValid = false
    }

	hashpass, err := hashPassword(password)
	if err != nil {
		panic(err)
	}

	credTime := time.Now().Unix()
	return q.CreateUser(ctx, database.CreateUserParams{
		Name: name,
		Email: sql.NullString{String: email, Valid: true},
		HashedPassword: sql.NullString{String: hashpass, Valid: true},
		Bio: sql.NullString{String: bio, Valid: bioValid},
		AuthToken: GenerateToken(string(rune(credTime)), name, email),
		CreatedAt: credTime,
		UpdatedAt:  sql.NullInt64{Int64: credTime, Valid: true},
	})

}


