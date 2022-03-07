CREATE TABLE reception
(
    "id"             SERIAL PRIMARY KEY,
    "reason"         TEXT,
    "observation"    TEXT,
    "status"         VARCHAR(20),
    "date_admission" date,
    "date_delivery"  date,
    "created_at"     timestamp,
    "updated_at"     timestamp,
    "customer"       integer,
    "equipment"      integer
);

