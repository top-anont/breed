CREATE TABLE IF NOT EXISTS breed (
    id VARCHAR(36) PRIMARY KEY,
    name_en VARCHAR NOT NULL,
    name_th VARCHAR NOT NULL,
    short_name VARCHAR(30) NOT NULL,
    remark VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by_id VARCHAR(36) NOT NULL,
    created_by VARCHAR NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by_id VARCHAR(36) NOT NULL,
    updated_by VARCHAR NOT NULL
);

CREATE INDEX idx_breed_id_names_short_name ON breed USING btree (id, name_en, name_th, short_name);