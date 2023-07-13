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
    ('name1', 'user1', 'user1.@icloud.com', '3c9915fc4adb8eba9035954e536d5b3433f6945f2342073737c72754486f8593'),
    ('name2', 'user2', 'user2.@fakeemail.com', 'd1e1775f6202e384364c0f2236cd141ef2374549bfd7212896764bc125863ead'),
    ('name3', 'user3', 'user3.@fakeemail.com', 'b2394c84491ee2b6cb3b95070bcc8b0137e20705da6ef78cfcda6f0517641383'),
    ('name4', 'user4', 'email4', 'f4aa0655cdb8d4fcf6f719c7a786de10556783c70bfb8ef1d78923482fe6ebbc');

-- ('name1', 'user1', 'user1.@icloud.com', 'user1placeholder'),
-- ('name2', 'user2', 'user2.@fakeemail.com', 'user2splaceholder'),
-- ('name3', 'user3', 'user3.@fakeemail.com', 'user3splaceholder'),
-- ('name4', 'user4', 'email4', 'user4splaceholder');


-- default testing data