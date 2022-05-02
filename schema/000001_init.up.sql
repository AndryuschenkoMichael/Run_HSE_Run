CREATE TABLE users
(
    id                 serial PRIMARY KEY,
    user_name          varchar(255) not null unique,
    user_mail          varchar(255) not null unique,
    authorization_code integer      not null
);

CREATE TABLE friends
(
    id       serial PRIMARY KEY,
    user_id1 int references users (id) on delete cascade not null,
    user_id2 int references users (id) on delete cascade not null
);