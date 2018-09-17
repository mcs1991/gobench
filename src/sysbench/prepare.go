package sysbench

import (
	"os/exec"
	"fmt"
	"bufio"
	"io"
	"tool"
)

func Prepare(gvinfo map[string]interface{})(err error) {
	sysbench, paramPre := tool.ParamSplice(gvinfo,"","","")
	cmd := exec.Command(sysbench,paramPre...)
	fmt.Printf("will exec cmd:%s\n",cmd.Args)
	/*等待执行完成一次性打印所有输出
	output, err := cmd.Output()
	fmt.Printf("output:%s\n",output)
	fmt.Printf("err:%s\n",err)
	*/
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	cmd.Start()
	//方法一：
	/*ioutil
	content , err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	cmd.Wait()
	fmt.Println(string(content))
	*/
	//方法二:
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取

	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err{
			break
		}

		fmt.Printf(string(line))
	}

	/*
	//方法三：
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("%s\n", text)
	}
	*/
	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return
}