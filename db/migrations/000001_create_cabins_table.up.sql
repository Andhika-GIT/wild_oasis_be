CREATE TABLE IF NOT EXISTS cabins (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    maxCapacity INT NOT NULL,
    regularPrice INT NOT NULL,
    discount INT,
    description TEXT,
    image VARCHAR(255)
);
