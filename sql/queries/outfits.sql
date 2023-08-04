-- name: GetPiecesByOutfitId :many
SELECT BIN_TO_UUID(p.id) AS piece_uuid, p.*, BIN_TO_UUID(v.id) AS variant_uuid, v.*
FROM outfit_pieces p
JOIN product_variants v ON p.variant_id=v.id
WHERE p.outfit_id=UUID_TO_BIN(sqlc.arg("OutfitUUID"))
ORDER BY p.created_at DESC;
