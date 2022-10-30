
create table IF NOT EXISTS twitter_user_identities (
  twitter_user_id VARCHAR(255) UNIQUE,
  user_id VARCHAR(255)
);