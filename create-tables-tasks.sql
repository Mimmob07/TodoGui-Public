DROP TABLE IF EXISTS task;
CREATE TABLE task (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  task TEXT NOT NULL,
  status TEXT NOT NULL,
  user INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO task
    (task, status, user)
VALUES
    ('Buy groceries', 'incomplete', 1),
    ('Complete programming project', 'incomplete', 1),
    ('Book a doctors appointment', 'incomplete', 2);
-- default testing data