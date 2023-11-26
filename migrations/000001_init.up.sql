CREATE TABLE quests(
    id BIGSERIAL NOT NULL Primary Key,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500) NOT NULL,
    author_comment VARCHAR(500),
    point VARCHAR(500),
    age_level int not null,
    difficult VARCHAR(100) NOT NULL,
    duration int not null,
    location VARCHAR(100),
    organizer VARCHAR(100) NOT NULL
);

CREATE TABLE users(
    id BIGSERIAL NOT NULL Primary Key,
    tg_user_id VARCHAR(300) NOT NULL,
    name VARCHAR(300) NOT NULL,
    age int not null,
    phone VARCHAR(12) NOT NULL
);