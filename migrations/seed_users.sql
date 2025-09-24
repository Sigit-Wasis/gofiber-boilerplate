INSERT INTO users (name, email) VALUES
('Admin','admin@example.com'),
('John Doe', 'john@example.com'),
('Jane Doe', 'jane@example.com');
ON CONFLICT (email) DO NOTHING;