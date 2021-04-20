# ChitChat

このサンプルは第2章で紹介した、単純なフォーラムを実現するアプリケーションです。

このコードは第2章の説明と、1対1には対応していません。第2章の例よりも細かな処理を行っている部分があります。

以下に相違点をあげます。

* config.json で設定を変更することができます
* chitchat.log ファイルにログが書き出されます
* テストファイルがあります
* 構造体がすべて細かく定義されています（本文中では明確には記述されていませんでした）
* 一部の関数は（パッケージの一部ではなく）構造体のメソッドになっています

## 実行の手順

1. ＜Postgresを起動＞
2. createdb chitchat
3. psql -f data/setup.sql -d chitchat
4. go build
5. ./chitchat

brew install postgresql@9.6

echo 'export PATH="/usr/local/opt/postgresql@9.6/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

brew services start postgresql@9.6

createuser --superuser --createdb --createrole postgres

export LDFLAGS="-L/usr/local/opt/postgresql@9.6/lib"
export CPPFLAGS="-I/usr/local/opt/postgresql@9.6/include"
export PKG_CONFIG_PATH="/usr/local/opt/postgresql@9.6/lib/pkgconfig"


createuser -P -d gwp
take0014

createdb gwp

psql -U gwp -f setup.sql -d gwp

psql -U gwp -d gwp

\d+ posts


brew services list
brew services start postgresql@9.6
brew services stop postgresql@9.6

