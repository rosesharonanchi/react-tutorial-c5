CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  total_savings VARCHAR(255) DEFAULT '0',
  total_withdrawals VARCHAR(255) DEFAULT '0'
);

CREATE TABLE transactions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  amount VARCHAR(255) NOT NULL,
  reason TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  type VARCHAR(25)
);

