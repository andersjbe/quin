-- +goose Up
CREATE TABLE IF NOT EXISTS product_saves (
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  product_id UUID NOT NULL,
  profile_id UUID NOT NULL,
  PRIMARY KEY (product_id,profile_id),
  CONSTRAINT product_saves_ibfk_1 FOREIGN KEY (product_id) REFERENCES products (id),
  CONSTRAINT product_saves_ibfk_2 FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

-- +goose Down
DROP TABLE product_saves;
