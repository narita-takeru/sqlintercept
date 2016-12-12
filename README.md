
mysqlとの通信を傍受して、 テーブル毎の`SELECT`の回数をカウントするツールです。

![説明画像] (https://github.com/narita-takeru/sqlintercept/blob/master/sqlintercept-readme.gif)

```
go run main.go -src=127.0.0.1:3305 -dst=127.0.0.1:3306
```

↑のように起動すると、ポート3305で待ち受けて待機状態になります。

```
mysql -u root hoge_database -h 127.0.0.1 -P 3305 --ssl-mode=DISABLED -e "select count(*) from hoge_table"
```

みたいに、3305に対してクエリを送信すると、`hoge_table`に打たれた`SELECT`の回数をカウントします。

