DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id INTEGER PRIMARY KEY,
	name TEXT
);

INSERT INTO users (name) values ("foo"), ("bar"), ("spam");
