package user

import (
	"log"
	"time"
	"txp/web-service-gin/src/data"
	"txp/web-service-gin/src/module/user/dto"
	"txp/web-service-gin/src/module/user/entity"
)

type UserRepository struct {
}

func (r *UserRepository) Create(p *dto.CreateUpdateUserBody) int {
	lastId := -1
	stmt, err := data.Db.PrepareNamed(
		"INSERT INTO users (name)" +
			"VALUES (:name) RETURNING id",
	)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Get(&lastId, p)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Println(err)
	}
	return lastId
}

func (r *UserRepository) FindAll() []entity.User {
	users := []entity.User{}
	err := data.Db.Select(
		&users,
		"SELECT * FROM users", // "SELECT * FROM users WHERE NOT id > ? ORDER BY created_at ASC",
		// fmt.Sprintf("%d", 0),
	)
	if err != nil {
		log.Println(err)
	}
	return users
}

func (r *UserRepository) FindOne(id string) (entity.User, error) {
	u := entity.User{}
	err := data.Db.Get(
		&u,
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func (r *UserRepository) Update(
	id string,
	u *entity.User,
) int64 {
	u.UpdatedAt = time.Now()
	query := "UPDATE users SET name=:name," +
				" updated_at = :updated_at WHERE id = " + id
	res, err := data.Db.NamedExec(
		query,
		u,
	)
	if err != nil {
		log.Println(err)
		return 0
	}
	return data.GetRowsAffected(res)
}

func (r *UserRepository) Delete(id string) int64 {
	return 0
}
