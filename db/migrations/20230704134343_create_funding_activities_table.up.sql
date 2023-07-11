create table funding_activities
(
    id          serial      not null,
    date_note   date      default current_date,
    description varchar(50) not null,
    amount      int         not null,
    year_period_id  int,
    type_transaction TRANSACTION_CATEGORY,
    created_at  timestamp default current_timestamp,
    primary key (id),
    constraint fk_year_period_funding foreign key (year_period_id) references year_periods (id)

)