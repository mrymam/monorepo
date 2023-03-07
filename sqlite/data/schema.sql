create table exp_users(
    id varchar(255) NOT NULL, 
    username varchar(255) NOT NULL,
    listenings TEXT NOT NULL
);

create table friends (
  exp_user_id varchar(255) NOT NULL,
  friend_id varchar(255) NOT NULL
);