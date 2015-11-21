;; $ sudo apt-get install emacs
;; $ curl -fsSL https://raw.githubusercontent.com/cask/cask/master/go | python
;; $ cask
;; $ make prepare
;; $ make all

(require 'cask "~/.cask/cask.el")
(cask-initialize)

(require 'esqlite)

(setq db-filepath "db.sqlite")

;; Stream API
(setq my-stream (esqlite-stream-open db-filepath))
(print (unwind-protect
    (esqlite-stream-read my-stream "SELECT * FROM users")
  (esqlite-stream-close my-stream)))

;; Sync Read API
(print (esqlite-read db-filepath "SELECT * FROM users"))
