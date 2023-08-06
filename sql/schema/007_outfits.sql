-- +goose Up
CREATE TABLE IF NOT EXISTS outfits (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  description TEXT,
  image_urls JSON NOT NULL DEFAULT json_build_array(),
  profile_id binary(16) NOT NULL,
  CONSTRAINT outfits_ibfk_1 FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

-- +goose Down
DROP TABLE outfits;
