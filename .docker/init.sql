CREATE TABLE request
(
    id           SERIAL PRIMARY KEY,
    currency     VARCHAR     NOT NULL,
    average      DOUBLE PRECISION,
    start_date   DATE        NOT NULL,
    end_date     DATE        NOT NULL,
    request_date TIMESTAMPTZ NOT NULL
);

SET timezone = 'Europe/Warsaw';