# thead
headコマンドの機能拡張 

<img src="./img/thead_icon.svg" width="100">

## Description
theadコマンドは指定されたファイルおよび指定されたディレクトリ内のファイルの先頭の内容を10行分出力する．


## Usage
```
thead [OPTIONS...] [FILEs...|DIRs...]  
　OPTIONS:  
    -c, --bytes <BYTES>     指定された文字数を各ファイルごとに出力する．
    -n, --lines <LINES>     指定された行数を各ファイルごとに出力する．
    -q, --quiet             複数ファイルを出力する際の各ファイル名を表示しない．
    -s, --squeeze-blank     各ファイルごとに連続した空行を1行にする．  
    -h, --help              このメッセージを出力する.  
    
　ARGUMENTS   
　　FILEs...
　　DIRs... 
```
  
## icon
<img src="./img/thead_icon.svg" width="100">  

アイコンの画像は [freesvg.org](https://freesvg.org/rejons-head-vector)から取得.
