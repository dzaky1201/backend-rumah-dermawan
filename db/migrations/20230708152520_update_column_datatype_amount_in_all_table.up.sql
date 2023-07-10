ALTER TABLE funding_activities
    ALTER COLUMN amount TYPE money;

ALTER TABLE operation_activities
    ALTER COLUMN amount TYPE money;

ALTER TABLE invests_activities
    ALTER COLUMN amount TYPE money;