CREATE TABLE IF NOT EXISTS cabins (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    max_capacity INT NOT NULL,
    regular_price INT NOT NULL,
    discount INT,
    description TEXT,
    image VARCHAR(255)
);
