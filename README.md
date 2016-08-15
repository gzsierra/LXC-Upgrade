# LXC-Upgrade
Go script that update all active container

## Requirements
* Admin access (VERY IMPORTANT)
* [Go-LXC](https://github.com/lxc/go-lxc/tree/v2)
* [Pb.v1](https://github.com/cheggaaa/pb/tree/v1.0.5)
* [Golang 1.x](https://golang.org/dl/)
* [LXC](https://github.com/lxc/lxc/releases)

## Installation
There is multiple way to get to it.

### Method #1
Install requirements manually
```
apt-get install -y pkg-config lxc-dev
go get gopkg.in/lxc/go-lxc.v2
go get gopkg.in/cheggaaa/pb.v1
```
Get the project
```
git clone https://github.com/gzsierra/LXC-Upgrade.git
```
or
```
go get github.com/gzsierra/LXC-Upgrade
```

### Method #2
Using [godep](https://github.com/tools/godep)
Get the project
```
git clone https://github.com/gzsierra/LXC-Upgrade.git
```
or
```
go get github.com/gzsierra/LXC-Upgrade
```
In the project folder
```
godep restore
```

## Usage
The script does everything for you.
```
go run upgrade.go
```
or
```
go build upgrade.go
./upgrade
```

# LICENCE
MIT
