-- name: createProduct :exec
INSERT INTO products (id, name, description, sourceUrl, available_to_buy, gender, categories, materials, image_url, profile_id)
VALUES (UUID_TO_BIN(UUID()), ?, ?, ?, ?, ?, ?, ?, ?, UUID_TO_BIN(sqlc.arg('profileUUID')));

-- name: getProductById :one
SELECT BIN_TO_UUID(id) AS uuid, * FROM products
WHERE id = UUID_TO_BIN(sqlc.arg('uuid'))
LIMIT 1;

-- name: updateProductById :exec
UPDATE products
SET name=COALESCE(sqlc.narg('name'), name),
   description=COALESCE(?,description),
   sourceUrl=COALESCE(?, sourceUrl),
   available_to_buy=COALESCE(sqlc.narg('availableToBuy'), available_to_buy),
   gender=COALESCE(?, gender),
   categories=COALESCE(?, categories),
   materials=COALESCE(sqlc.narg('materials'), materials),
   image_url=COALESCE(sqlc.narg('imageUrl'), image_url)
WHERE id = UUID_TO_BIN(sqlc.arg('uuid'));
