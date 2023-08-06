-- +goose Up
CREATE TYPE product_gender AS ENUM ('male','female','unisex');
CREATE TYPE product_category AS ENUM ('headwear','eyewear','tops','bottoms','dresses','outerwear','accessories','shoes','swimwear','piercings','necklaces','rings');

CREATE TABLE IF NOT EXISTS products (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(127) NOT NULL,
  description TEXT,
  sourceUrl VARCHAR(255) DEFAULT NULL,
  available_to_buy BOOLEAN NOT NULL,
  gender product_gender DEFAULT NULL,
  categories product_category DEFAULT NULL,
  materials JSON NOT NULL DEFAULT json_build_array(),
  image_url VARCHAR(255) NOT NULL,
  profile_id UUID NOT NULL,
  CONSTRAINT products_ibfk_1 FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

-- +goose Down
DROP TYPE product_gender;
DROP TYPE product_category;
DROP TABLE products;
