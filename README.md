# gRPC通信サンプル

Go と Node.js/TypeScript での gRPC 通信実装例です。以下の2つのシナリオをサポート：

1. **Go ⇄ Go**: `client.go` と `server.go` で Go 同士の gRPC 疎通確認
2. **Node.js ⇄ Go**: `src/client.ts` と `server.go` で Node.js/TypeScript クライアントと Go サーバーの gRPC 疎通確認

## 概要

### Go 同士（client.go + server.go）
- サーバー: ポート 8080 でリッスン
- クライアント: サーバーに接続してメッセージをエコーバック

### Node.js + Go（src/client.ts + server.go）
- サーバー: 同じポート 8080 でリッスン（Go サーバー共通）
- クライアント: Node.js/TypeScript で実装、Protocol Buffers で通信

## 前提条件

- Go 1.21 以上
- Node.js 18 以上
- npm
- protoc（Protocol Buffers コンパイラ）

## インストール

```bash
# Go と Node.js 依存関係をインストール
go mod download
npm install
```

## セットアップ

### Go サーバー用コード生成（初回のみ）

```bash
protoc --go_out=. --go-grpc_out=. message.proto
```

このコマンドで自動生成されます：
- `protoc/pb/message.pb.go` - Protocol Buffers メッセージコード
- `protoc/pb/message_grpc.pb.go` - gRPC サービスコード

Node.js クライアント側は`@grpc/proto-loader`で動的ロードするため、別途生成は不要です。

## 実行方法

### シナリオ1: Go ⇄ Go 通信確認

ターミナル1（サーバー）:
```bash
go run server.go
```

ターミナル2（クライアント）:
```bash
go run client.go
```

出力例：
```
Server listening on :8080
Client sent - ID: 42, Content: Hello from client!
Server received - ID: 42, Content: Hello from client!
Server sent - ID: 43, Content: Echo: Hello from client!
Client received - ID: 43, Content: Echo: Hello from client!
```

### シナリオ2: Node.js ⇄ Go 通信確認

ターミナル1（Go サーバー）:
```bash
go run server.go
```

ターミナル2（Node.js/TypeScript クライアント）:
```bash
npm run dev
```

## テストスクリプト

両シナリオを自動実行できます：

```bash
# Node.js クライアントでテスト（デフォルト）
./run_test.sh

# または明示的に指定
./run_test.sh node

# Go クライアントでテスト
./run_test.sh go

# 使用方法の表示
./run_test.sh help
```

**出力例（Go クライアント）:**
```
=== Go クライアント（Go同士通信） ===
Server listening on :8080
Client sent - ID: 42, Content: Hello from client!
Server received - ID: 42, Content: Hello from client!
Server sent - ID: 43, Content: Echo: Hello from client!
Client received - ID: 43, Content: Echo: Hello from client!
```

## ファイル構成

```
.
├── message.proto                 # Protocol Buffers メッセージ定義
├── server.go                     # Go gRPC サーバー実装
├── client.go                     # Go gRPC クライアント実装（Go同士通信確認用）
├── src/
│   └── client.ts                 # Node.js/TypeScript gRPC クライアント
├── protoc/pb/
│   ├── message.pb.go             # 自動生成（protoc）
│   └── message_grpc.pb.go        # 自動生成（protoc）
├── package.json                  # Node.js プロジェクト設定
├── tsconfig.json                 # TypeScript コンパイラ設定
├── go.mod / go.sum               # Go モジュール依存情報
└── run_test.sh                   # テストスクリプト（go/node の両方に対応）
```

## トラブルシューティング

| エラー | 解決方法 |
|--------|--------|
| `protoc: command not found` | protoc をインストール |
| `port 8080 already in use` | `lsof -i :8080` で確認し、プロセスを終了 |
| `Cannot find module '@grpc/grpc-js'` | `npm install` を実行 |

## 参考リンク

- [Protocol Buffers ドキュメント](https://developers.google.com/protocol-buffers)
- [Go gRPC](https://pkg.go.dev/google.golang.org/grpc)
- [Node.js gRPC](https://www.npmjs.com/package/@grpc/grpc-js)
