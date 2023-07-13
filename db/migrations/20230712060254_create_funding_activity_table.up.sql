create table funding_activities
(
    id               serial      not null,
    date_note        varchar(50),
    description      varchar(50) not null,
    amount           money         not null,
    year_period_id   int,
    type_transaction TRANSACTION_CATEGORY,
    created_at       timestamp default current_timestamp,
    primary key (id),
    constraint fk_year_period_funding foreign key (year_period_id) references year_periods (id)

);

create index desc_funding_search_index on funding_activities using gin (to_tsvector('indonesian', description));
create index date_note_funding_search_index on funding_activities using gin (to_tsvector('indonesian', date_note));