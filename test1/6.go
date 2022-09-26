package test1

import (
	"fmt"
	"os"
)

func ninenine() {
	file, err := os.Create("ninenine.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// 打印9行
	for i := 1; i <= 9; i++ {
		for j := 1; j < i; j ++ {
			_, err := fmt.Fprintf(file, "%dx%d=%d ", j, i, j*i)
			if err != nil {
				return 
			}
		}
		_, err := fmt.Fprintf(file, "%dx%d=%d\n", i, i, i*i)
		if err != nil {
			return 
		}
	}

}