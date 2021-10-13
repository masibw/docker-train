# baseイメージとしてgolangの1.17.2がインストールされたものを使う
FROM golang:1.17.2
# 現在の作業ディレクトリを/workdirに変更する
WORKDIR /workdir

# ホストからgo.modをコピーする
COPY go.mod .
RUN go mod download

# ホストからコマンド実行したディレクトリの中身をすべてコピーする
COPY . .

# ビルドする
RUN go build -o docker-train .

# docker runされた時に実行されるコマンドを指定する
ENTRYPOINT ["/workdir/docker-train"]