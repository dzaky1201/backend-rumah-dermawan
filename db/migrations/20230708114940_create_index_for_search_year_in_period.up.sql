create index year_search_index on year_periods using gin (to_tsvector('indonesian', info_period ->> 'Month'));