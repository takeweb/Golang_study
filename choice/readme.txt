// モジュールの作成
% mkdir choice
% cd choice
% go mod init choice
go: creating new go.mod: module choice

// モジュールのビルド
main.goを作成後
% go build
→choiceというフォルダ名を冠した実行ファイルが作成される

// モジュールの実行
./choice 29 30 31 32 44
→ランダムで引数から選択されて表示する
→44
→29
