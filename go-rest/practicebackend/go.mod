module example.com/practicebackend

go 1.19

replace example.com/backend => ../backend

require example.com/backend v0.0.0-00010101000000-000000000000

require github.com/mattn/go-sqlite3 v1.14.16 // indirect