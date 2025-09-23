CREATE TABLE IF NOT EXISTS users (id VARCHAR, email VARCHAR, points VARCHAR);
CREATE TABLE IF NOT EXISTS points (id VARCHAR, points INT, user_id VARCHAR);
INSERT INTO users (id, email, points)
VALUES ('ID-123', 'abc@mail.com', 1),
  ('ID-168', 'test@test.com', 2),
  ('ID-166', 'helloo@world.com', 3);
INSERT INTO points (id, points, user_id)
VALUES ('111', 12, 'ID-123'),
  ('222', 1, 'ID-168'),
  ('333', 99, 'ID-166')