package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/andersjbe/quin/graph/model"
	"github.com/andersjbe/quin/internal/auth"
	"github.com/andersjbe/quin/internal/auth/jwt"
	"github.com/andersjbe/quin/internal/database"
)

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input model.CreateUser) (string, error) {
	hashedPass, err := auth.HashPassword(input.Password)
	if err != nil {
		return "", err
	}

	tx, err := r.Conn.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()
	qtx := r.DB.WithTx(tx)

	_, err = qtx.CreateUser(ctx, database.CreateUserParams{
		FirstName: input.FirstName,
		LastName:  *input.LastName,
		Email:     input.Email,
		Password:  hashedPass,
	})
	if err != nil {
		return "", err
	}
	user, err := qtx.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", err
	}
	_, err = qtx.CreateProfile(ctx, database.CreateProfileParams{
		Username:  input.Username,
		UUIDTOBIN: user.Uuid,
	})
	if err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user, err := r.DB.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return "", errors.New("email not found")
	}

	authed := auth.CheckPassword(input.Password, user.Password)
	if !authed {
		return "", errors.New("incorrect email or password")
	}

	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, token string) (string, error) {
	email, err := jwt.ParseToken(token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	newToken, err := jwt.GenerateToken(email)
	if err != nil {
		return "", err
	}
	return newToken, nil
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, input map[string]interface{}) (*model.User, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.User{}, errors.New("access denied")
	}

	userUpdate := database.UpdateUserParams{ID: string(user.ID)}
	profileUpdate := database.UpdateProfileParams{UserId: string(user.ID)}

	if gender, ok := input["gender"]; ok {
		profileUpdate.Gender.Valid = true
		switch model.Gender(gender.(string)) {
		case "MAN":
			profileUpdate.Gender.ProfilesGender = database.ProfilesGenderMan
		case "WOMAN":
			profileUpdate.Gender.ProfilesGender = database.ProfilesGenderWoman
		case "NONBINARY":
			profileUpdate.Gender.ProfilesGender = database.ProfilesGenderNonbinary
		}
	}
	if firstName, ok := input["firstName"].(string); ok {
		userUpdate.FirstName.Valid = true
		userUpdate.FirstName.String = firstName
	}
	if lastName, ok := input["lastName"].(string); ok {
		userUpdate.LastName.Valid = true
		userUpdate.LastName.String = lastName
	}
	if heightInches, ok := input["heightInches"].(int); ok {
		profileUpdate.HeightInches.Valid = true
		profileUpdate.HeightInches.Int32 = int32(heightInches)
	}
	if weightLbs, ok := input["weightLbs"].(int); ok {
		profileUpdate.WeightLbs.Valid = true
		profileUpdate.WeightLbs.Int32 = int32(weightLbs)
	}
	if skinPigmentationValue, ok := input["skinPigmentationValue"].(int); ok {
		profileUpdate.SkinPigmentation.Valid = true
		profileUpdate.SkinPigmentation.Int32 = int32(skinPigmentationValue)
	}
	if hairColor, ok := input["hairColor"].(string); ok {
		profileUpdate.HairColor.Valid = true
		profileUpdate.HairColor.String = hairColor
	}
	if eyeColor, ok := input["eyeColor"].(string); ok {
		profileUpdate.EyeColor.Valid = true
		profileUpdate.EyeColor.String = eyeColor
	}

	sizePreferences := model.ProductSize{}
	if bustInches, ok := input["bustInches"].(int); ok {
		sizePreferences.BustInches = &bustInches
	}
	if waistInches, ok := input["waistInches"].(int); ok {
		sizePreferences.WaistInches = &waistInches
	}
	if hipInches, ok := input["hipInches"].(int); ok {
		sizePreferences.HipInches = &hipInches
	}
	if pantLengthInches, ok := input["pantLengthInches"].(int); ok {
		sizePreferences.PantLengthInches = &pantLengthInches
	}
	if shoeSize, ok := input["shoeSize"].(int); ok {
		sizePreferences.ShoeSize = &shoeSize
	}
	var err error
	profileUpdate.SizePreferences, err = json.Marshal(sizePreferences)
	if err != nil {
		return &model.User{}, err
	}

	tx, err := r.Conn.Begin()
	if err != nil {
		return &model.User{}, err
	}
	defer tx.Rollback()
	qtx := r.DB.WithTx(tx)

	err = qtx.UpdateUser(ctx, userUpdate)
	if err != nil {
		return &model.User{}, err
	}
	err = qtx.UpdateProfile(ctx, profileUpdate)
	if err != nil {
		return &model.User{}, err
	}

	newUser, err := qtx.GetUserByID(ctx, string(user.ID))
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
		ID:        newUser.Uuid,
		FirstName: newUser.FirstName,
		LastName:  &newUser.LastName,
		Email:     newUser.Email,
	}, nil
}

// FollowProfile is the resolver for the followProfile field.
func (r *mutationResolver) FollowProfile(ctx context.Context, profileID string) (bool, error) {
	panic(fmt.Errorf("not implemented: FollowProfile - followProfile"))
}

// UnfollowProfile is the resolver for the unfollowProfile field.
func (r *mutationResolver) UnfollowProfile(ctx context.Context, profileID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnfollowProfile - unfollowProfile"))
}

// SubmitProduct is the resolver for the submitProduct field.
func (r *mutationResolver) SubmitProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Product{}, errors.New("access denied")
	}

	return &model.Product{}, nil
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, productID string, input map[string]interface{}) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// SavedProduct is the resolver for the savedProduct field.
func (r *mutationResolver) SavedProduct(ctx context.Context, productID string) (bool, error) {
	panic(fmt.Errorf("not implemented: SavedProduct - savedProduct"))
}

// UnsaveProduct is the resolver for the unsaveProduct field.
func (r *mutationResolver) UnsaveProduct(ctx context.Context, productID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnsaveProduct - unsaveProduct"))
}

// CreateOutfit is the resolver for the createOutfit field.
func (r *mutationResolver) CreateOutfit(ctx context.Context, input model.NewOutfit) (*model.Outfit, error) {
	panic(fmt.Errorf("not implemented: CreateOutfit - createOutfit"))
}

// UpdateOutfit is the resolver for the updateOutfit field.
func (r *mutationResolver) UpdateOutfit(ctx context.Context, outfitID string, input model.NewOutfit) (*model.Outfit, error) {
	panic(fmt.Errorf("not implemented: UpdateOutfit - updateOutfit"))
}

// LikeOutfit is the resolver for the likeOutfit field.
func (r *mutationResolver) LikeOutfit(ctx context.Context, outfitID string) (bool, error) {
	panic(fmt.Errorf("not implemented: LikeOutfit - likeOutfit"))
}

// UnlikeOutfit is the resolver for the unlikeOutfit field.
func (r *mutationResolver) UnlikeOutfit(ctx context.Context, outfitID string) (bool, error) {
	panic(fmt.Errorf("not implemented: UnlikeOutfit - unlikeOutfit"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
