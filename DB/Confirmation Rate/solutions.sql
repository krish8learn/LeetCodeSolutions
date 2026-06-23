-- Solution: Confirmation Rate
-- 
-- Key Concepts:
-- 1. LEFT JOIN Signups with Confirmations to include users with no confirmations
-- 2. Use CASE statement to count only 'confirmed' actions
-- 3. Handle division by zero for users with no confirmation requests
-- 4. Round result to 2 decimal places

SELECT 
    s.user_id,
    ROUND(
        COALESCE(
            SUM(CASE WHEN c.action = 'confirmed' THEN 1 ELSE 0 END) / 
            COUNT(c.user_id),
            0
        ),
        2
    ) AS confirmation_rate
FROM Signups s
LEFT JOIN Confirmations c ON s.user_id = c.user_id
GROUP BY s.user_id
ORDER BY s.user_id;

-- ============================================================
-- EXPLANATION OF THE CALCULATION:
-- ============================================================
--
-- Example with User 7 (from the problem):
-- - User 7 appears in 3 rows after the LEFT JOIN:
--   Row 1: user_id=7, action='confirmed' -> CASE returns 1
--   Row 2: user_id=7, action='confirmed' -> CASE returns 1
--   Row 3: user_id=7, action='confirmed' -> CASE returns 1
-- - SUM(CASE...) = 1 + 1 + 1 = 3 (total confirmed)
-- - COUNT(c.user_id) = 3 (total confirmation requests)
-- - Division: 3 / 3 = 1.0
-- - ROUND(1.0, 2) = 1.00
--
-- Example with User 6 (no confirmations):
-- - User 6 appears in 1 row after the LEFT JOIN:
--   Row 1: user_id=6, c.user_id=NULL (no matching confirmation)
-- - SUM(CASE...) = 0 (no confirmed actions)
-- - COUNT(c.user_id) = 0 (no confirmation requests, COUNT ignores NULLs)
-- - Division: 0 / 0 = NULL (division by zero)
-- - COALESCE(NULL, 0) = 0 (convert NULL to 0)
-- - ROUND(0, 2) = 0.00
--
-- Example with User 2 (mixed confirmations):
-- - User 2 appears in 2 rows:
--   Row 1: user_id=2, action='confirmed' -> CASE returns 1
--   Row 2: user_id=2, action='timeout' -> CASE returns 0
-- - SUM(CASE...) = 1 + 0 = 1 (one confirmed)
-- - COUNT(c.user_id) = 2 (two total requests)
-- - Division: 1 / 2 = 0.5
-- - ROUND(0.5, 2) = 0.50
--
-- ============================================================
-- KEY COMPONENTS:
-- ============================================================
--
-- LEFT JOIN:
--   Ensures all users from Signups are included, even if they
--   have no entries in Confirmations (c.user_id will be NULL)
--
-- CASE WHEN c.action = 'confirmed' THEN 1 ELSE 0 END:
--   Converts 'confirmed' actions to 1 and others (like 'timeout')
--   to 0, so SUM can count only the confirmed messages
--
-- COUNT(c.user_id):
--   Counts non-NULL confirmation records for each user
--   If no confirmations exist, COUNT returns 0
--
-- COALESCE(..., 0):
--   Handles the division by zero case when a user has no
--   confirmation requests (returns 0 instead of NULL)
--
-- ROUND(..., 2):
--   Rounds the confirmation rate to 2 decimal places
--
-- GROUP BY s.user_id:
--   Groups by user to aggregate their confirmation statistics

