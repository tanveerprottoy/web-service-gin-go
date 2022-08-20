package entities

import (
	"time"
	"txp/web-service-gin/src/core"
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