#!/bin/bash

# デフォルト値: node
CLIENT_TYPE="${1:-node}"

# 使用方法
usage() {
  echo "Usage: $0 [go|node]"
  echo "  go   - Go クライアントでテスト（Go同士の通信）"
  echo "  node - Node.js クライアントでテスト（Node.js-Go の通信）"
  echo ""
  echo "デフォルト: node"
  exit 1
}

# 引数の検証
case "$CLIENT_TYPE" in
  go)
    echo "=== Go クライアント（Go同士通信） ==="
    ;;
  node)
    echo "=== Node.js クライアント（Node.js-Go 通信） ==="
    ;;
  *)
    echo "エラー: 不正な引数 '$CLIENT_TYPE'"
    usage
    ;;
esac

# サーバーを起動
go run server.go &
SERVER_PID=$!

# サーバーが起動するのを待機
sleep 2

# クライアントを実行
if [ "$CLIENT_TYPE" = "go" ]; then
  go run client.go
else
  npm run dev
fi

# サーバーを終了
kill $SERVER_PID 2>/dev/null || true
