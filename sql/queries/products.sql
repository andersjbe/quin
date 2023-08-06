-- name: createProduct :exec
INSERT INTO products (name, description, sourceUrl, available_to_buy, gender, categories, materials, image_url, profile_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: getProductById :one
SELECT * FROM products
WHERE id = UUID_TO_BIN(sqlc.arg('productId'))
LIMIT 1;

-- name: updateProductById :exec
UPDATE products
SET name=COALESCE(sqlc.narg('name'), name),
   description=COALESCE($1,description),
   sourceUrl=COALESCE($2, sourceUrl),
   available_to_buy=COALESCE(sqlc.narg('availableToBuy'), available_to_buy),
   gender=COALESCE($3, gender),
   categories=COALESCE($4, categories),
   materials=COALESCE(sqlc.narg('materials'), materials),
   image_url=COALESCE(sqlc.narg('imageUrl'), image_url)
WHERE id = sqlc.arg('id')::UUID;
