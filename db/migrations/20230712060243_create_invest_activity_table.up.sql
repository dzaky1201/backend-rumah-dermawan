create table invests_activities
(
    id               serial      not null,
    date_note        text        not null,
    description      varchar(50) not null,
    amount           int       not null,
    type_transaction TRANSACTION_CATEGORY,
    year_period_id   int,
    created_at       timestamp default current_timestamp,
    updated_at       timestamp default current_timestamp,
    primary key (id),
    constraint fk_year_period_invests foreign key (year_period_id) references year_periods (id)

);

create index desc_invest_search_index on invests_activities using gin (to_tsvector('indonesian', description));
create index date_note_invest_search_index on invests_activities using gin (to_tsvector('indonesian', date_note));