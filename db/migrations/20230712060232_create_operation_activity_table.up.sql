create table operation_activities
(
    id          serial      not null,
    date_note   varchar(50) not null,
    description varchar(50) not null,
    amount      money         not null,
    type_transaction TRANSACTION_CATEGORY,
    year_period_id  int,
    created_at  timestamp default current_timestamp,
    primary key (id),
    constraint fk_year_period_operation foreign key (year_period_id) references year_periods (id)

);

create index desc_operation_search_index on operation_activities using gin (to_tsvector('indonesian', description));
create index type_transaction_operation_search_index on operation_activities using gin (to_tsvector('indonesian', type_transaction));
create index date_note_operation_search_index on operation_activities using gin (to_tsvector('indonesian', date_note));