package main

func Example_help() {
	goMain([]string{"/some/path/of/thead", "-h"})
	// Output:
	// thead [OPTIONS]  [FILEs...]
	// OPTIONS
	// 	-c, --bytes <BYTES>      指定された文字数を各ファイルごとに出力する．
	// 	-n, --lines <LINES>      指定された行数を各ファイルごとに出力する．
	// 	-q, --quiet              複数ファイルを出力する際の各ファイル名を表示しない．
	// 	-s, --squeeze-blank      各ファイルごとに連続した空行を1行にする．
	// 	-h, --help               このメッセージを出力する.
	// ARGUMENTS
	// 	FILEs...                 カウント対象を指定する．theadはzip/tar/tar.gz/tar.bz2/jar/warファイルを受け付ける．
	// 	DIRs...                  指定したディレクトリ内のファイルを入力ファイルとする．
}
