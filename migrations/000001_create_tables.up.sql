CREATE TYPE payment_method AS enum ('cash', 'debit_card', 'online_transfer');
CREATE TYPE payment_status AS enum ('pending', 'failed', 'completed', 'refunded');

CREATE TABLE payments (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid NOT NULL,
    amount DECIMAL(10, 2),
    method payment_method NOT NULL,
    status payment_status NOT NULL
);