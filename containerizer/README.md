# go run containerizer.go run <cmd> <args>

Must have some bin/libs (or better an fs) in ./containerized/

Easiest way to accomplish this :  
 * mkdir ./containerized/bin/ ./containerized/lib/ ./containerized/lib64/ ./containerized/usr/
 * cp /bin /bash ./containerized/bin/bash
 * sudo mount -o bind /usr ./containerized/lib
 * sudo mount -o bind /usr ./containerized/lib64
 * sudo mount -o bind /usr ./containerized/usr
 * (sudo -E) go run containerizer.go run /bin/bash

