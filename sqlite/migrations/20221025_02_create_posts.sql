
create table posts (
  id VARCHAR(255) PRIMARY KEY,
  title VARCHAR(255),
  user_id VARCHAR(255),
  foreign key (user_id) references users(id)
);