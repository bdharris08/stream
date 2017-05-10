package main

/*
Simulate a stream of data from (CSV) to (endpoint)

1. stream data from local file to console log
2. send data from remote file to remote endpoint
3. send data from multiple files to endpoint
4. send data from multiple files to multiple endpoints
5. GUI to turn all of this on and off
*/

import (
  "fmt"
  "log"
  "os"
  "io"
  "encoding/csv"
  "bufio"
  "flag"
  "errors"
)

// log a line - later, send it somewhere
func output(line []string, target string) {
  if target == "local" {
    fmt.Println(line)
  }
}

// check for errors
func check(e error) {
  if e != nil {
    panic(e)
  }
}

// command line Flags
var target string
func init() {
    const (
      defaultTarget = "local"
      usage         = "web address to send stream to"
    )
    flag.StringVar(&target, "target", defaultTarget, usage)
    flag.StringVar(&target, "t", defaultTarget, usage+" (shorthand)")
}

func main() {
  flag.Parse()
  args := flag.Args()
  if len(args) == 0 {
    check(errors.New("Usage: $stream [-t http://target.com:port] path/to/source.csv [path/to/more.csv ...]"))
  }

  fmt.Println(args)

  f, err := os.Open(args[0])
  check(err)
  defer f.Close() //close it once everything else is done

  r := csv.NewReader(bufio.NewReader(f))

  for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatal(err)
    }

    output(record, target)
  }
}


/* // listens on 3001 and prints the rest of the path
{

  /*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  })

  log.Fatal(http.ListenAndServe(":3001", nil))
}
*/
