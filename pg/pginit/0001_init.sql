-- User authorization entity
create table user_auth (
    id  bigserial primary key,
    user_name varchar(100) not null unique,
    user_password varchar not null,
    private_key varchar not null,
    expired_at timestamp not null -- if expired need login again
);

-- Films entity
create table film (
    id              bigserial primary key,
    user_id         bigint not null references user_auth (id),
    type            varchar(30),
    name            varchar(255) not null,
    release_date    timestamp,
    duration        varchar(10),
    score           integer check(score >= 0 and score <= 100),
    comment         text
);

-- Serials entity
create table serial (
    id              bigserial primary key,
    user_id         bigint not null references user_auth (id),
    name            varchar(255) not null,
    release_date    timestamp,
    duration        varchar(10),
    score           integer check(score >= 0 and score <= 100),
    comment         text
);

-- Seasons entity
create table season (
    id  bigserial primary key,
    serial_id bigint references serial (id),
    number integer not null,
    series json -- e.g. {"1": "42m", "2": "46m"}
);