package tool

import (
	"fmt"
	"os"
	"time"
	"strings"
	"os/exec"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//create gobench database
func CreateDB(gvinfo map[string]interface{})(err error){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema",gvinfo["user"],gvinfo["password"],gvinfo["host"],gvinfo["port"])
	fmt.Printf("dsn:%s\n",dsn)
	conn ,err := sql.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer conn.Close()
	createCmd := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;",gvinfo["db"])
	_, err = conn.Exec(createCmd)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func ParamSplice(gvinfo map[string]interface{},th string,mo string,fmo string)(bench string,param []string){
	var sysbench string
	var readwrite string
	var readonly string
	var writeonly string
	if gvinfo["benchpath"] == "" {
		sysbench = "/usr/local/bin/sysbench"
		readwrite = "/usr/local/share/sysbench/oltp_read_write.lua"
		readonly = "/usr/local/share/sysbench/oltp_read_only.lua"
		writeonly = "/usr/local/share/sysbench/oltp_write_only.lua"
	}else {
		sysbench = gvinfo["benchpath"].(string) + "/bin/sysbench"
		readwrite = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_read_write.lua"
		readonly = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_read_only.lua"
		writeonly = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_write_only.lua"
	}
	if gvinfo["fileio"] == false {
		if gvinfo["command"] == "prepare" {
			paramPre := []string{"--db-driver=mysql"}
			paramPre = append(paramPre, fmt.Sprintf("--threads=%s", gvinfo["thread"]))
			paramPre = append(paramPre, fmt.Sprintf("--mysql-host=%s", gvinfo["host"]))
			paramPre = append(paramPre, fmt.Sprintf("--mysql-port=%s", gvinfo["port"]))
			paramPre = append(paramPre, fmt.Sprintf("--mysql-user=%s", gvinfo["user"]))
			paramPre = append(paramPre, fmt.Sprintf("--mysql-password=%s", gvinfo["password"]))
			paramPre = append(paramPre, fmt.Sprintf("--mysql-db=%s", gvinfo["db"]))
			paramPre = append(paramPre, fmt.Sprintf("--table-size=%d", gvinfo["tablesize"]))
			paramPre = append(paramPre, fmt.Sprintf("--tables=%d", gvinfo["tablecount"]))
			switch gvinfo["mode"] {
			case "readwrite":
				paramPre = append(paramPre, readwrite)
			case "readonly":
				paramPre = append(paramPre, readonly)
			case "writeonly":
				paramPre = append(paramPre, writeonly)
			default:
				paramPre = append(paramPre, readwrite)
			}
			paramPre = append(paramPre,"prepare")
			return sysbench, paramPre
		}
		if gvinfo["command"] == "run" {
			paramRun := []string{"--db-driver=mysql"}
			paramRun = append(paramRun, fmt.Sprintf("--threads=%s", th))
			paramRun = append(paramRun, fmt.Sprintf("--mysql-host=%s", gvinfo["host"]))
			paramRun = append(paramRun, fmt.Sprintf("--mysql-port=%s", gvinfo["port"]))
			paramRun = append(paramRun, fmt.Sprintf("--mysql-user=%s", gvinfo["user"]))
			paramRun = append(paramRun, fmt.Sprintf("--mysql-password=%s", gvinfo["password"]))
			paramRun = append(paramRun, fmt.Sprintf("--mysql-db=%s", gvinfo["db"]))
			paramRun = append(paramRun, fmt.Sprintf("--table-size=%d", gvinfo["tablesize"]))
			paramRun = append(paramRun, fmt.Sprintf("--tables=%d", gvinfo["tablecount"]))
			paramRun = append(paramRun, fmt.Sprintf("--time=%d", gvinfo["runtime"]))
			paramRun = append(paramRun, fmt.Sprintf("--report-interval=%d", gvinfo["interval"]))
			switch mo {
			case "readwrite":
				paramRun = append(paramRun, readwrite)
			case "readonly":
				paramRun = append(paramRun, readonly)
			case "writeonly":
				paramRun = append(paramRun, writeonly)
			default:
				paramRun = append(paramRun, readwrite)
			}
			paramRun = append(paramRun,"run")
			return sysbench, paramRun
		}
		if gvinfo["command"] == "cleanup" {
			paramClean := []string{"--db-driver=mysql"}
			paramClean = append(paramClean, fmt.Sprintf("--threads=%s", gvinfo["thread"]))
			paramClean = append(paramClean, fmt.Sprintf("--mysql-host=%s", gvinfo["host"]))
			paramClean = append(paramClean, fmt.Sprintf("--mysql-port=%s", gvinfo["port"]))
			paramClean = append(paramClean, fmt.Sprintf("--mysql-user=%s", gvinfo["user"]))
			paramClean = append(paramClean, fmt.Sprintf("--mysql-password=%s", gvinfo["password"]))
			paramClean = append(paramClean, fmt.Sprintf("--mysql-db=%s", gvinfo["db"]))
			paramClean = append(paramClean, fmt.Sprintf("--table-size=%d", gvinfo["tablesize"]))
			paramClean = append(paramClean, fmt.Sprintf("--tables=%d", gvinfo["tablecount"]))
			switch gvinfo["mode"] {
			case "readwrite":
				paramClean = append(paramClean, readwrite)
			case "readonly":
				paramClean = append(paramClean, readonly)
			case "writeonly":
				paramClean = append(paramClean, writeonly)
			default:
				paramClean = append(paramClean, readwrite)
			}
			paramClean = append(paramClean,"cleanup")
			return sysbench, paramClean
		}
	}else {
		if gvinfo["command"] == "prepare" {
			paramFilePre := []string{"fileio"}
			paramFilePre = append(paramFilePre,fmt.Sprintf("--threads=%s",gvinfo["thread"]))
			paramFilePre = append(paramFilePre,fmt.Sprintf("--file-total-size=%s",gvinfo["ftotalsize"]))
			paramFilePre = append(paramFilePre,fmt.Sprintf("--file-num=%d",gvinfo["filenum"]))
			paramFilePre = append(paramFilePre,fmt.Sprintf("--file-block-size=%d",gvinfo["fblocksize"]))
			paramFilePre = append(paramFilePre,fmt.Sprintf("prepare"))
			return sysbench,paramFilePre
		}
		if gvinfo["command"] == "run" {
			paramFileRun := []string{"fileio"}
			paramFileRun = append(paramFileRun,fmt.Sprintf("--threads=%s",th))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-total-size=%s",gvinfo["ftotalsize"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-num=%d",gvinfo["filenum"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-block-size=%d",gvinfo["fblocksize"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-test-mode=%s",fmo))
			if gvinfo["fflag"] != "" {
				paramFileRun = append(paramFileRun,fmt.Sprintf("--file-extra-flags=%s",gvinfo["fflag"]))
			}
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-io-mode=%s",gvinfo["fiomode"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-fsync-all=%s",gvinfo["ffsyncall"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-fsync-end=%s",gvinfo["ffsyncend"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-fsync-freq=%d",gvinfo["ffsyncfreq"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--file-rw-ratio=%s",gvinfo["frwratio"]))
			paramFileRun = append(paramFileRun,fmt.Sprintf("--time=%d",gvinfo["runtime"]))
			paramFileRun = append(paramFileRun,"run")
			return sysbench,paramFileRun
		}
		if gvinfo["command"] == "cleanup" {
			paramFileClean := []string{"fileio"}
			paramFileClean = append(paramFileClean,fmt.Sprintf("--threads=%s",gvinfo["thread"]))
			paramFileClean = append(paramFileClean,fmt.Sprintf("--file-total-size=%s",gvinfo["ftotalsize"]))
			paramFileClean = append(paramFileClean,fmt.Sprintf("--file-num=%d",gvinfo["filenum"]))
			paramFileClean = append(paramFileClean,fmt.Sprintf("--file-block-size=%d",gvinfo["fblocksize"]))
			paramFileClean = append(paramFileClean,fmt.Sprintf("cleanup"))
			return sysbench,paramFileClean
		}
	}
	return "",nil
}

func Addlog(gvinfo map[string]interface{})(f *os.File,err error){
	var logfile string
	if gvinfo["logfile"] != "" {
		logfile = gvinfo["logfile"].(string)
	}else {
		t := time.Now()
		logfile = fmt.Sprintf("/tmp/gobench_%s.log",t.Format("2006-01-02_15:04:05"))

	}
	f, err = os.Create(logfile)
	return f,err
}

func Check(gvinfo map[string]interface{})(string,bool, error){
	var sysbench string
	var readwrite string
	var readonly string
	var writeonly string
	if gvinfo["benchpath"] == "" {
		sysbench = "/usr/local/bin/sysbench"
		readwrite = "/usr/local/share/sysbench/oltp_read_write.lua"
		readonly = "/usr/local/share/sysbench/oltp_read_only.lua"
		writeonly = "/usr/local/share/sysbench/oltp_write_only.lua"
		if isPath, err := PathExists(sysbench); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(readwrite); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(readonly); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(writeonly); isPath == false && err != nil {
			return "", isPath, err
		}
	}else {
		sysbench = gvinfo["benchpath"].(string) + "/bin/sysbench"
		readwrite = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_read_write.lua"
		readonly = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_read_only.lua"
		writeonly = gvinfo["benchpath"].(string) + "/share/sysbench/oltp_write_only.lua"
		if isPath, err := PathExists(sysbench); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(readwrite); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(readonly); isPath == false && err != nil {
			return "", isPath, err
		}
		if isPath, err := PathExists(writeonly); isPath == false && err != nil {
			return "", isPath, err
		}
	}
	cmd := exec.Command(sysbench,"--version")
	output, err := cmd.Output()
	if err != nil {
		return "", false, err
	}
	return string(output), true, nil
}

func CheckFlag(gvinfo map[string]interface{}){
	if gvinfo["command"] == "" {
		fmt.Printf("-c parameter value is not defined\n")
		os.Exit(-1)
	}
	if gvinfo["command"] != "run" && gvinfo["command"] != "prepare" && gvinfo["command"] != "cleanup" {
		fmt.Printf("-c parameter value is incorrect!\n")
		os.Exit(-1)
	}
	if gvinfo["command"] != "run" && gvinfo["logfile"] != "" {
		fmt.Printf("You can use -l parameter only in run mode!\n")
		os.Exit(-1)
	}
	if gvinfo["command"] != "run" && gvinfo["autotest"] != false {
		fmt.Printf("You can use -f parameter only in run mode!\n")
		os.Exit(-1)
	}
	if gvinfo["command"] != "run" && gvinfo["benchcount"] != 1 {
		fmt.Printf("You can use -C parameter only in run mode!\n")
		os.Exit(-1)
	}

	if gvinfo["autotest"] == false && (strings.Contains(gvinfo["thread"].(string),",") == true || strings.Contains(gvinfo["mode"].(string),",") == true || strings.Contains(gvinfo["ftestmode"].(string),",") == true) {
		fmt.Printf("You can not enter more than one thread or mode or ftestmode when -f is not defined!\n")
		os.Exit(-1)
	}
	if gvinfo["sleeptime"].(int) < 0 || gvinfo["benchcount"].(int) < 0 || gvinfo["runtime"].(int) < 0 || gvinfo["tablesize"].(int) < 0 || gvinfo["tablecount"].(int) < 0 {
		fmt.Printf("You have entered the wrong value!\n")
		fmt.Printf("Please check the parameter!\n")
		os.Exit(-1)
	}
	if gvinfo["fileio"] == false && gvinfo["ftestmode"] != "" {
		fmt.Printf("You can not test file io when fileio parameter is not defined!\n")
		os.Exit(-1)
	}
	if gvinfo["fileio"] == true && gvinfo["ftestmode"] == "" {
		fmt.Printf("Missing required argument: --file-test-mode!\n")
		os.Exit(-1)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

func HelpInfo(){
	fmt.Printf("Usage of gobench\n")
	fmt.Printf("Dsn:\n")
	fmt.Printf("-u     MySQL user (default root)\n")
	fmt.Printf("-p     MySQL password (default 123456)\n")
	fmt.Printf("-P     MySQL server port (default 3306)\n")
	fmt.Printf("-H     MySQL server host (default 127.0.0.1)\n")
	fmt.Printf("-db    MySQL database name (default gobench)\n")
	fmt.Printf("\ngobench options:\n")
	fmt.Printf("-t     number of threads to use (default 8)\n")
	fmt.Printf("-C     each test exec times (default 1)\n")
	fmt.Printf("-i     Interval for outputting statistics (default 10)\n")
	fmt.Printf("-l     define the logfile to save gobench result\n")
	fmt.Printf("-m     sysbench mode:readwrite,readonly,writeonly (default readwrite)\n")
	fmt.Printf("-s     Time interval between two benches (default 10)\n")
	fmt.Printf("-f     open gobench auto test mode\n")
	fmt.Printf("-time  limit for total execution time in seconds(Default 1800) (default 1800)\n")
	fmt.Printf("-path        sysbench install path\n")
	fmt.Printf("-tbcount     number of tables (default 10)\n")
	fmt.Printf("-tbsize      data size for per table (default 100000)\n")
	fmt.Printf("\ngobench fileio options:\n")
	fmt.Printf("-fileio      sysbench fileio test\n")
	fmt.Printf("-filenum     number of files to create (default 128)\n")
	fmt.Printf("-fblocksize  block size to use in all IO operations (default 16384)\n")
	fmt.Printf("-ftotalsize  total size of files to create (default 2g)\n")
	fmt.Printf("-ftestmode   test mode {seqwr, seqrewr, seqrd, rndrd, rndwr, rndrw}\n")
	fmt.Printf("-fiomode     file operations mode{sync,async,mmap} (default sync)\n")
	fmt.Printf("-fflag       list of additional flags to use to open files {sync,dsync,direct}\n")
	fmt.Printf("-ffsyncall   execute fsync every time a write operation is performed (default off)\n")
	fmt.Printf("-ffsyncend   execute fsync at the end of the test (default on)\n")
	fmt.Printf("-ffsyncfreq  frequency of executing fsync() (default 100)\n")
	fmt.Printf("-frwratio    Read and write ratio at test (default 1.5)\n")
	os.Exit(1)
}
