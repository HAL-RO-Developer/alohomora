# web template

goでweb開発をするためのオレなりのtemplate

## 使用パッケージ
- gin
- gorm
- sql-migrate
- golem


## 初期設定
```sh
// ファイルの取得
git clone git@github.com:makki0205/go_template.git ./{Repository名}
// DBファイル
cp dbconfig.yml.template dbconfig.yml
// 依存ライブラリのインストール
glide up
```

## 起動
```
go run main.go
```