package sysbench

import (
	"fmt"
	"os/exec"
	"bufio"
	"io"
	"time"
	"os"
	"strings"
	"tool"
)


func RunDB(gvinfo map[string]interface{})(err error){
	var th []string
	var mo []string
	//获取日志信息
	f, err := tool.Addlog(gvinfo)
	if err != nil {
		fmt.Printf("create logfile failed!\n")
		os.Exit(-1)
	}
	//判断执行自动化测试还是单次测试
	if gvinfo["autotest"] == true {
		//解析thread字符串
		t := strings.Split(gvinfo["thread"].(string),",")
		for i := range t {
			th = append(th, t[i])
		}
		//解析mode字符串
		m := strings.Split(gvinfo["mode"].(string),",")
		for i := range m {
			mo = append(mo,m[i])
		}
		//每一个thread的每一种模式均会执行count次
		for thcount := 0;thcount < len(th);thcount++ {
			for mcount :=0; mcount < len(mo); mcount++ {
				for count := 0; count < gvinfo["benchcount"].(int); count++ {
					//fmt.Printf("th[%d]:%s\n",count,th[thcount])
					//fmt.Printf("mo[%d]:%s\n",count,mo[mcount])
					//fmt.Printf("benchcount:%d\n",count)
					err := RunCommand(gvinfo, th[thcount], mo[mcount],"", f)
					if err != nil {
						fmt.Println(err.Error())
						return err
					}
					fmt.Printf("-------------------sleep %d seconds-------------------\n", gvinfo["sleeptime"])
					f.WriteString(fmt.Sprintf("-------------------sleep %d seconds-------------------\n", gvinfo["sleeptime"]))
					time.Sleep(time.Duration(gvinfo["sleeptime"].(int)) * time.Second)
				}
			}
		}
	}else {
		err := RunCommand(gvinfo,gvinfo["thread"].(string),gvinfo["mode"].(string),"",f)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func RunFile(gvinfo map[string]interface{})(err error){
	var fmo []string
	var th []string
	//获取日志信息
	f, err := tool.Addlog(gvinfo)
	if err != nil {
		fmt.Printf("create logfile failed!\n")
		os.Exit(-1)
	}
	//判断执行自动化测试还是单次测试
	if gvinfo["autotest"] == true {
		m := strings.Split(gvinfo["ftestmode"].(string),",")
		for i := range m {
			fmo = append(fmo,m[i])
		}
		t := strings.Split(gvinfo["thread"].(string),",")
		for i := range t {
			th = append(th, t[i])
		}
		for tcount :=0;tcount < len(th);tcount ++ {
			for fcount := 0; fcount < len(fmo); fcount ++ {
				for count := 0; count < gvinfo["benchcount"].(int); count++ {
					err := RunCommand(gvinfo, th[tcount],"", fmo[fcount],f)
					if err != nil {
						fmt.Println(err.Error())
						return err
					}
					fmt.Printf("-------------------sleep %d seconds-------------------\n", gvinfo["sleeptime"])
					f.WriteString(fmt.Sprintf("-------------------sleep %d seconds-------------------\n", gvinfo["sleeptime"]))
					time.Sleep(time.Duration(gvinfo["sleeptime"].(int)) * time.Second)
				}
			}
		}
	}else {
		err := RunCommand(gvinfo, gvinfo["thread"].(string),"", gvinfo["ftestmode"].(string),f)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

//调用sysbench run函数
func RunCommand(gvinfo map[string]interface{},th string,mo string,fmo string,f *os.File)(err error){
	sysbench, paramRun := tool.ParamSplice(gvinfo,th,mo,fmo)
	cmd := exec.Command(sysbench,paramRun...)
	fmt.Printf("will exec cmd:%s\n",cmd.Args)
	t := time.Now()
	f.WriteString(fmt.Sprintf("gobench start:%s\n",t.Format("2006-01-02_15:04:05")))
	f.WriteString(fmt.Sprintf("gobench thread:%s\n",th))
	f.WriteString(fmt.Sprintf("gobench mode:%s\n",mo))
	f.WriteString(fmt.Sprintf("gobench mode:%s\n",fmo))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	cmd.Start()
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Printf("%s",line)
		f.WriteString(string(line))
	}
	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return nil
}