package main

import "log"

func main() {
	// if normal
	x := 1
	if x > 0 {
	}

	// if has init
	if err := file.Chmod(0664); err != nil { // ; 前初始化 err 后续判断 err 是否为 nil
		log.Print(err) // 如果不为空 输出 err
		return err     // 保存 当前err 变量
	}
}
