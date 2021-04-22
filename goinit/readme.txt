◆インストール(設定ファイルを所定の位置に配置、バイナリをGOPATH/binに追加)
./install.sh

◆「hogehoge」というモジュールを作成したい場合
goinit -m hogehoge
※「-m」省略時は「test」というモジュール名になる。

◆「hogehoge」というモジュールを任意のフォルダ「/usr/local/dev」に作成したい場合
goinit -m hogehoge -d /usr/local/dev
※「-d」省略時は「~/dev/Golang_study/」というディレクトリ配下に作成される。

◆「hogehoge」というモジュールで、main関数を「hogehoge.go」というファイル名に作成したい場合
goinit -m hogehoge -f hogehoge
※「-f」省略時は「main.go」というファイル名で作成される。

以上