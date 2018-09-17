package sysbench

import (
	"fmt"
	"os/exec"
	"bufio"
	"io"
	"tool"
)

func Cleanup(gvinfo map[string]interface{})(err error){
	sysbench, paramClean := tool.ParamSplice(gvinfo,"","","")
	cmd := exec.Command(sysbench,paramClean...)
	fmt.Printf("will exec cmd:%s\n",cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(line)
	}
	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return nil
}