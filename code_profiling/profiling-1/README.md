
[How to profile Go with pprof in 30 seconds](https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a)

$ brew
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
$ brew install graphviz

$ mkdir install_pprof
$ cd install_pprof
$ got mod init install_pprof
$ go get -u github.com/google/pprof

$ go tool pprof http://localhost:6060/debug/pprof/heap
$ go tool pprof http://localhost:6060/debug/pprof/goroutine
$ go tool pprof http://localhost:6060/debug/pprof/block
$ go tool pprof http://localhost:6060/debug/pprof/mutex
