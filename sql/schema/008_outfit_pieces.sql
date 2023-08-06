-- +goose Up
CREATE TABLE IF NOT EXISTS outfit_pieces (
  created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  sizes JSON NOT NULL DEFAULT json_build_object(),
  outfit_id UUID NOT NULL,
  variant_id UUID NOT NULL,
  PRIMARY KEY (outfit_id,variant_id),
  KEY variant_id (variant_id),
  CONSTRAINT outfit_pieces_ibfk_1 FOREIGN KEY (outfit_id) REFERENCES outfits (id),
  CONSTRAINT outfit_pieces_ibfk_2 FOREIGN KEY (variant_id) REFERENCES product_variants (id)
);

-- +goose Down
DROP TABLE outfit_pieces;
