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

Run domain server
```bash
cd domain/cmd
go build -o domain
./cmd/domain
```

Run user server
```bash
cd user/cmd
go build -o user
./cmd/user
```

Run view server
```bash
cd view/cmd
go build -o view
./cmd/view
```
Now you can use your browser to access the website:
- [create](http://localhost:3000/new) a new page/post (post is not implemented yet)
- [list](http://localhost:3000/page/) all the pages you created
- [login](http://localhost:3000/login)
- [register](http://localhost:3000/register) new user account