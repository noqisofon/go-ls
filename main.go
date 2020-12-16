package main

import (
    "fmt"
    "os"

    flag "github.com/spf13/pflag"
)

type file_type int
const (
    unknown file_type = iota
    fifo
    chardev
    directory
    blockdev
    normal
    symbolic_link
    sock
    whiteout
    arg_directory
)

type file_info struct {
    /* ファイルの名前 */
    name   string

    /* リンクの名前 */
    link_name string

    /* 絶対的ファイル名 */
    absolute_name string

    stat os.FileInfo

    file_type file_type
}

func main() {
    var (
    	all        bool
	all_most   bool
	author     bool
	escape     bool
	block_size uint
    )

    flag.BoolVarP( &all       , "all"       , "a", false, "do not ignore entries starting with ."           )
    flag.BoolVarP( &all_most  , "all-most"  , "A", false, "do not list implied . and .."                    )
    flag.BoolVar(  &author    , "auther"         , false, "with -l, print the author of each file"          )
    flag.BoolVarP( &escape    , "escape"    , "b", false, "print C-style escapes for nongraphic characters" )
    flag.UintVar(  &block_size, "block-size"     ,     4, "with -l, scale sizes by SIZE when printing them; e.g., '--block-size=M'; see SIZE format below");

    flag.Parse()

    fmt.Println( "all       : ", all        )
    fmt.Println( "all_most  : ", all_most   )
    fmt.Println( "author    : ", author     )
    fmt.Println( "escape    : ", escape     )
    fmt.Println( "block_size: ", block_size )

    // dir, err := os.Getwd()
}