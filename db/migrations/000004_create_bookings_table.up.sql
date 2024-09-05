CREATE TABLE IF NOT EXISTS bookings (
    id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    startDate TIMESTAMP,
    endDate TIMESTAMP,
    numNights INT,
    numGuests INT,
    cabinPrice FLOAT4,
    extrasPrice FLOAT4,
    totalPrice FLOAT4,
    status VARCHAR(255),
    hasBreakfast BOOL,
    isPaid BOOL,
    observations VARCHAR(255),
    CONSTRAINT fk_cabin
        FOREIGN KEY (cabinId) REFERENCES cabins(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_guest
        FOREIGN KEY (guestId) REFERENCES guests(id)
        ON DELETE SET NULL
);
