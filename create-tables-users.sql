DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name TEXT NOT NULL,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user
    (name, username, email, password)
VALUES
    ('name1', 'user1', 'user1.@icloud.com', 'user1placeholder'),
    ('name2', 'user2', 'user2.@fakeemail.com', 'user2splaceholder'),
    ('name3', 'user3', 'user3.@fakeemail.com', 'user3spaceholder'),
    ('name4', 'user4', 'email4', 'user4splaceholder');

-- default testing data