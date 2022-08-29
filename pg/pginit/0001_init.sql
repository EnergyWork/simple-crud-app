create table film (
    id              bigserial not null,
    type            varchar(30),
    name            varchar(255) not null,
    year            int,
    duration        varchar(25),
    serial_count    int not null default 1, 
    score           int check(score <= 10),
    comment         text
);