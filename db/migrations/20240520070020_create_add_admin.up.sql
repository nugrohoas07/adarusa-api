ALTER TABLE users ALTER COLUMN password TYPE VARCHAR(255);

INSERT INTO users(email,password,role_id,status) VALUES ('admin','$2a$10$k9ABmFXo3CmISrvsVvA4eOn3I8OhXGjlFWKJbK8UzMQqxS9/NDknO',1,'verified')