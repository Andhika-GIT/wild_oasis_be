CREATE TABLE IF NOT EXISTS settings (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    minBookingLength INT NOT NULL,
    maxBookingLength INT NOT NULL,
    maxGuestsPerCabin INT NOT NULL,
    breakfastPrice FLOAT4
);
