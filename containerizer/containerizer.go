package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	// Must run a child process to play with namespaces & cgroups
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	err_handler(cmd.Run())
}

func child() {
	fmt.Printf("Running %v \n", os.Args[2:])

	set_cgroups()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err_handler(syscall.Sethostname([]byte("containerized")))
	fmt.Println("Chrooting")
	err_handler(syscall.Chroot("./containerized"))
	fmt.Println("Chdir")
	err_handler(os.Chdir("/"))
	//err_handler(syscall.Mount("proc", "proc", "proc", 0, ""))
	//err_handler(syscall.Mount("thing", "mytemp", "tmpfs", 0, ""))
	err_handler(cmd.Run())
	//err_handler(syscall.Unmount("proc", 0))
	//err_handler(syscall.Unmount("thing", 0))
}

func set_cgroups() {
	containerizer_dir := filepath.Join("/sys/fs/cgroup/pids/", "containerizer")
	os.Mkdir(containerizer_dir, 0755)
	err_handler(ioutil.WriteFile(filepath.Join(containerizer_dir, "pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	err_handler(ioutil.WriteFile(filepath.Join(containerizer_dir, "notify_on_release"), []byte("1"), 0700))
	err_handler(ioutil.WriteFile(filepath.Join(containerizer_dir, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func err_handler(err error) {
	if err != nil {
		panic(err)
	}
}
