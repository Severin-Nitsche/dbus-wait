package main

import (
  "os"
  "os/exec"
  "flag"
  "log"

  dbus "github.com/godbus/dbus/v5"
)

func main() {
  log.SetFlags(log.Lshortfile)

  session := flag.Bool(
    "session", 
    false, 
    "Connect to the session bus instead of the system bus",
  )

  flag.Parse()

  var before *exec.Cmd
  matchOptions := make(
    []dbus.MatchOption,
    0,
    flag.NArg() / 2,
  )
  for i := 0; i+1 < flag.NArg(); i += 2 {
    key := flag.Arg(i)
    value := flag.Arg(i+1)

    if key == "--" {
      before = exec.Command(value, flag.Args()[i+2:]...)
      break;
    }

    matchOptions = append(
      matchOptions,
      dbus.WithMatchOption(key, value),
    )
  }

  var connect func(...dbus.ConnOption) (*dbus.Conn, error)
  if *session {
    connect = dbus.ConnectSessionBus
  } else {
    connect = dbus.ConnectSystemBus
  }

  conn, err := connect()
  if err != nil {
    log.Fatalf(
      "Could not connect to the (session: %v) D-Bus: %v",
      *session,
      err,
    )
  }

  err = conn.AddMatchSignal(matchOptions...)
  if err != nil {
    log.Fatalln("Failed to listen for signals:", err)
  }


  c := make(chan *dbus.Signal, 10)
  conn.Signal(c)
  log.Println("Waiting for signal...")

  if before != nil {
    before.Stdout = os.Stdout;
    err = before.Run()
    if err != nil {
      log.Fatalln("Error starting command:", err)
    }
  }

  _ = <-c
}
