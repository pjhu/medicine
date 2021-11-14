CREATE TABLE member (
  id BIGINT NOT NULL,
  phone VARCHAR(32) NOT NULL,
  nickname VARCHAR(255),
  password VARCHAR(255),
  created_at TIMESTAMP,
  last_modified_at TIMESTAMP,
  PRIMARY KEY (id)
);