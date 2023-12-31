// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type ProductCategory string

const (
	ProductCategoryHeadwear    ProductCategory = "headwear"
	ProductCategoryEyewear     ProductCategory = "eyewear"
	ProductCategoryTops        ProductCategory = "tops"
	ProductCategoryBottoms     ProductCategory = "bottoms"
	ProductCategoryDresses     ProductCategory = "dresses"
	ProductCategoryOuterwear   ProductCategory = "outerwear"
	ProductCategoryAccessories ProductCategory = "accessories"
	ProductCategoryShoes       ProductCategory = "shoes"
	ProductCategorySwimwear    ProductCategory = "swimwear"
	ProductCategoryPiercings   ProductCategory = "piercings"
	ProductCategoryNecklaces   ProductCategory = "necklaces"
	ProductCategoryRings       ProductCategory = "rings"
)

func (e *ProductCategory) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProductCategory(s)
	case string:
		*e = ProductCategory(s)
	default:
		return fmt.Errorf("unsupported scan type for ProductCategory: %T", src)
	}
	return nil
}

type NullProductCategory struct {
	ProductCategory ProductCategory
	Valid           bool // Valid is true if ProductCategory is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProductCategory) Scan(value interface{}) error {
	if value == nil {
		ns.ProductCategory, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProductCategory.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProductCategory) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProductCategory), nil
}

type ProductGender string

const (
	ProductGenderMale   ProductGender = "male"
	ProductGenderFemale ProductGender = "female"
	ProductGenderUnisex ProductGender = "unisex"
)

func (e *ProductGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProductGender(s)
	case string:
		*e = ProductGender(s)
	default:
		return fmt.Errorf("unsupported scan type for ProductGender: %T", src)
	}
	return nil
}

type NullProductGender struct {
	ProductGender ProductGender
	Valid         bool // Valid is true if ProductGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProductGender) Scan(value interface{}) error {
	if value == nil {
		ns.ProductGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProductGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProductGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProductGender), nil
}

type ProfileGender string

const (
	ProfileGenderMan       ProfileGender = "man"
	ProfileGenderWoman     ProfileGender = "woman"
	ProfileGenderNonbinary ProfileGender = "nonbinary"
)

func (e *ProfileGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProfileGender(s)
	case string:
		*e = ProfileGender(s)
	default:
		return fmt.Errorf("unsupported scan type for ProfileGender: %T", src)
	}
	return nil
}

type NullProfileGender struct {
	ProfileGender ProfileGender
	Valid         bool // Valid is true if ProfileGender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProfileGender) Scan(value interface{}) error {
	if value == nil {
		ns.ProfileGender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProfileGender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProfileGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProfileGender), nil
}

type Outfit struct {
	ID          uuid.UUID
	CreatedAt   sql.NullTime
	Description sql.NullString
	ImageUrls   json.RawMessage
	ProfileID   interface{}
}

type OutfitLike struct {
	CreatedAt sql.NullTime
	OutfitID  uuid.UUID
	ProfileID uuid.UUID
}

type OutfitPiece struct {
	CreatedAt sql.NullTime
	Sizes     json.RawMessage
	OutfitID  uuid.UUID
	VariantID uuid.UUID
	Key       interface{}
}

type Product struct {
	ID             uuid.UUID
	CreatedAt      sql.NullTime
	Name           string
	Description    sql.NullString
	Sourceurl      sql.NullString
	AvailableToBuy bool
	Gender         NullProductGender
	Categories     NullProductCategory
	Materials      json.RawMessage
	ImageUrl       string
	ProfileID      uuid.UUID
}

type ProductSafe struct {
	CreatedAt sql.NullTime
	ProductID uuid.UUID
	ProfileID uuid.UUID
}

type ProductVariant struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	Name      string
	Colors    json.RawMessage
	Pattern   string
	ImageUrls pqtype.NullRawMessage
	ProductID uuid.UUID
}

type Profile struct {
	ID               uuid.UUID
	CreatedAt        sql.NullTime
	Username         string
	ImageUrl         string
	Gender           NullProfileGender
	HeightInches     sql.NullInt32
	WeightLbs        sql.NullInt32
	SkinPigmentation sql.NullInt32
	HairColor        sql.NullString
	EyeColor         sql.NullString
	SizePreferences  pqtype.NullRawMessage
	UserID           uuid.UUID
}

type ProfileFollow struct {
	CreatedAt   sql.NullTime
	FollowerID  uuid.UUID
	FollowingID uuid.UUID
	Key         interface{}
}

type User struct {
	ID        uuid.UUID
	CreatedAt sql.NullTime
	FirstName string
	LastName  string
	Email     string
	Password  string
}
