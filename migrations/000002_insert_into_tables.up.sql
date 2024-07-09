INSERT INTO payments (id, reservation_id, amount, method, status)
VALUES
('a1234567-89ab-cdef-0123-456789abcdef', '550e8400-e29b-41d4-a716-446655440001', 100.00, 'cash', 'completed'),
('b2345678-9abc-def0-1234-56789abcdef0', '550e8400-e29b-41d4-a716-446655440002', 50.25, 'debit_card', 'completed'),
('c3456789-abcd-ef01-2345-6789abcdef01', '550e8400-e29b-41d4-a716-446655440003', 75.80, 'online_transfer', 'pending'),
('d456789a-bcde-f012-3456-789abcdef012', '550e8400-e29b-41d4-a716-446655440004', 125.50, 'cash', 'completed'),
('e56789ab-cdef-0123-4567-89abcdef0123', '550e8400-e29b-41d4-a716-446655440005', 30.00, 'debit_card', 'failed'),
('f6789abc-def0-1234-5678-9abcdef01234', '550e8400-e29b-41d4-a716-446655440006', 200.00, 'online_transfer', 'completed'),
('91815a5b-d382-45b1-8962-a21c2f78f9c4', '550e8400-e29b-41d4-a716-446655440007', 80.75, 'cash', 'refunded'),
('e42ef6f2-f9ff-4cc5-9178-ce4fadf1131f', '550e8400-e29b-41d4-a716-446655440008', 150.00, 'debit_card', 'completed'),
('81c602c8-04c6-441a-ae79-91f744f62a3c', '550e8400-e29b-41d4-a716-446655440009', 45.50, 'online_transfer', 'pending'),
('6542fa34-bf97-4d56-bf21-3bbc789a5c25', '550e8400-e29b-41d4-a716-446655440010', 300.00, 'cash', 'completed');