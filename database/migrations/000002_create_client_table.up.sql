CREATE TABLE IF NOT EXISTS tb_clients (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, name VARCHAR(255) NOT NULL, birth_date DATE NOT NULL, sexo CHAR(1) NOT NULL CHECK (sexo IN ('f', 'm', 'u')), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
);