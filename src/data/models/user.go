package model

import (
	"fmt"
	"log"
	"time"

	"biyelap.com/biyelap-core/app/core"
	"biyelap.com/biyelap-core/app/datum"
)

type User struct {
	CreatedAt        time.Time        `db:"created_at" json:"createdAt,omitempty"`
	UpdatedAt        time.Time        `db:"updated_at" json:"updatedAt,omitempty"`
	Id               int              `db:"id" json:"id"`
	Phone            string           `db:"phone" json:"phone,omitempty"`
	FirstName        core.NullString  `db:"first_name" json:"firstName,omitempty"`
	LastName         core.NullString  `db:"last_name" json:"lastName,omitempty"`
	Email            core.NullString  `db:"email" json:"email,omitempty"`
	Password         string           `db:"-" json:"password"`
	PasswordHash     core.NullString  `db:"password_hash" json:"-"`
	NidNumber        core.NullString  `db:"nid_number" json:"nidNumber,omitempty"`
	Age              core.NullInt64   `db:"age" json:"age,omitempty"`
	Dob              core.NullTime    `db:"dob" json:"dob,omitempty"`
	Height           core.NullFloat64 `db:"height" json:"height,omitempty"`
	Weight           core.NullInt64   `db:"weight" json:"weight,omitempty"`
	CreatedBy        core.NullString  `db:"created_by" json:"created_by,omitempty"`
	Religion         core.NullString  `db:"religion" json:"religion,omitempty"`
	ShortBio         core.NullString  `db:"short_bio" json:"shortBio,omitempty"`
	Gender           core.NullString  `db:"gender" json:"gender,omitempty"`
	Hometown         core.NullString  `db:"hometown" json:"hometown,omitempty"`
	Education        core.NullString  `db:"education" json:"education,omitempty"`
	Profession       core.NullString  `db:"profession" json:"profession,omitempty"`
	AddressCurrent   core.NullString  `db:"address_current" json:"address_current,omitempty"`
	AddressPermanent core.NullString  `db:"address_permanent" json:"address_permanent,omitempty"`
	Country          core.NullString  `db:"country" json:"country,omitempty"`
	BloodGroup       core.NullString  `db:"blood_group" json:"bloodGroup,omitempty"`
	MaritalStatus    core.NullString  `db:"marital_status" json:"maritalStatus,omitempty"`
	ImageUrl         core.NullString  `db:"image_url" json:"imageUrl,omitempty"`
	FatherName       core.NullString  `db:"father_name" json:"fatherName,omitempty"`
	MotherName       core.NullString  `db:"mother_name" json:"motherName,omitempty"`
	SiblingCount     core.NullInt64   `db:"sibling_count" json:"siblingCount"`
	FamilyType       core.NullString  `db:"family_type" json:"familyType,omitempty"`
	IsActive         bool             `db:"is_active" json:"isActive,omitempty"`
	IsRegistered     bool             `db:"is_registered" json:"isRegistered,omitempty"`
}

func GetUsers(
	id int,
	lastId int,
	gender string,
	religion string,
	maritalStatus string,
) []User {
	users := []User{}
	if gender == "" && religion == "" && maritalStatus == "" {
		_ = datum.Db.Select(
			&users,
			"SELECT id, first_name, last_name, age FROM users WHERE NOT id = ?"+
				"AND id > ? ORDER BY first_name ASC",
			fmt.Sprintf("%d", id),
			fmt.Sprintf("%d", lastId),
		)
	} else {
		_ = datum.Db.Select(
			&users,
			"SELECT id, first_name, last_name, age FROM users WHERE NOT id = "+
				" AND id > ? AND gender = ?"+" AND religion = ? AND marital_status = ?"+
				" ORDER BY first_name ASC",
			fmt.Sprintf("%d", id),
			fmt.Sprintf("%d", lastId),
			gender,
			religion,
			maritalStatus,
		)
	}
	return users
}

func GetUser(
	id int,
	isAllCol bool,
) (User, error) {
	u := User{}
	var err error
	if isAllCol {
		err = datum.Db.Get(
			&u,
			"SELECT * FROM users WHERE id = $1 LIMIT 1",
			id,
		)
	} else {
		err = datum.Db.Get(
			&u,
			"SELECT id, first_name, last_name, age, religion, education, marital_status, blood_group, dob FROM users WHERE id = $1 LIMIT 1",
			id,
		)
	}
	return u, err
}

func GetUserByPhone(phone string) (User, error) {
	u := User{}
	err := datum.Db.Get(&u, "SELECT id, is_active, is_registered FROM users WHERE phone = $1 LIMIT 1", phone)
	return u, err
}

func GetUserPassHash(id int) (string, error) {
	u := User{}
	err := datum.Db.Get(&u, "SELECT password_hash FROM users WHERE id = $1 LIMIT 1", id)
	return u.PasswordHash.String, err
}

func InsertUser(l *LoginBody) int {
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
