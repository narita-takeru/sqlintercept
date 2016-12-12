
<h1 align="center">
    <img alt="sqlintercept" src="https://github.com/narita-takeru/sqlintercept/blob/master/sqlintercept-readme.gif" />
</h1>

# sqlintercept

## About

sqlintercept is designed to help you increase your application's performance by reducing the number of queries it makes. It will watch your queries and visualize tables are too much accessed.

## Requirements

- go 
- mysql

#Installation

```bash
$ go get github.com/narita-takeru/sqlintercept/cmd/sqlintercept
```

#Usage

```bash
$ sqlintercept -src=127.0.0.1:3305 -dst=127.0.0.1:3306
```

##Try to run query
```bash
$ mysql -u root hoge_database -h 127.0.0.1 -P 3305 --ssl-mode=DISABLED -e "select count(*) from one table"
```

