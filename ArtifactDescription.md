# Artifact Description

**クレジット**：

* 

## 概要：{PDFビューアーの単語数カウント機能の追加}

以下，記述事項の説明．

* 改変対象OSSのURL:https://github.com/ledongthuc/pdf
  + このOSSはpdfファイルを内容を表示するOSSである
* 改変内容を説明
  + pdfファイル内にある単語(空白で区切られたもの)の数をカウントし、その単語数を出力するcount_words.goファイルの追加
  + testディレクトリに1~1000の単語をランダムに出力しsample.txtを作るgenerate_txt.goファイルの追加
  + 今回はpdfのプレーンテキストだけを出力する(改行やフォントを無視する)ファイルをmain.goとし、単語カウントの追加を行う
  

## クイックスタート

* Dockerイメージをpullしてrunする手順を具体的に示す．

```
docker pull h1910541/2024-h1910541-pdf:latest
docker run -it --rm --name console h1910541/2024-h1910541-pdf:latest
```

* コンテナの中に入った後
```
cd test
./test.sh
```

## 評価手順

例
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

**例外**

一応、Wordで作成しpdfにエクスポートした`test_pdf.pdf`を用意している。
こちらで`go run main.go test_pdf.pdf`を実行すると

```
The test pdf file No,1910541 Name Shinij Hikari  This is a sample pdf file. The number of words is 20. 
The PDF contains 20 words.
```
これにより,や.も空白に分けられていなければ、1つの単語として扱われることがわかる。

また空ファイル`sample2.pdf`も用意し実行すると

```

The PDF contains 0 words.
```
という結果を得る。

更に日本語のテキスト入れた`test_pdf2.pdf`を実行してみると

```
これは テスト pdf ファイル です。 
The PDF contains 5 words.
```
という単語数のカウントはあっているが、wordのように日本語は単語数=文字数、pdfは英単語なので1単語というカウントなので
単語数は14と表示される。

最後にヘッダーとフッターを追加した`test_pdf3.pdf`を実行すると

```
Test start   Test finish This is test. The number of words is 13.  
The PDF contains 13 words.
```
という結果を得る。

## 制限と展望

変更内容としてすごい簡単なものになってしまったのは、正直このリポジトリに何の機能がつけばいいものができるか思いつかなかったのが大きい。
もし時間があれば、pdfビューアするだけでなく編集可能にしたり、読み込んだ英単語を翻訳する機能を追加するなどやれることは多いが
今回は時間がなかったので断念した。
評価方法は可能な限り行ってみた。さすがに中国語などの特殊文字はきりがないので行わなかった。
ヘッダーとフッターの内容が先に来てしまうのは、なぜかわからなかった。
wordでは単語数の数え方が言語によって違うのはおそらく認識した文字コードなどでカウントする方法が変化しているので、最終的には
そこまで改変してみたいと思う。


## 更なる使い方（オプション）

もしこれが日本語にも対応すれば、あるpdfファイルの文字数などを調べられるのだが
これは英語の単語カウント数なので、英語でのプレゼンテーションで単語数を調べたいときぐらいに使えるかと思う。
