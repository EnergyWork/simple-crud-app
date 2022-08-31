-- Authorization entity
create table auth (
    id  bigserial primary key,
    user_name varchar(100) not null,
    user_password varchar(255) not null,
    api_key varchar unique,
    expired_at timestamp
)

-- Films entity
create table film (
    id              bigserial primary key,
    type            varchar(30),
    name            varchar(255) not null,
    release_date    timestamp,
    duration        varchar(10),
    score           integer check(score >= 0 and score <= 100),
    comment         text
)

-- Serial entity
create table serial (
    id              bigserial primary key,
    name            varchar(255) not null,
    release_date    timestamp,
    duration        varchar(10),
    score           integer check(score >= 0 and score <= 100),
    comment         text
)

-- Season entity
create table season (
    id  bigserial primary key,
    serial_id bigint references serial (id),
    number integer not null,
    series json -- e.g. {"1": "42m", "2": "46m"}
);