// Copyright © 2018 iota-tangle.org
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iota-tangle-io/tanglifier/cmd"
	"github.com/pkg/errors"
)

func main() {
	cmd.Execute()
	var err error
	defer func() {
		if err != nil {
			// Print out a stack trace
			if err, ok := err.(stackTracer); ok {
				log.Println("Spammer has failed. Stack trace:")
				for _, f := range err.StackTrace() {
					fmt.Printf("%+s:%d", f, f)
				}
				os.Exit(1)
			} else {
				log.Println(err)
			}
		}
		// Do other stuff to shutdown
	}()
	err = cmd.RootCmd.Execute()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
