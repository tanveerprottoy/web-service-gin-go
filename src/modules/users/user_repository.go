package users

import (
	"fmt"
	"txp/web-service-gin/src/data"
	"txp/web-service-gin/src/modules/users/entities"
)

type UserRepositoy struct{}

func Create(l *LoginBody) int {
	lastId := 0
	stmt, err := datum.Db.PrepareNamed(
		"INSERT INTO users (phone)" +
			"VALUES (:phone) RETURNING id",
	)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Get(&lastId, l)
	if err != nil {
		log.Println(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Println(err)
	}
	return lastId
}

func (repository *UserRepositoy) FindAll() []entities.User {
	users := []entities.User{}
	_ = data.Db.Select(
		&users,
		"SELECT * FROM users WHERE id > ? AND isDeleted = ?"+
			" ORDER BY createdAt ASC",
		fmt.Sprintf("%d", 0),
		fmt.Sprintf("%t", false),
	)
	return users
}

func FindOne(
	id int,
) (entities.User, error) {
	u := entities.User{}
	var err error
	err = data.Db.Get(
		&u,
		"SELECT * FROM users WHERE id = $1 LIMIT 1",
		id,
	)
	return u, err
}

func UpdateUser(u *User) int64 {
	var query string
	u.UpdatedAt = time.Now()
	if u.Password != "" {
		u.PasswordHash.String = core.GenerateHashFromPassword(u.Password)
		query = "UPDATE users SET updated_at = :updated_at," +
			" first_name=:first_name.string, last_name=:last_name.string, email=:email.string," +
			" password_hash=:password_hash.string, nid_number=:nid_number.string, age=:age.int64, dob=:dob.time, height=:height.float64, weight=:weight.int64," +
			" created_by=:created_by.string, religion=:religion.string, short_bio=:short_bio.string," +
			" gender=:gender.string, hometown=:hometown.string, education=:education.string, profession=:profession.string, address_current=:address_current.string," +
			" address_permanent=:address_permanent.string, country=:country.string, father_name=:father_name.string, " +
			" mother_name=:mother_name.string, sibling_count=:sibling_count.int64, family_type=:family_type.string, " +
			" blood_group=:blood_group.string, marital_status=:marital_status.string, image_url=:image_url.string," +
			" is_registered=True WHERE id = " + fmt.Sprintf("%d", u.Id)
	} else if u.ImageUrl.String == "" {
		query = "UPDATE users SET updated_at = :updated_at," +
			" first_name=:first_name.string, last_name=:last_name.string, email=:email.string," +
			" age=:age.int64, dob=:dob.time, height=:height.float64, weight=:weight.int64, religion=:religion.string, short_bio=:short_bio.string," +
			" gender=:gender.string, hometown=:hometown.string, education=:education.string, profession=:profession.string, address_current=:address_current.string," +
			" address_permanent=:address_permanent.string, country=:country.string, father_name=:father_name.string," +
			" mother_name=:mother_name.string, sibling_count=:sibling_count.int64, family_type=:family_type.string, " +
			" blood_group=:blood_group.string, marital_status=:marital_status.string," +
			" is_registered=True WHERE id = " + fmt.Sprintf("%d", u.Id)
	} else {
		query = "UPDATE users SET updated_at = :updated_at," +
			" first_name=:first_name.string, last_name=:last_name.string, email=:email.string," +
			" age=:age.int64, dob=:dob.time, height=:height.float64, weight=:weight.int64, religion=:religion.string, short_bio=:short_bio.string," +
			" gender=:gender.string, hometown=:hometown.string, education=:education.string, profession=:profession.string, address_current=:address_current.string," +
			" address_permanent=:address_permanent.string, country=:country.string, father_name=:father_name.string," +
			" mother_name=:mother_name.string, sibling_count=:sibling_count.int64, family_type=:family_type.string, " +
			" blood_group=:blood_group.string, marital_status=:marital_status.string, image_url=:image_url.string," +
			" is_registered=True WHERE id = " + fmt.Sprintf("%d", u.Id)
	}
	r, err := datum.Db.NamedExec(
		query,
		u,
	)
	if err != nil {
		log.Println(err)
		return 0
	}
	return GetRowsAffected(r)
}
