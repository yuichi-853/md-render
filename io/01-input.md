# **Python講座 1 - Review 1~5**
**2024年12月9日**
<br>
**角田雄一**

<br>

# 試験対策について
推奨する勉強法としては、まずこのテキストと解説動画を見て、大まかな内容を把握する。そして、Stevensの授業スライドを最初から読んでいき、スライドにある練習問題を一つずつ解いていく。
<br>
プログラミングを学ぶ際には、どうしても言葉で説明されてもイメージしにくい概念が多く登場する。そのため、プログラミングをマスターするためには、自分で手を動かしてコードを書くことが何よりも重要である。だから、テキストを読み込むということより、コーディングの問題を解くことに重点をおいてほしい。
<br>
そもそもPythonは比較的簡単な言語であるとはいえ、通常習得には数ヶ月はかかる。だから勉強スケジュールを考えると相当ヤバめ。そのためテスト対策には緊張感を持って取り組んで欲しい。
<br>
とはいっても、一つ一つの内容は難しいものではないため、必要以上に苦手意識をもつ必要は無い。量をこなせば絶対身につくので、とりあえず手を動かして、コードを書くこと。
<br>
なお、自前のテキストは中間テストの範囲である基本事項の分のみ作り、それ以降の内容はStevensのスライドを使って解説する。

# Table of Contents

<!-- TOC GFM -->

* [L-01 Print, Comments](#l-01-print-comments)
    * [プログラムの実行](#プログラムの実行)
    * [Print（プリント出力）](#printプリント出力)
    * [Comments（コメントについて）](#commentsコメントについて)
* [L-02 StringのQuotes, Variables, User Input](#l-02-stringのquotes-variables-user-input)
    * [String（文字列）を囲むquotes（引用符）](#string文字列を囲むquotes引用符)
    * [Variables（変数）](#variables変数)
    * [User Input（ユーザー入力）](#user-inputユーザー入力)
    * [cf. データ型（Data Types）](#cf-データ型data-types)
* [L-03&04 Branching, Condition, Boolean, Format strings, Printing options](#l-0304-branching-condition-boolean-format-strings-printing-options)
    * [if, elif, else によるBranching（分岐）](#if-elif-else-によるbranching分岐)
    * [Condition（条件式）の書き方](#condition条件式の書き方)
    * [Boolean variable（論理値）](#boolean-variable論理値)
    * [Format strings（フォーマット文字列）](#format-stringsフォーマット文字列)
    * [Printing options （プリント出力のオプション）](#printing-options-プリント出力のオプション)
* [L-05 For loops & range()](#l-05-for-loops--range)
    * [1. For loops と range() 関数の基本](#1-for-loops-と-range-関数の基本)
    * [2. range() 関数の追加機能](#2-range-関数の追加機能)

<!-- TOC -->
<br>
<br>

# L-01 Print, Comments

## プログラムの実行
Pythonのプログラムは、上から順に１行ずつ実行される。

## Print（プリント出力）
値や文字列を画面に出力するために使う。
``` python
print("Hello, World!")
print(10 * 3)
```
> Hello, World!
<br>
> 30

## Comments（コメントについて）
プログラムはそもそも読みにくい。一流のエンジニアでも、コードだけ読んで理解するのは難しい。
だから、コードの中にコメントを書いて、コードの目的や処理内容を説明することが非常に重要となる。
<br>
つまり、あくまでコメントはプログラマのためのコードの補足説明文であり、実際のプログラムには影響を与えない。
<br>

- 複数行でコメントを書く場合は、コメント文の上下を `'''` または `"""` で囲む。なお、複数行でのコメントは docstring と呼ばれる。
- 一行でコメントを書く場合は、コメントの頭に `#` を使う。`#` を用いたコメントは、コードの末尾にも書くことができる。

``` python
'''
複数行に渡るコメントは、このように書くことができる
複数行
だよ
'''

# このように、一行のコメントも書くことができる
print("Hello, World!") # コードのあとにコメントを書くこともできる
```
<br>
<br>

# L-02 StringのQuotes, Variables, User Input

## String（文字列）を囲むquotes（引用符）
Pythonで文字列（string）は、文字や数字、記号の並びのことである。
文字列を表現するには４つ方法がある。

1. シングルクォート（`'`string`'`）
    - アポストロフィー `'` をそのまま含められない。含める場合は直前にバックスラッシュ `\` が必要。(スラッシュ `/` との違いに注意)
``` python
# シングルクォート
'Hello'
'It\'s a pen'  # アポストロフィを含める場合は \ でエスケープ
```

2. ダブルクォート（`"`string`"`）
    - アポストロフィー `'` をそのまま含められる。
``` python
# ダブルクォート
"Hello"
"It's a pen"  # アポストロフィはそのまま使用可能
"He said \"Hello\""  # ダブルクォートを含む場合は \ でエスケープ
```

3. 三重クォート（`'''`string`'''` または `"""`string`"""`）
    - シングルクォート `'` またはダブルクォート `"` の両方をそのまま含められる。
    - 複数行の文字列を表現可能。
``` python
# 三重クォート
# シングルクォートとダブルクォートの両方をそのまま使用可能
'''It's a "pen".'''
'''It's a 'pen'.'''
"""It's a 'pen'."""
"""It's a "pen"."""

# 複数行の文字列を表現可能
text = '''This is
a multi-line
string.'''
```

4. 文字列の結合（String Concatenation）
    - `+` 演算子を使って文字列を結合することができる。
``` python
print('H' + 'e' + 'l' + 'l' + 'o')
print('yah' + 'ooo!')
print('Hi ' + 'there' + ' ' + 'Bob!') # スペースを入れることもできる
```
> Hello
<br>
> yahooo!
<br>
> Hi there Bob!

## Variables（変数）
変数は、データを格納するための箱のようなものである。変数には名前がついており、その名前を使ってデータを参照することができる。
``` python
# 変数 "a" に値を代入する
a = 1 + 2
print(a)
```
> 3
``` python
a = 1 + 2
a = 3 + 4   # 変数に新しい値を代入するため、元の値は上書きされて消える
print(a)
```
> 7
``` python
a = 1
a = a + 2   # 右辺は「a の値を評価した結果」の 1 を使って， 1 + 2 となる
print(a)
```
> 3

通常、Variableは中のデータの情報を表す名前をつける。
<br>
Variable（変数）にはどんなデータが入っているのかがわかりやすく、ほかと明確に区別できる名前をつけることが非常に重要になってくる。

その上で、Pythonの変数名は以下のルールに従う必要がある。
- 使用可能な文字
    - アルファベット: 大文字・小文字（`A-Z`、`a-z`）
    - 数字: 0～9（ただし、先頭には使えない）
    - アンダースコア: `_`（自由に使えるが、特別な意味を持つ場合もある）
- 使用不可な文字
    - スペースや記号（例: `@`, `#`, `!`）は使用不可。
    - Pythonの文法で使用されている予約語（`if`, `else`, `for` など）は使えない。
- 命名の注意点
    - 先頭に数字は不可
    - 大文字と小文字は区別される（'apple' と 'Apple' は別の変数）
``` python
# 有効な変数名
x1 = 10                # シンプルな例
my_variable = 20       # 小文字とアンダースコアを使った例（スネークケース）
myVariable123 = 30     # キャメルケースを使った例
MY_VARIABLE = 40       # 大文字のみ（定数として使うことが多い）

# 無効な変数名
2variable = 10      # 無効: 先頭に数字が使われている
my-variable = 20    # 無効: ハイフン（`-`）は使用できない
my variable = 30    # 無効: スペースが含まれている
if = 40             # 無効: `if`はPythonの予約語
class = 50          # 無効: `class`も予約語
my@variable = 60    # 無効: `@`のような記号は使用できない
variable! = 70      # 無効: `!`のような記号も使用不可
```

## User Input（ユーザー入力）
ユーザーからの入力を受け取るには、`input()` 関数を使う。
``` python
# ユーザーから名前を入力してもらい、入力後、それを表示する
name = input("What is your name? ")
print("Hello, " + name + "!")
```
input()の引数（カッコの中の文字列。ここでは "What is your name? "）は、"**prompt**"（刺激） と呼ばれ、ユーザーに対して表示されるメッセージである。

## cf. データ型（Data Types）
Pythonで扱えるデータ型として、以下のものがある。

| データ型・関数（変換のため） | 説明                              | 表記例                 |
|---------------------------|-----------------------------------|------------------------|
| 整数: `int()`             | 整数を表す                       | `10`, `-3`, `0`       |
| 浮動小数点数: `float()`   | 小数点以下を含む数値を表す       | `3.14`, `-0.1`, `0.0` |
| 文字列: `str()`           | 文字の並びを表す **（`string()`でないことに注意！）** | `"Hello"`, `'World'`, `"123"` |
| 論理値: `bool()`          | 条件判断に使います               | `True`, `False`       |

変数に代入されているオブジェクトの型を確認するには、`type()` 関数を使う。

``` python
a = 1
b = 1/2
c = "ABC"
print(type(a))
print(type(b))
print(type(c))
```
> <class 'int'>
<br>
> <class 'float'>
<br>
> <class 'str'>

関数（`int()`や`str()`）はどう使うのか？変換とはどういう意味か？
``` python
# 浮動小数点数を整数に変換（小数点以下切り捨て）
number = int(3.99)
print(number)  # 出力: 3

# 文字列を浮動小数点数に変換
decimal = float("3.14")
print(decimal)  # 出力: 3.14

# 整数を文字列に変換
text = str(123)
print(text)  # 出力: "123"
```
応用すると、ユーザーからの入力を数値として扱いたい場合に、`int()` や `float()` を使って変換することができる。（ユーザーからの入力は、`input()` 関数で文字列として扱われるから）
``` python
# ユーザーが任意の数値を入力する

# ユーザーからの入力は文字列として扱われてしまう
num1 = input("Enter a number: ")

# ユーザーからの入力を整数としてnum2に代入することができる
num2 = int(input("Enter another number: "))
```

<br>
<br>

# L-03&04 Branching, Condition, Boolean, Format strings, Printing options

## if, elif, else によるBranching（分岐）
**条件式のあとに、コロン `:` をつけるのを忘れないように！**
また、条件式が成立した場合に実行されるブロックは、読みやすさのためインデント（スペース４つ分、左側を空ける）で表現される。
``` python
# 例１
if 条件式:
    条件式が成立した場合に実行されるブロック
else:
    条件式が成立しなかった場合に実行されるブロック

# 例２
if 条件式１:
    条件式１が成立した場合に実行されるブロック
elif 条件式２:
    条件式１が成立しなかった場合で、条件式２が成立した場合に実行されるブロック
    elif は、else if の略
else:
    すべての条件式が成立しなかった場合に実行されるブロック
```

``` python
# 入力された数値が正の数か、負の数か、0かを判定するプログラム

# 数値を入力
num = int(input("Enter a number: "))

# 条件判定
if num > 0:
    print("The number is positive.")  # 正の数の場合
elif num < 0:
    print("The number is negative.")  # 負の数の場合
else:
    print("The number is zero.")      # 0の場合

```

## Condition（条件式）の書き方
### 1. 数値の比較
条件の代表的なものは数値の比較である。比較するための演算子は以下の通り。
- `==`: 等しい
- `!=`: 等しくない
- `>`: より大きい
- `<`: より小さい
- `>=`: 以上
- `<=`: 以下

数学の記号と基本的に同じだが、**等しいことを表す記号は `=` ではなく `==` であることに注意！**

### 2. String comparison（文字列の比較）
文字列は、辞書順（lexicographical）で比較される。
``` python
if "apple" > "banana":
    print("apple は辞書順で banana より後に位置します")
else:
    print("apple は辞書順で banana より先に位置します")
```
> apple は辞書順で banana より先に位置します

### 3. Compound conditions（条件の組み合わせ）
複数の条件を組み合わせるために、論理演算子 `and`, `or`, `not` が用意されている。
数学の式と同じように、括弧 `()` の中の条件式が最初に評価される。
* `and`: 日本語で言うと「かつ」。両方の条件が成立する場合にTrue。
``` python
# 天気（weather）が晴れ（sunny）かつ、気温（temp）が暖かい（warm）場合
if weather == 'sunny' and temp == 'warm':
```
+ `or`: 日本語で言うと「または」。どちらかの条件が成立する場合にTrue。
``` python
# 天気（weather）が雨（rainy）または、気温（temp）が寒い（cold）場合
if weather == 'rainy' or temp == 'cold':
```
- `not`: 日本語で言うと「でない」。単に条件を反転させる。
``` python
# 天気（weather）が晴れ（sunny）かつ、気温（temp）が暖かい（warm）でない場合
if not (weather == 'sunny' and temp == 'warm'):
```

### 4. 演算の優先順位
なお、Pythonでは演算子に優先順位があり、

- 算術演算（`+`, `-`, `*`, `/`, `**`）は 比較演算（`==`, `<`, `>`, `<=`, `>=`, `!=`）よりも優先されます。
- 比較演算（`==`, `<`, `>`, `<=`, `>=`, `!=`）は 論理演算（`and`, `or`, `not`）よりも優先されます。

``` python
# 算術演算が比較演算より優先
result = 3 + 2 > 4  # まず 3 + 2 が計算され、その後 > 4 が評価される
print(result)  # 出力: True

# 比較演算が論理演算より優先
result = 5 > 3 and 2 < 4  # まず 5 > 3 と 2 < 4 が評価され、その後 and が適用される
print(result)  # 出力: True
```
ここらへんは直感的にわかると思うので、難しく考える必要はない。
<br>
<br>
そして、結構大事なのは**プログラムを読みやすくするためにカッコ `()` で明示的に優先順位を示すこと**である。
``` python
# どちらも動作は同じだが、下のコードの方が読みやすいだろう
a == 1 and b != 0
(a == 1) and (b != 0)
```

## Boolean variable（論理値）
Boolean（ブーリアン）とは、True（真）または False（偽）の値を持つデータ型である。（`int` や `str` と同じ扱い）
普通に、条件が成立するかどうかを判定するために使われるだけで、難しく考える必要はない。
条件式の結果も、True または False という論理値を返す（下の例参照）。
``` python
x = 15
if x >= 10: # x が 10 以上の場合、条件式は True を返す
two_digits = True # two_digits に boolean値の True を代入
else:
two_digits = False # two_digits に boolean値の False を代入

print(two_digits) # 出力: True
```

ややこしい例を挙げると、以下のようなものがある。
``` python
# 例１: 数字が 0 かどうかを判定するプログラム
num = 0
is_zero = num == 0  # num が 0 と等しい場合、is_zero に True を代入
print(is_zero)  # 出力: True

# 例２: 数値が50以上かつ100未満であることを判定するプログラム
num = 70
is_between_50_and_100 = 50 <= num and num < 100
print(is_between_50_and_100)  # 出力: True

# 例３: 子供で、かつ女性であることを判定するプログラム
age = 13
is_male = False
is_minor_and_female = age < 18 and not is_male
# わかりやすく書くと
# is_minor_and_female = (age < 18) and (not is_male) となる
print(is_minor_and_female)  # 出力: True

# 例４: Boolean variable は、if文やelif文の条件式としても使える
two_digits = True
if two_digits:
    print("The number is two digits long.")
```
``` python
# なお、よくある間違いとして、
if a <= b <= c:
# という条件式があるが、「三方向比較」はPythonではダメである。正しくは、
if a <= b and b <= c:
```

## Format strings（フォーマット文字列）
Pythonのフォーマット文字列（f-strings）は、値を効率的かつ簡単に整形して出力するための方法。フォーマット文字列では、f を文字列の前に置き、文字列中に {} を使って変数や式を埋め込むことができる。

- `f"..."` の形式で文字列を定義
- `{}` の中に変数や式を入れて出力

イメージとしては、`{}` の中の変数や式が、実行時にその値に置き換わり、それ以外の `" "` の中の文字列はそのまんま出力される、と考えるとわかりやすい。
``` python
x = 5
y = 10
print(f'x = {x}')
print(f'10x^2 = {10 * x ** 2}')
print(f'x = {x}, y = {y}, sum = {x + y}')   # 複数の {} を使うことも可能
```
> x = 5
<br>
> 10x^2 = 250
<br>
> x = 5, y = 10, sum = 15

## Printing options （プリント出力のオプション）
`print()` 関数には、いくつかのオプションが用意されている。

### 1. 複数の要素をカンマ ',' を使って区切って出力する
カンマ `,` を使って複数の要素を区切ると、自動的にスペースが挿入される。
``` python
print("I'm", 2024-1976, "years old.")
```
> I'm 48 years old.

### 3. sep を使って区切り文字を指定する
`sep` パラメータを使うことで、カンマ `,` で区切った部分の出力で、自動的にスペースが挿入される部分を自由にカスタマイズすることができる。
``` python
print("I'm", 2024-1976, "years old.", sep='/')
print('Direcroty:', 'home', 'user', sep='/')
```
> I'm/48/years old.
<br>
> Direcroty:/home/user

### 3. end を使って同じ行に出力する
通常 `print()` を使うと、各出力のあとに自動的に改行が挿入される。しかし、`end=''` を指定すると、改行されず、次の出力が同じ行に続けて表示される。
``` python
print('Installing your program...', end='')
print('done')
```
> Installing your program...done

### sep と end の応用
`sep` と `end` を組み合わせることで、さらに出力をカスタマイズすることができる。
``` python
print('Installing', '...', sep='', end='Done\n')
print('Starting the next operation')
```
> Installing...Done
<br>
> Starting the next operation

### 4. Escape sequence
Escape sequenceは、文字列内に特殊な文字や機能を埋め込むための2文字のコードである。Pythonでは、`\`（バックスラッシュ）を使う。
<br>

主なEscape sequenceは以下の通り。

- `\n`: 改行を挿入する
``` python
print('1\n2\n3')

# 出力
# 1
# 2
# 3
```

- `\t`: タブ（スペースより広い空白、一般的にスペース４個分）を挿入する
``` python
print('1\t2\t3')

# 出力
# 1    2    3
```

- `\\`: バックスラッシュ `\` を出力する（バックスラッシュのままではエラーになる）
``` python
print('C:\\Program Files\\Python')

# 出力
# C:\Program Files\Python
```

- `\'`: シングルクォート `'` を出力する
``` python
# print('It's a sunny day.') だとシングルクォートで囲まれた中でシングルクォートが使われているためエラーになる！

print('It\'s a sunny day.')

# 出力
# It's a sunny day.
```

- `\"`: ダブルクォート `"` を出力する
``` python
print("She said \"Hello\"")

# 出力
# She said "Hello"
```
# L-05 For loops & range()
プログラミングによりコンピュータが威力を発揮するのは、多くの処理を速い速度で実施できるときである。ここでは、for文を使って、繰り返しを自動化する方法を紹介する。

## 1. For loops と range() 関数の基本
`for` ループとは、**特定の回数**だけコードを繰り返したいときに使用する。
``` python
# 文法
for ループ変数 in range(回数):
    body（繰り返し処理）

# 例１
for i in range(5):
    print('hi', end='') # 出力: hihihihihi

# 例２
for i in range(3):
    print(f'#{i + 1}', end='') # 出力: #1#2#3
```
`range()` という関数は、0から指定した回数-1までの範囲の数値を生成する。つまり、例１のrange(5)という関数は、0 から「添え字の値 - 1」 の「4」までの 5 個の`値`を生成する関数で，for 文は生成された値を変数 i に入れて，body（繰り返し処理）の部分を繰り返し実行する。
<br>
したがって、range(n) の範囲でのループ変数 i のカウントの流れは以下の通りとなる。
- 1回目: i = 0
- 2回目: i = 1
- 3回目: i = 2
- n回目: i = n - 1

これを踏まえると、例２の処理が理解できるだろう。

``` python
# 例３: 正の整数を入力し、1 からその数までの合計を求めるプログラム
n = int(input('Enter positive number: '))
if n > 0:
    sum = 0
    for x in range(n):
        sum = sum + x + 1
    print(f'The sum from 1 to {n} is {sum}')
else:
    print('The number is not positive.')
```

## 2. range() 関数の追加機能
### range(start, end) の使い方
`range(start, end)` を使うことで、特定の範囲内の数値を生成できる。
``` python
# 文法
for i in range(start, end):
    処理

```
- `start`: 開始値（含まれる）
- `end`: 終了値（含まれない）
- 生成される値: `start` から `end - 1` までの数値
- 繰り返し回数: `end - start`

``` python
# 例１
for i in range(1, 5):
    print(i, end='') 
# 出力: 1234

# 例２
for i in range(10, 15):
    print(i, end=', ')
# 出力: 10, 11, 12, 13, 14

# 例３: 100から200までの合計を求めるプログラム
sum = 0
for i in range(100, 201):   # 100～200を生成
    sum = sum + i 
print(f'The sum from 100 to 200 is {sum}')
# 出力: The sum from 100 to 200 is 15150
```
### range(start, end, increment) の使い方
`range(start, end, increment)` を使うことで、指定したincrement（間隔）で値を生成できる。
``` python
# 文法
for i in range(start, end, increment):
    処理
```
- `increment`: 増加量（負の値にすると、減少する値も生成可能）
- 繰り返し回数: `floor((end - start) / increment)`

``` python
# 例: 3の倍数の合計を求めるプログラム
sum = 0
for i in range(3, 301, 3):  # 3～300の3の倍数を生成
    sum = sum + i
print(f'The sum of the first 100 multiples of 3 is {sum}.')
# 出力: The sum of the first 100 multiples of 3 is 15050.
```

なお、実際は `range(start, end)` や、`range(start, end, increment)` を使うことはあまりなく、覚えておく必要はそんなに無いのだが、試験に出るかどうかとなると話は全く変わってくるため、一応理解しておく必要がある。
