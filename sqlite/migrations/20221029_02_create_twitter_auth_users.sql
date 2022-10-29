
create table IF NOT EXISTS twitter_auths (
  twitter_user_id VARCHAR(255) UNIQUE,
  user_id VARCHAR(255)
);