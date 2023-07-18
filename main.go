package main

import (
	"time"
)

func main() {
  var query queryRows 
  query.connect()
 
  // time interval based query
  endtime := time.Now()
  startTime := time.Now().Add(-24 *time.Hour)
  query.queryClientTime(24, startTime, endtime)
  query.queryPrint()
  query.queryEmpty()

  //insert
  query.insert(24, 8)

  //client_id based query
  query.queryClient(24)
  query.queryPrint()
  query.queryEmpty()
  
  query.queryClose()
}
