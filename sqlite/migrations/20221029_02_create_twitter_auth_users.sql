
create table IF NOT EXISTS twitter_auth_users (
  twitter_user_id VARCHAR(255) UNIQUE,
  user_id VARCHAR(255)
);