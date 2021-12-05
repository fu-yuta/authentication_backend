# authentication_backend

### サーバーの起動
```
bee run -downdoc=true -gendoc=true
```

### swaggerのurl
```
http://127.0.0.1:8080/swagger/
```

### リクエスト
```
curl -X GET "http://127.0.0.1:8080/v1/hello/" -H  "accept: application/json" -H "Authorization: XXXX" --verbose
```

### データベース(mysql)の操作
```
# ログイン
mysql --user root --password

# データベースの作成
mysql> CREATE DATABASE user_db;

# データベースの確認
mysql> show databases;

# データベースの利用
mysql> use user_db;

# テーブルの確認
mysql> show tables;

# カラムの確認
mysql> SHOW COLUMNS FROM users;

＃ データの削除
mysql> DELETE FROM users;
```