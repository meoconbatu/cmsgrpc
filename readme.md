# CMS
A simple web app, using gRPC

## Requirements 

- [Go](https://golang.org)
- [Postgres](https://www.postgresql.org)

## Usage 

Start postgres database

```bash
pg_ctl -D /usr/local/var/postgres start
```
Connect to the default database - postgres
```bash
psql postgres
```
Then run sql file to create database and tables
```bash
\i domain/001_init.up.sql
```
Run domain server
```bash
go run domain/cmd/main.go
```
Run view server
```bash
go run view/cmd/main.go
```
Now you can use your browser to access the website:
- [create](http://localhost:3000/new) a new page/post (post is not implemented yet)
- [list](http://localhost:3000/page/) all the pages you created