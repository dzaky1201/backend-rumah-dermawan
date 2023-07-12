create table year_periods
(
    id          serial not null,
    info_period jsonb,
    created_at  timestamp default current_timestamp,
    primary key (id)

);

create index year_search_index on year_periods using gin (to_tsvector('indonesian', info_period ->> 'Month'));
create index year_month_search_index on year_periods using gin (to_tsvector('indonesian', info_period ->> 'Year'));

create type TRANSACTION_CATEGORY as enum ('credit', 'debit');