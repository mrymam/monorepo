
create table IF NOT EXISTS slack_user_identities (
  slack_user_id VARCHAR(255) UNIQUE,
  slack_team_id VARCHAR(255),
  user_id VARCHAR(255)
);