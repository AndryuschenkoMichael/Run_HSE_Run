CREATE TABLE users
(
    id       serial PRIMARY KEY,
    nickname varchar(255) not null unique,
    email    varchar(255) not null unique,
    image    int          not null
);

CREATE TABLE friends
(
    id       serial PRIMARY KEY,
    user_id1 int references users (id) on delete cascade not null,
    user_id2 int references users (id) on delete cascade not null
);

CREATE TABLE rooms
(
    id        serial PRIMARY KEY,
    code      varchar(255) not null,
    campus_id int          not null
);

CREATE TABLE edges
(
    id            serial PRIMARY KEY,
    room_start_id int references rooms (id) on delete cascade not null,
    room_end_id   int references rooms (id) on delete cascade not null,
    campus_id     int                                         not null
);