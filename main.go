package main

import (
    "fmt"
    "path/filepath"
    "os"
    "strconv"
    "strings"
)

var stdout_file_no : int = os.Stdout.Fd()

// プログラムにおける exit_code
const (
    EXIT_SUCCESS = 0
    EXIT_FAILURE = 1
)

// C 言語の size_t
type size_t uint

// ファイル・タイプ
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

// ls の表示フォーマット。
// デフォルトは横に複数表示する。
const (
    LS_LS          = 1
    LS_MULTI_COL   = 2
    LS_LONG_FORMAT = 3
)
var ls_mode = LS_MULTI_COL

// ls -l としたときに、表示される日付として、どれを表示するか
type time_type int
const (
    time_mtime time_type = iota    // default
    time_ctime                     // -c
    time_atime                     // -u
    time_btime                     // birth time

    time_num_types
)
var time_type time_type = time_mtime

// ソートする際、どれを元にソートするか
type sort_type int
const (
    sort_none sort_type = iota
    sort_name
    sort_extension
    sort_size
    sort_version
    sort_time

    sort_num_types
)
var sort_type sort_type = sort_name

// ls のファイル内容構造体
type file_info struct {
    // ファイル名
    name                string

    // リンクファイル名
    link_name           string

    // 絶対パス
    absolute_name       string

    stat                Stat

    link_mode           mode_t

    security_context    string

    stat_ok             bool

    link_ok             bool

    access_type         access_type

    has_capability      bool

    quoted              int
}


// func (self file_info) print_name_with_quoting(symlink_target bool, stack obstack, start_col size_t) {
//      var name string

     
// }

type pending struct {
    name              string
    real_name         string
    command_line_arg  bool
    next             *pending
}

var program_name string
var line_length  size_t

func initialize_main(argv []string, argc int) {
}

func set_program_name(name string) {
    program_name = strings.Replace( filepath.Base( name ), filepath.Ext( name ), "", -1 )
}


func decode_switches(argv []string, argc int) int {
    var (
        time_style_option   string
        sort_type_specified bool   = false
        kibibytes_specified bool   = false
        qmark_funny_chars   bool   = false
    )

    switch ls_mode {
    case LS_LS: 
        if is_a_tty( stdout_file_no ) {
            format            = many_per_line
            set_quoting_style( nil, shell_escape_quoting_style )
            qmark_funny_chars = true
        } else {
            format            = one_per_line
            qmark_funny_chars = false
        }
        
    case LS_MULTI_COL:
        format                = many_per_line
        set_quoting_style( nil, escape_quoting_style )

    case LS_LONG_FORMAT*
        format                = long_format
        set_quoting_style( nil, escape_quoting_style )

    default:
        abort()
    }

    time_type := time_mtime
    sort_type := sort_name

    getenv_quoting_style()

    line_length := 80

    {
        raw_columns := os.Getenv( "COLUMNS" )
        if raw_columns && !set_line_length( raw_columns ) {
            error( 0, 0, "ignoring invalid width in environment variable COLUMNS: %s", quote( raw_columns) )
        }
    }

    {
        raw_tabsize := os.Getenv( "TABSIZE" )
        tabsize      = 8
        if raw_tabsize {
            var tmp uintmax_t

            tmp, err = strconv.Atoi( raw_tabsize )
            if err != nil {
                tabsize = tmp
            } else {
                error( 0, 0, "ignoring invalid tab size in environment variable TABSIZE: %s", quote( raw_tabsize ) )
            }
        }
    }

    for {
        c, oi := getopt_long( argc, argv, "abcdfghiklmnopqrstuvw:xABCDFGHI:LNQRST:UXZ1", long_options, oi )
	
    }

    return 0
}

func main() {
    var (
        i            int
        this_pend    pending
        amount_files int
    )

    initialize_main( os.Args[1:], len( os.Args ) - 1 )
    set_program_name( os.Args[0] )

    //set_locale( LC_ALL, "" )
    //bind_text_domain( PACKAGE, LOCALEDIR )
    // text_domain( PACKAGE )

    fmt.Println( program_name )

    exit_status    := EXIT_SUCCESS
    print_dir_name := true
    // pending_dirs   := nil

    i = decode_switches( os.Args[1:], len( os.Args ) - 1  )
}