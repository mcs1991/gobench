**gobench**

    Use Sysbench to benchmark mysql

**how to install**

	export GOBIN=../gobench/bin
	export GOPATH=../gobench:$GOPATH
    go install gobench/src/gobench/gobench.go

**how to use**

    Usage of gobench
    Dsn:
    -u MySQL user (default root)
    -p MySQL password (default 123456)
    -P MySQL server port (default 3306)
    -H MySQL server host (default 127.0.0.1)
    -dbMySQL database name (default gobench)
    
    gobench options:
    -t number of threads to use (default 8)
    -C each test exec times (default 1)
    -i Interval for outputting statistics (default 10)
    -l define the logfile to save gobench result
    -m sysbench mode:readwrite,readonly,writeonly (default readwrite)
    -s Time interval between two benches (default 10)
    -f open gobench auto test mode
    -time  limit for total execution time in seconds(Default 1800) (default 1800)
    -pathsysbench install path
    -tbcount number of tables (default 10)
    -tbsize  data size for per table (default 100000)
    
    gobench fileio options:
    -fileio  sysbench fileio test
    -filenum number of files to create (default 128)
    -fblocksize  block size to use in all IO operations (default 16384)
    -ftotalsize  total size of files to create (default 2g)
    -ftestmode   test mode {seqwr, seqrewr, seqrd, rndrd, rndwr, rndrw}
    -fiomode file operations mode{sync,async,mmap} (default sync)
    -fflag   list of additional flags to use to open files {sync,dsync,direct}
    -ffsyncall   execute fsync every time a write operation is performed (default off)
    -ffsyncend   execute fsync at the end of the test (default on)
    -ffsyncfreq  frequency of executing fsync() (default 100)
    -frwratioRead and write ratio at test (default 1.5)

   
	
**详细使用说明见gobench_test.md**