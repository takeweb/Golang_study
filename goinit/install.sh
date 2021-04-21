#!/usr/bin/env bash

# 設定ファイルを配置
mkdir -p ~/.config/goinit
cp ./config.json ~/.config/goinit/

# 普通に使えるコマンド化
go install goinit
