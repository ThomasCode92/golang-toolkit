--- Create Tables
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(60),
    user_active INTEGER DEFAULT 0,
    is_admin INTEGER DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE plans (
    id SERIAL PRIMARY KEY,
    plan_name VARCHAR(255),
    plan_amount INTEGER,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE user_plans (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
    plan_id INTEGER REFERENCES plans (id) ON DELETE CASCADE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

--- Seed Data
INSERT INTO users (email, first_name, last_name, password, user_active, is_admin, created_at, updated_at)
VALUES ('admin@example.com', 'Admin', 'User', '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe', 1, 1, '2022-03-14 00:00:00', '2022-03-14 00:00:00');

INSERT INTO plans (plan_name, plan_amount, created_at, updated_at)
VALUES
    ('Bronze Plan', 1000, '2022-05-12 00:00:00', '2022-05-12 00:00:00'),
    ('Silver Plan', 2000, '2022-05-12 00:00:00', '2022-05-12 00:00:00'),
    ('Gold Plan', 3000, '2022-05-12 00:00:00', '2022-05-12 00:00:00');
