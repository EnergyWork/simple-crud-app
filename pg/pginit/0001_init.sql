set time zone 'Europe/Moscow';

-- Sessons entity
create table sessions (
    id bigserial primary key,
    token varchar unique,
    created timestamp not null default now(),
    deadline timestamp
);
-- User authorization entity
create table users (
    id  bigserial primary key,
    session_id bigint not null references sessions (id),
    login varchar(100) not null unique,
    password varchar not null,
    access_key varchar not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- Films entity
create table films (
    id              bigserial primary key,
    user_id         bigint not null references users (id),
    name            varchar(255) not null,
    release_date    timestamp,
    duration        varchar(10),
    score           integer check(score >= 0 and score <= 100),
    comment         text,
    created_at      timestamp not null default now(),
    updated_at      timestamp
);

-- Serials entity
create table serials (
    id              bigserial primary key,
    user_id         bigint not null references users (id),
    name            varchar(255) not null,
    release_date    timestamp,
    score           integer check(score >= 0 and score <= 100),
    comment         text,
    created_at      timestamp not null default now(),
    updated_at      timestamp
);

-- Seasons entity
create table seasons (
    id              bigserial primary key,
    serial_id       bigint references serials (id),
    number          integer not null, -- 
    series          json, -- e.g. {"1": "42m", "2": "46m"}
    created_at      timestamp not null default now(),
    updated_at      timestamp
);