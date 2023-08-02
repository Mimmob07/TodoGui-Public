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
    ('name1', 'user1', 'user1.@icloud.com', '3c41f8b6fd29b4377585b5a411a6bd0db12427712bce04c41b10b8e8dc8b49e0'),
    ('name2', 'user2', 'user2.@fakeemail.com', '8d5303e5f7b71eaeaee5aae65563d3eea483c5bf6a4d9e9406875bc08061d969'),
    ('name3', 'user3', 'user3.@fakeemail.com', '99ae24526dee0f3b422595113cf2396adfc3c4f684319864a3e38a47d464aefb'),
    ('name4', 'user4', 'email4', 'e1bd35c8f15d49d6703a7dfadaf6ea174326fac5dd212dc3585518c79e676d3e');

-- ('name1', 'user1', 'user1.@icloud.com', 'user1placeholder'),
-- ('name2', 'user2', 'user2.@fakeemail.com', 'user2splaceholder'),
-- ('name3', 'user3', 'user3.@fakeemail.com', 'user3splaceholder'),
-- ('name4', 'user4', 'email4', 'user4splaceholder');


-- default testing data
-- paswords are hashed using sha256