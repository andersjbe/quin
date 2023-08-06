-- +goose Up
CREATE TYPE profile_gender AS ENUM ('man','woman','nonbinary');

CREATE TABLE IF NOT EXISTS profiles (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  username VARCHAR(127) UNIQUE NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  gender profile_gender DEFAULT NULL,
  height_inches INT DEFAULT NULL,
  weight_lbs INT DEFAULT NULL,
  skin_pigmentation INT DEFAULT NULL,
  hair_color VARCHAR(63) DEFAULT NULL,
  eye_color VARCHAR(63) DEFAULT NULL,
  size_preferences JSON DEFAULT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT profiles_ibfk_1 FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
DROP TYPE gender;
DROP TABLE profiles;
