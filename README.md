[![CI](https://github.com/akanatr/thead/actions/workflows/build.yml/badge.svg)](https://github.com/akanatr/thead/actions/workflows/build.yml)

# thead
headコマンドの機能拡張 

<img src="./img/thead_icon.svg" width="100">

## Description
既存のheadコマンドは複数のファイル指定し，各ファイルごとの先頭10行分の内容を出力することが可能である．
しかし，特定のディレクトリ直下のファイル全てを指定する際には，正規表現を工夫する必要があり面倒である．
そこでtheadではheadコマンドの機能の一部に加え，ディレクトリを指定するだけでその直下のファイル全てを対象とする機能を追加する．


## Usage
```
thead [OPTIONS...] <FILEs...|DIRs...>  
　OPTIONS:  
    -c, --bytes <BYTES>     指定された文字数を各ファイルごとに出力する．
    -n, --lines <LINES>     指定された行数を各ファイルごとに出力する．
    -q, --quiet             複数ファイルを出力する際の各ファイル名を表示しない．
    -s, --squeeze-blank     各ファイルごとに連続した空行を1行にする．  
    -h, --help              このメッセージを出力する.  
    
　ARGUMENTS   
　　FILEs...                 カウント対象を指定する．theadはzip/tar/tar.gz/tar.bz2/jar/warファイルを受け付ける．
　　DIRs...                  指定したディレクトリ内のファイルを入力ファイルとする．
```
  
## icon
<img src="./img/thead_icon.svg" width="100">  

アイコンの画像は [freesvg.org](https://freesvg.org/rejons-head-vector)から取得.
