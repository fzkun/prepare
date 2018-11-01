package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		ctx         context.Context
		cancelFucn  context.CancelFunc
		cmd         *exec.Cmd
		resultChain chan *result
		res         *result
	)

	resultChain = make(chan *result, 1000)

	ctx, cancelFucn = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2; echo hello;")

		output, err = cmd.CombinedOutput()

		resultChain <- &result{
			output: output,
			err:    err,
		}

	}()

	time.Sleep(1 * time.Second)
	cancelFucn()

	res = <-resultChain
	fmt.Println(res.err, string(res.output))

}
