// Copyright 2014 wego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"sync"
	"io"
	"time"
	"log"
	"os"
	//"fmt"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

// Logger defines the behavior of a log provider.
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time,msg string,level int) error
	Destroy()
	Flush()
}

type WegoLogger struct {
	lock		sync.Mutex // ensures atomic writes; protects the following fields
	prefix		string     // prefix to write at beginning of each line
	flag		int        // properties
	out			io.Writer  // destination for output
	buf			[]byte     // for accumulating text to write
	wg			sync.WaitGroup
	outputs		[]*nameLogger
}

type nameLogger struct {
	Logger
	name string
}

func init()  {
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("Open log file failed：",err)
	}
	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
}
