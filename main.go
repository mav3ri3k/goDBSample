package main

import (
	"fmt"
	"time"
)

func main() {
  var query queryRows 

  endtime := time.Now()
  startTime := time.Now().Add(-24 *time.Hour)
  query.queryClientTime(24, startTime, endtime)
  fmt.Println(query.rows)
}
