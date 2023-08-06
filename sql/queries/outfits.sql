-- name: GetPiecesByOutfitId :many
SELECT p.*, v.*
FROM outfit_pieces p
JOIN product_variants v ON p.variant_id=v.id
WHERE p.outfit_id=$1
ORDER BY p.created_at DESC;
