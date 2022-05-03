CREATE TABLE users
(
    id       serial PRIMARY KEY,
    nickname varchar(255) not null unique,
    email    varchar(255) not null unique
);

CREATE TABLE friends
(
    id       serial PRIMARY KEY,
    user_id1 int references users (id) on delete cascade not null,
    user_id2 int references users (id) on delete cascade not null
);