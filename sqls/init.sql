CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE service_versions (
    id SERIAL PRIMARY KEY,
    service_id INTEGER REFERENCES services(id),
    version VARCHAR(255) NOT NULL
);


-- Inserting 100 services
DO $$
BEGIN
    FOR i IN 1..100 LOOP
        INSERT INTO services (name, description)
        VALUES ('Service ' || i, 'Description for Service ' || i);
    END LOOP;
END $$;

-- Inserting versions for each service
DO $$
BEGIN
    FOR i IN 1..100 LOOP
        INSERT INTO service_versions (service_id, version)
        VALUES (i, 'v1.0');
        INSERT INTO service_versions (service_id, version)
        VALUES (i, 'v1.1');
    END LOOP;
END $$;