# Write your MySQL query statement below
SELECT Tweets.tweet_id FROM Tweets WHERE LENGTH(Tweets.content) > 15;