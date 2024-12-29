FROM ubuntu:24.04

ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone
RUN apt update && apt upgrade -y

# カレントディレクトリ
WORKDIR /artifact

# 必要なAPTパッケージを適当にインストール
RUN apt install -y git golang nano curl unzip enscript ghostscript
# Gitリポジトリを展開しても良い
RUN git clone https://github.com/oss-experiment-uec/2024-h1910541-pdf .

RUN go mod init github.com/oss-experiment-uec/2024-h1910541-pdf || true

# 依存関係をダウンロード
RUN go mod tidy

# Dockerfileを実行する場所からファイルをコピーする場合
# COPY . /artifact