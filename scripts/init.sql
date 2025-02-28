CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    price BIGINT NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE stats (
    user_id BIGINT NOT NULL,
    service_id BIGINT NOT NULL,
    count BIGINT NOT NULL,
    PRIMARY KEY (user_id, service_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
);

INSERT INTO services (name, description, price) VALUES
    ('Service A', 'Description for Service A', 100),
    ('Service B', 'Description for Service B', 200);

INSERT INTO users (name) VALUES
    ('User 1'),
    ('User 2');

INSERT INTO stats (user_id, service_id, count) VALUES
    (1, 1, 5),
    (1, 2, 3),
    (2, 1, 2);