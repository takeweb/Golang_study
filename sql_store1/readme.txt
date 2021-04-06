1. Postgresを起動
    brew install postgresql@9.6

    echo 'export PATH="/usr/local/opt/postgresql@9.6/bin:$PATH"' >> ~/.zshrc
    source ~/.zshrc

    brew services start postgresql@9.6

    createuser --superuser --createdb --createrole postgres

    export LDFLAGS="-L/usr/local/opt/postgresql@9.6/lib"
    export CPPFLAGS="-I/usr/local/opt/postgresql@9.6/include"
    export PKG_CONFIG_PATH="/usr/local/opt/postgresql@9.6/lib/pkgconfig"

2. DBユーザ「gwp」を作成
    createuser -P -d gwp
    （passwd gwp）

3. DB「gwp」を作成
    creatdb gwp

4. テーブル初期化用SQL「setup.sql」を実行
    psql -U gwp -f setup.sql -d gwp
   （一番最初は ERRORが表示されるが、そのままでOK)
   psql -U gwp -d gwp
   \d+ posts

5. go get "github.com/lib/pq"

6. go run store.go
    or
   go build
   ./store

以上