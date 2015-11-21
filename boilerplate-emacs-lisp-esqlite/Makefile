all:
	emacs --script main.el

prepare: clean
	sqlite3 db.sqlite < users.sql

clean:
	rm -rf db.sqlite
