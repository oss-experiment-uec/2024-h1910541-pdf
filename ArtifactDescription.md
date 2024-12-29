# Artifact Description

**クレジット**：

* 

## 概要：{PDFビューアーの単語数カウント機能の追加}

以下，記述事項の説明．

* 改変対象OSSのURL:https://github.com/ledongthuc/pdf
  + このOSSはpdfファイルを内容を表示するOSSである
* 改変内容を説明
  + pdfファイル内にある単語(空白や改行で区切られたもの)の数をカウントし、その単語数を出力するcount_words.goファイルの追加
  + testディレクトリに1~1000の単語をランダムに出力しsample.txtを作るgenerate_txt.goファイルの追加
  + 今回はpdfのプレーンテキストだけを出力するファイルをmain.goとし、単語カウントの追加
  

## クイックスタート

以下，記述事項の説明．

* Dockerイメージをpullしてrunする手順を具体的に示す．

```
docker pull h1910541/2024-h1910541-pdf:latest
docker run -it --rm --name example h1910541/2024-h1910541-pdf:latest
```

* コンテナの中に入った後，最低限の動作チェックを行う方法を示す．
* シングルコマンドであることが望ましい．
  + 例えば，`make` を使う．

```
cd test
./test.sh
```

## 評価手順

以下，記述事項の説明．


1. 入力`go run generate_text.go` を与えて実行する．

```
go run generate_text.go
```

を実行すると

```
Generated sample.txt with 6 words.
```

が出力される．これは単語数6のsample.txtを作成しました。
内容を`cat sample.txt`で確認すると

```
cxdj p zryyfzkecq jiainriysh ttgvci phh
```
と単語数6のsample.txtが作成される。

2. 入力`enscript sample.txt --no-header -o - | ps2pdf - sample.pdf `を与えて実行する

```
enscript sample.txt --no-header -o - | ps2pdf - sample.pdf 
```
を実行すると
```
[ 1 page * 1 copy ] left in -
```
と出力される。これでsample.txtはsample.pdfに変換されます。なお、`--no-header`を入れて出力内容に作成した日付を消します。

3. 入力`go run main.go sample.pdf`を与えて実行する

```
go run main.go sample.pdf 
```
を実行すると
```
cxdj p zryyfzkecq jiainriysh ttgvci phh
The PDF contains 6 words.
```
と出力される。上記の単語数が6でカウント数は6となっている。

## 制限と展望

以下，記述事項の説明．

* 改変内容や評価方法について諦めた点があれば説明する．
  + 意図的に行わなかった事柄も含む．
* 諦めた理由は問わないが，理由の説明は要する．
* 時間に余裕があればやりたかった事柄も説明する．
* 何もないなら「特になし」と明記する．

## 更なる使い方（オプション）

以下，記述事項の説明．

* より現実的な応用例や利用例を説明する．
* ソフトウェアを使いたくさせる説明が理想的．
* この節の見出しは適当に変えてよいし，複数の節に分けてもよい．
* 必須ではないが，書けるなら書いた方が評価者には好印象を与える．
