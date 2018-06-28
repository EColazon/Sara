package handleShared

import (
	//"io"
	"fmt"
	"bufio"
	"os/exec"
)

func ExecShell(linuxCmd string) {

	cmd := exec.Command("/bin/bash", "-c", linuxCmd)
	fmt.Println("---> Shell:", cmd.Args[0])

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("---> Error: Can Not Obtain Stdout Pipe For CMD: %s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("---> Error: The CMD is err: ", err)
		return
	}
	//使用带缓冲的读取器
	outputBuff := bufio.NewReader(stdout)

	for {
		//一次读一行,获取当前行是否被读完
		
		output, _, err := outputBuff.ReadLine()
		if err != nil {
			//判断是否到文件的结尾,否则出粗啦
			if err.Error() != "EOF" {
				fmt.Printf("---> Error: EOF %s\n", err)
			}
			return
		}
		fmt.Println(string(output))
		
		
		/*
		line, err := outputBuff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Printf(line)
		*/
	}
	//cmd.Wait()
	/*
	//wait方法会一直阻塞到其所属的命令完全运行结束为止
	if err := cmd.Wait(); err != nil {
		fmt.Println("---> Wait: ", err.Error())
		return
	}
	*/
}

