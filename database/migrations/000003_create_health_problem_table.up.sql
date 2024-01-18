CREATE TABLE IF NOT EXISTS tb_health_problems (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY, client_id uuid REFERENCES tb_clients (id), name VARCHAR(255) NOT NULL, grau INT CHECK (grau IN (1, 2)), created_at TIMESTAMPTZ DEFAULT current_timestamp, updated_at TIMESTAMPTZ DEFAULT current_timestamp
);