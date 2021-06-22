package main

func Example_help() {
	goMain([]string{"/some/path/of/thead", "-h"})
	// Output:
	// thead [OPTIONS]  [FILEs...]
	// OPTIONS
	// 	-c, --bytes <BYTES>             指定されたバイト数を各ファイルごとに出力する．
	// 	-n, --lines <LINES>             指定された行数を各ファイルごとに出力する．
	// 	-q, --quiet                     複数ファイルを出力する際の各ファイル名を表示しない．
	// 	-h, --help                      このメッセージを出力する.
	// ARGUMENTS
	// 	FILEs...                        カウント対象を指定する．
	// 	DIRs...                         指定したディレクトリ内のファイルを入力ファイルとする．
}
