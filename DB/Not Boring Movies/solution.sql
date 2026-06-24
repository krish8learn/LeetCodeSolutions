-- SELECT *
-- FROM Movies
-- WHERE id NOT IN (SELECT id FROM Movies WHERE description = 'boring')
SELECT
    *
FROM
    Cinema
WHERE
    description != 'boring'
    AND MOD(id, 2) != 0
ORDER BY
    rating DESC;