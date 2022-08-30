package content

import (
	"fmt"
	"log"
	"time"
	"txp/web-service-gin/src/data"
	"txp/web-service-gin/src/module/content/dto"
	"txp/web-service-gin/src/module/content/entity"
)

type ContentRepository struct {
}

func (r *ContentRepository) Create(d *dto.CreateUpdateContentBody) (
	int,
	error,
) {
	lastId := -1
	stmt, err := data.Db.PrepareNamed(
		"INSERT INTO contents (name)" +
			"VALUES (:name) RETURNING id",
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = stmt.Get(&lastId, d)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Println(err)
	}
	return lastId, nil
}

func (r *ContentRepository) FindAll() (
	[]entity.Content,
	error,
) {
	contents := []entity.Content{}
	err := data.Db.Select(
		&contents,
		"SELECT * FROM contents WHERE id > $1",
		fmt.Sprintf("%d", 0),
	)
	if err != nil {
		log.Println(err)
		return contents, err
	}
	return contents, nil
}

func (r *ContentRepository) FindOne(id string) (
	entity.Content,
	error,
) {
	u := entity.Content{}
	err := data.Db.Get(
		&u,
		"SELECT * FROM contents WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func (r *ContentRepository) Update(
	id string,
	u *entity.Content,
) (int64, error) {
	u.UpdatedAt = time.Now()
	q := "UPDATE Contents SET name=:name," +
		" updated_at = :updated_at WHERE id = " + id
	res, err := data.Db.NamedExec(
		q,
		u,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}

func (r *ContentRepository) Delete(id string) (
	int64,
	error,
) {
	q := "DELETE FROM contents WHERE id = $1"
	res, err := data.Db.Exec(
		q,
		id,
	)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}
