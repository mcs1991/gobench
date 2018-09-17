package main

import(
	"fmt"
	"sysbench"
	"os"
	"tool"
)

func main(){
	//get flag
	gvinfo := GetValue()
	//fmt.Printf("command:gvinfo:%s\n",gvinfo)
	//check whether the sysbench is exist
	checkFirst,isPath,err := tool.Check(gvinfo)
	if isPath == false && err != nil {
		fmt.Printf("check failed:\n%s\n",err)
		os.Exit(-1)
	}
	fmt.Printf("check ok:\n%s\n",checkFirst)
	//check flag
	tool.CheckFlag(gvinfo)
	//exec sysbench prepare
	if gvinfo["command"] == "prepare" {
		if gvinfo["fileio"] == false {
			if err := tool.CreateDB(gvinfo); err != nil {
				fmt.Printf("create database error!\n")
				os.Exit(-1)
			}
			if err := sysbench.Prepare(gvinfo); err != nil {
				fmt.Printf("sysbench prepare error!\n")
				os.Exit(-1)
			}
		}else {
			if err := sysbench.Prepare(gvinfo); err != nil {
				fmt.Printf("sysbench fileio prepare error!\n")
				os.Exit(-1)
			}
		}
	}
	//exec sysbench run
	if gvinfo["command"] == "run" {
		if gvinfo ["fileio"] == false {
			if err := sysbench.RunDB(gvinfo); err != nil {
				fmt.Printf("sysbench run error!")
				os.Exit(-1)
			}
		}else {
			if err := sysbench.RunFile(gvinfo); err != nil {
				fmt.Printf("sysbench fileio run error!")
				os.Exit(-1)
			}
		}
	}
	//exec sysbench cleanup
	if gvinfo["command"] == "cleanup"{
		if gvinfo["fileio"] == false {
			if err := sysbench.Cleanup(gvinfo); err != nil {
				fmt.Printf("sysbench cleanup error!")
				os.Exit(-1)
			}
		}else {
			if err := sysbench.Cleanup(gvinfo); err != nil {
				fmt.Printf("sysbench cleanup error!")
				os.Exit(-1)
			}
		}
	}
}
