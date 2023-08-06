-- name: CreateProfile :execlastid
INSERT INTO profiles (username, image_url, user_id)
VALUES ($1, "https://cdn.onlinewebfonts.com/svg/img_110805.png", $2);

-- name: GetProfileByID :one
SELECT u.id, u.first_name, u.last_name, p.*
FROM profiles p
JOIN users u ON p.user_id=u.id
WHERE p.id = $1
LIMIT 1;

-- name: GetProfileByUserID :one
SELECT u.id, u.first_name, u.last_name, p.*
FROM profiles p
JOIN users u ON p.user_id=u.id
WHERE p.user_id = $1
LIMIT 1;

-- name: UpdateProfile :exec
UPDATE profiles
SET
  image_url=coalesce(sqlc.narg('image_url'), image_url),
  gender=coalesce(sqlc.narg('gender'), gender),
  height_inches=coalesce(sqlc.narg('height_inches'), height_inches),
  weight_lbs=coalesce(sqlc.narg('weight_lbs'), weight_lbs),
  skin_pigmentation=coalesce(sqlc.narg('skin_pigmentation'), skin_pigmentation),
  hair_color=coalesce(sqlc.narg('hair_color'), hair_color),
  eye_color=coalesce(sqlc.narg('eye_color'), eye_color),
  size_preferences=coalesce(sqlc.narg('size_preferences'), size_preferences)
WHERE user_id = sqlc.arg('userId')::UUID;

-- name: DeleteProfile :exec
DELETE FROM profiles WHERE id=$1;

-- name: GetFollowerProfiles :many
SELECT p.*
FROM profiles p
JOIN profile_follows f ON p.id=f.following_id
WHERE f.following_id=$1
ORDER BY f.created_at DESC;

-- name: GetFollowingProfiles :many
SELECT p.*
FROM profiles p
JOIN profile_follows f ON p.id=f.follower_id
WHERE f.follower_id=$1
ORDER BY f.created_at DESC;

-- name: CreateProfileFollow :exec
INSERT INTO profile_follows(follower_id, following_id)
VALUES(@follower_id, @following_id);

-- name: DeleteProfileFollow :exec
DELETE FROM profile_follows
WHERE follower_id=@follower_id AND following_id=@following_id;
