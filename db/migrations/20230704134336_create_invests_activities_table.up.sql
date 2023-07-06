create table invests_activities
(
    id               serial      not null,
    date_note        date      default current_date,
    description      varchar(50) not null,
    amount           int         not null,
    type_transaction TRANSACTION_CATEGORY,
    periode_id       int,
    created_at       timestamp default current_timestamp,
    primary key (id),
    constraint fk_year_period_invests foreign key (periode_id) references year_periods (id)

)