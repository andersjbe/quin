-- +goose Up
CREATE TABLE IF NOT EXISTS product_variants (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(127) NOT NULL,
  colors JSON NOT NULL DEFAULT json_build_array(),
  pattern VARCHAR(127) NOT NULL,
  image_urls JSON DEFAULT json_build_array(),
  product_id UUID NOT NULL,
  CONSTRAINT product_variants_ibfk_1 FOREIGN KEY (product_id) REFERENCES products (id)
);

-- +goose Down
DROP TABLE product_variants;
