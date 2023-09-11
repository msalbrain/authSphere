PRAGMA foreign_keys = ON;


CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  email text,
  hashed_password text,
  bio text,
  refresh_token text NOT NULL,
  auth_token text NOT NULL,
  created_at int NOT NULL,
  updated_at int
);


CREATE TABLE IF NOT EXISTS user_sessions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  session_token text NOT NULL,
  expires int NOT NULL,
  login_time int NOT NULL,
  ip_address text,
  user_agent text,
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE IF NOT EXISTS app_state (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  jwtoption text CHECK (jwtoption IN ('symmetric', 'asymmtric'))
)



CREATE TABLE IF NOT EXISTS mail (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  subject text,
  body text,
  updated int
)

