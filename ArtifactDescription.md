# Artifact Description

**クレジット**：

* 

## 概要：{PDFビューアーの単語数カウント機能の追加}

以下，記述事項の説明．

* 改変対象OSSのURL:https://github.com/ledongthuc/pdf
  + このOSSはpdfファイルを内容を表示するOSSである
* 改変内容を説明
<<<<<<< HEAD
  + pdfファイル内にある単語(空白や改行で区切られたもの)の数をカウントし、その単語数を出力するcount_words.goファイルの追加
  + testディレクトリに1~1000の単語をランダムに出力しsample.txtを作るgenerate_txt.goファイルの追加
  + 今回はpdfのプレーンテキストだけを出力するファイルをmain.goとし、単語カウントの追加
=======
  + pdfファイル内にある単語(空白で区切られたもの)の数をカウントし、その単語数を出力するcount_words.goファイルの追加
  + testディレクトリに1~1000の単語をランダムに出力しsample.txtを作るgenerate_txt.goファイルの追加
  + 今回はpdfのプレーンテキストだけを出力する(改行やフォントを無視する)ファイルをmain.goとし、単語カウントの追加を行う
>>>>>>> 98bdc2c3d41879a3ce511c5901d6857bf97c2938
  

## クイックスタート

<<<<<<< HEAD
以下，記述事項の説明．

=======
>>>>>>> 98bdc2c3d41879a3ce511c5901d6857bf97c2938
* Dockerイメージをpullしてrunする手順を具体的に示す．

```
docker pull h1910541/2024-h1910541-pdf:latest
<<<<<<< HEAD
docker run -it --rm --name example h1910541/2024-h1910541-pdf:latest
```

* コンテナの中に入った後，最低限の動作チェックを行う方法を示す．
* シングルコマンドであることが望ましい．
  + 例えば，`make` を使う．

=======
docker run -it --rm --name h1910541/2024-h1910541-pdf:latest
```

* コンテナの中に入った後
>>>>>>> 98bdc2c3d41879a3ce511c5901d6857bf97c2938
```
cd test
./test.sh
```

## 評価手順

<<<<<<< HEAD
以下，記述事項の説明．


=======
例
>>>>>>> 98bdc2c3d41879a3ce511c5901d6857bf97c2938
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

<<<<<<< HEAD
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
=======
**例外**

一応、Wordで作成しpdfにエクスポートした`test_pdf.pdf`も用意している。
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
>>>>>>> 98bdc2c3d41879a3ce511c5901d6857bf97c2938
