package main

import (
	"flag"
	"fmt"
	"os"
	"tool"
)

type Flag struct {
	thread string
	host string
	port string
	user string
	password string
	db string
	runtime int
	tablesize int
	tablecount int
	command string
	interval int
	mode string
	benchcount int
	autotest bool
	sleeptime int
	logfile string
	benchpath string
	fileio bool
	filenum int
	fblocksize int
	ftotalsize string
	ftestmode string
	fiomode string
	fflag string
	ffsyncall string
	ffsyncend string
	ffsyncfreq int
	frwratio string
	help bool
}

func (f *Flag)FlagParse(){
	thread := flag.String("t","8","number of threads to use")
	host := flag.String("H","127.0.0.1","MySQL server host")
	port := flag.String("P","3306","MySQL server port")
	user := flag.String("u","root","MySQL user")
	password := flag.String("p","123456","MySQL password")
	db := flag.String("db","gobench","MySQL database name")
	runtime := flag.Int("time",1800,"limit for total execution time in seconds(Default 1800)")
	tablesize := flag.Int("tbsize",100000,"data size for per table")
	tablecount := flag.Int("tbcount",10,"number of tables")
	command := flag.String("c","","sysbench command:prepare,run,cleanup")
	interval := flag.Int("i",10,"Interval for outputting statistics")
	mode := flag.String("m","readwrite","sysbench mode:readwrite,readonly,writeonly")
	benchcount := flag.Int("C",1,"each thread exec times")
	autotest := flag.Bool("f",false,"open gobench auto test mode")
	sleeptime := flag.Int("s",10,"Time interval between two benches")
	logfile := flag.String("l","","define the logfile to save gobench result")
	benchpath := flag.String("path","","sysbench install path")
	fileio := flag.Bool("fileio",false,"sysbench fileio test")
	filenum := flag.Int("filenum",128,"number of files to create")
	fblocksize := flag.Int("fblocksize",16384,"block size to use in all IO operations")
	ftotalsize := flag.String("ftotalsize","2g","total size of files to create")
	ftestmode := flag.String("ftestmode","","test mode {seqwr, seqrewr, seqrd, rndrd, rndwr, rndrw}")
	fiomode := flag.String("fiomode","sync","file operations mode{sync,async,mmap} (default sync)")
	fflag := flag.String("fflag","","list of additional flags to use to open files {sync,dsync,direct} (default null)")
	ffsyncall := flag.String("ffsyncall","off","Execute fsync every time a write operation is performed")
	ffsyncend := flag.String("ffsyncend","on","Execute fsync at the end of the test")
	ffsyncfreq := flag.Int("ffsyncfreq",100,"Frequency of executing fsync()")
	frwratio := flag.String("frwratio","1.5","Read and write ratio at test")
	help := flag.Bool("h",false,"show how to use gobench")

	flag.Parse()

	//NFlag返回解析时进行了设置的flag的数量
	if flag.NFlag() == 0 {
		fmt.Println("please use -h for help")
		os.Exit(1)
	}

	//打印帮助信息
	if *help == true {
		tool.HelpInfo()
	}

	f.thread = *thread
	f.host = *host
	f.port = *port
	f.user = *user
	f.password = *password
	f.db = *db
	f.runtime = *runtime
	f.tablesize = *tablesize
	f.tablecount = *tablecount
	f.command = *command
	f.interval = *interval
	f.mode = *mode
	f.benchcount = *benchcount
	f.autotest = *autotest
	f.sleeptime = *sleeptime
	f.logfile = *logfile
	f.benchpath = *benchpath
	f.fileio = *fileio
	f.filenum = *filenum
	f.fblocksize = *fblocksize
	f.ftotalsize = *ftotalsize
	f.ftestmode = *ftestmode
	f.fiomode = *fiomode
	f.fflag = *fflag
	f.ffsyncall = *ffsyncall
	f.ffsyncend = *ffsyncend
	f.ffsyncfreq = *ffsyncfreq
	f.frwratio = *frwratio
}

func GetValue()map[string]interface{}{
	gv := Flag{}
	gv.FlagParse()
	gvinfo := map[string]interface{}{
		"thread" : gv.thread,
		"host" : gv.host,
		"port" : gv.port,
		"user" : gv.user,
		"password" : gv.password,
		"db" : gv.db,
		"runtime" : gv.runtime,
		"tablesize" : gv.tablesize,
		"tablecount" : gv.tablecount,
		"command" : gv.command,
		"interval" : gv.interval,
		"mode" : gv.mode,
		"benchcount" : gv.benchcount,
		"autotest" : gv.autotest,
		"sleeptime" : gv.sleeptime,
		"logfile" : gv.logfile,
		"benchpath" : gv.benchpath,
		"fileio" : gv.fileio,
		"filenum" : gv.filenum,
		"fblocksize" : gv.fblocksize,
		"ftotalsize" : gv.ftotalsize,
		"ftestmode" : gv.ftestmode,
		"fiomode" : gv.fiomode,
		"fflag" : gv.fflag,
		"ffsyncall" : gv.ffsyncall,
		"ffsyncend" : gv.ffsyncend,
		"ffsyncfreq" : gv.ffsyncfreq,
		"frwratio" : gv.frwratio,
	}
	return gvinfo
}
