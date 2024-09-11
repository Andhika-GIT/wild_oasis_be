CREATE TABLE IF NOT EXISTS settings (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    min_booking_length INT NOT NULL,
    max_booking_length INT NOT NULL,
    max_guests_per_cabin INT NOT NULL,
    breakfast_price FLOAT4
);
