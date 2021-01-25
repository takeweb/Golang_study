◆暗号キーの生成
cd cmd/keygen
./keygen

◆暗号キーを環境変数「ANGO_KEY」へ設定
export ANGO_KEY=6YwwhWFItyH9/nL2A/hbBa0aFC01Z4HxViwoW/4CCzw=

◆暗号化
cd cmd/ango
secret.txtに元の文字列を用意する。改行、日本語OK
cat secret.txt | ./ango > ../fukugo/encrypted.secret.txt

◆復号化
cd cmd/fukugo
cat encrypted.secret.txt | ./fukugo > secret.txt