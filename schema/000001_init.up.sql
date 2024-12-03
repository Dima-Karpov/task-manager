CREATE TABLE users (
  id serial NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE posts_lists (
    id serial NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(500) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users_lists (
  id serial NOT NULL UNIQUE,
  user_id INT REFERENCES users (id) ON DELETE cascade NOT NULL,
  post_id INT REFERENCES posts_lists (id) ON DELETE cascade NOT NULL
);