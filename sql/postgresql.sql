create table users(
    user_id uuid not null,
    username varchar(50) not null,
    password varchar(100) not null,
    active boolean not null,
    primary key(user_id),
    unique(username)
);
