CREATE TYPE otp_auth_status AS ENUM ('created', 'expired', 'validated');

CREATE TABLE otp_auth (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  otp INT NOT NULL,
  otp_expired_at TIMESTAMPTZ NOT NULL,
  status otp_auth_status DEFAULT 'created',
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);