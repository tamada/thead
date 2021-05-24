# thead
二地点の緯度経度から距離の算出を行う．  

<img src="./img/thead_icon.svg" width="100">

## Description
私は地理情報を用いた研究を行っている．その中で二点間の緯度経度から距離の算出を行うことがあった．しかし，距離の算出方法は一つではない．私が行った研究内では，harversine公式を使用し，距離を算出した．  
本ソフトウェアでは二点間の緯度経度から距離を算出するが，harversine公式以外にヒュベニの公式での距離算出を行えるソフトウェアの開発を行う．またosm(Open Street Map)データを想定したファイルの読み込み，距離を算出するファイルを出力するなどの機能を加える．

## Usage
```
thead [OPTIONS...] [FILEs...|DIRs...]  
　OPTIONS:  
　　-h,--help このメッセージを出力する.  
    
　ARGUMENTS   
　　FILEs... 
　　DIRs... 
```
  
## icon
<img src="./img/thead_icon.svg" width="100">  

アイコンの画像は [freesvg.org](https://freesvg.org/rejons-head-vector)から取得.
