CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT,
  hashed_password TEXT,
  bio TEXT,
  auth_token TEXT NOT NULL,
  created_at int NOT NULL,
  updated_at int
);

CREATE TABLE user_sessions (
  id SERIAL PRIMARY KEY,
  session_token TEXT NOT NULL,
  login_time int NOT NULL,
  logout_time int,
  ip_address TEXT,
  user_agent TEXT,
  is_active BOOLEAN NOT NULL DEFAULT true,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
