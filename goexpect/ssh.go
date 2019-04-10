package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	"golang.org/x/crypto/ssh"

	"google.golang.org/grpc/codes"

	"github.com/google/goexpect"
	"github.com/google/goterm/term"
)

const (
	timeout = 10 * time.Minute
)

var (
  addr  = flag.String("address", "sbx-nxos-mgmt.cisco.com:8181", "address of ssh server")
	user  = flag.String("user", "admin", "username to use")
	pass1 = flag.String("pass1", "Admin_1234!", "password to use")
	pass2 = flag.String("pass2", "pass2", "alternate password to use")
  cmd = flag.String("cmd", "sh ip route", "show routing table")

  //userRE   = regexp.MustCompile("username:")
  //passRE   = regexp.MustCompile("Password:")
  promptRE = regexp.MustCompile("sbx-n9kv-ao#")

)

func main() {
	flag.Parse()
	fmt.Println(term.Bluef("SSH to %q", *addr))

	sshClt, err := ssh.Dial("tcp", *addr, &ssh.ClientConfig{
		User:            *user,
		Auth:            []ssh.AuthMethod{ssh.Password(*pass1)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("ssh.Dial(%q) failed: %v", *addr, err)
	}
	defer sshClt.Close()

  fmt.Println("Waiting spawn...")
	e, _, err := expect.SpawnSSH(sshClt, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()


  fmt.Println("Waiting expectBatch...")
	e.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile(`sbx-n9kv-ao#`), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile(`router#`), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile(`Login: `), S: *user,
				T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong username")), Rt: 3},
			&expect.Case{R: regexp.MustCompile(`Password: `), S: *pass1, T: expect.Next(), Rt: 1},
			&expect.Case{R: regexp.MustCompile(`Password: `), S: *pass2,
				T: expect.Continue(expect.NewStatus(codes.PermissionDenied, "wrong password")), Rt: 1},
		}},
	}, timeout)

  //e.Expect(userRE, timeout)
  //e.Send(*user + "\n")
  //e.Expect(passRE, timeout)
  //e.Send(*pass1 + "\n")
  //e.Expect(promptRE, timeout)
  fmt.Println("Waiting cmd...")
  e.Send(*cmd + "\n")
  result, _, _ := e.Expect(promptRE, timeout)
  e.Send("exit\n")

  fmt.Println(term.Greenf("%s: result: %s\n", *cmd, result))



	fmt.Println(term.Greenf("All done"))
}

