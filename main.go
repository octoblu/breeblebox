package main

import (
  "log"
  "os"
  "github.com/codegangsta/cli"
  "github.com/garyburd/redigo/redis"
)

func main() {
  app := cli.NewApp()
  app.Name = "redis-interval-work-queue"
  app.Usage = "Run the work queue"
  app.Action = doNothing
  app.Run(os.Args)
}

func doNothing(context *cli.Context) {
  redisConn, err := redis.Dial("tcp", ":6379")
  if err != nil {
    log.Fatalf("Redis Connect error: %v", err.Error())
  }
  defer redisConn.Close()

  for {
    result,err := redisConn.Do("BRPOP", "linear-job-queue", 500)
    if err != nil {
      log.Fatalf("Redis BRPOP error: %v", err.Error())
    }
    results := result.([]interface{})
    jobStr := string(results[1].([]uint8))
    log.Printf("Got a job %v", jobStr)
  }
}
