create table users
(
    id         serial,
    name       varchar(20)  not null,
    email      varchar(50)  not null,
    password   varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    primary key (id),
    constraint email_unique unique (email)
);