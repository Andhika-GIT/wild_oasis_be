CREATE TABLE IF NOT EXISTS bookings (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    num_nights INT,
    num_guests INT,
    cabin_price FLOAT4,
    extras_price FLOAT4,
    total_price FLOAT4,
    status VARCHAR(255),
    has_breakfast BOOL,
    is_paid BOOL,
    observations VARCHAR(255),
    cabin_id INT,
    guest_id INT,
    CONSTRAINT fk_cabin
        FOREIGN KEY (cabin_id) REFERENCES cabins(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_guest
        FOREIGN KEY (guest_id) REFERENCES guests(id)
        ON DELETE SET NULL
);
