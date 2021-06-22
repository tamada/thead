package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

type options struct {
	bytes int
	lines int
	quiet bool
	help  bool
	args  []string
}

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS]  [FILEs...]
OPTIONS
	-c, --bytes <BYTES>             指定されたバイト数を各ファイルごとに出力する．
	-n, --lines <LINES>             指定された行数を各ファイルごとに出力する．
	-q, --quiet                     複数ファイルを出力する際の各ファイル名を表示しない．
	-h, --help                      このメッセージを出力する.
ARGUMENTS
	FILEs...                        カウント対象を指定する．
	DIRs...                         指定したディレクトリ内のファイルを入力ファイルとする．
`, programName)
}

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("thead", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.IntVarP(&opts.bytes, "byte", "c", 0, "指定されたバイト数を各ファイルごとに出力する．")
	flags.IntVarP(&opts.lines, "lines", "n", 0, "指定された行数を各ファイルごとに出力する．")
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "複数ファイルを出力する際の各ファイル名を表示しない．")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを出力する.")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	opts.args = flags.Args()[1:]

	return opts, nil
}

//標準出力
func stdout(o *options, filename string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	if o.lines > 0 {
		for i := 0; i < o.lines; i++ {
			s, err := rd.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(s)
		}
	} else {
		for i := 0; i < 10; i++ {
			s, err := rd.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(s)
		}
	}
	file.Close()

	// fmt.Println(filename)

}

func perform(o *options, filenames []string) {

	for _, filename := range filenames {
		stdout(o, filename)
	}
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	//オプション解析にエラーがあればエラー出力
	if err != nil {
		fmt.Printf("%s", err.Error())
		fmt.Println()
		fmt.Println(helpMessage(args[0]))
	}
	//オプション解析でhelpが指定されたらhelpメッセージを出力
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return 1
	}
	//エラーがない場合，通常の処理を行う
	perform(opts, opts.args)

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
