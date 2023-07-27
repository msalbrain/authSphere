PRAGMA foreign_keys = ON;


CREATE TABLE users (
  id int PRIMARY KEY,
  name text NOT NULL,
  email text,
  hashed_password text,
  bio text,
  auth_token text NOT NULL,
  created_at int NOT NULL,
  updated_at int
);

CREATE TABLE user_sessions (
  id int PRIMARY KEY,
  session_token text NOT NULL,
  login_time int NOT NULL,
  logout_time int,
  ip_address text,
  user_agent text,
  is_active BOOLEAN NOT NULL DEFAULT true,
  user_id int NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);
