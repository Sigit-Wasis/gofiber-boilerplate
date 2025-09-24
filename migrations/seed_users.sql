INSERT INTO users (name, email) VALUES
('Admin','admin@example.com')
ON CONFLICT (email) DO NOTHING;