# swiftscan
a really fast portscanner written in go

goals are:
- easy to use
- scan a few or all 65,535 ports as fast as possible
- minimal output. easy to read and send to other tools like nmap
- only features needed for identifying open ports, nothing else
- no external libraries
- cross platform compatible in windows, linux and osx

## install
just 'go build' and you're set.

## usage

scan 1024 ports
```
sws -ip 1.1.1.1
```

scan all 65535 ports
```
sws -ip 1.1.1.1 -full
```

## output
ex.
```
22,80,443,3389
```
simple, only open ports. perfect for sending over to nmap for the heavy lifting but without having to wait 10 minutes for a 65k scan.

### todo
- cidr and range support
- show filtered ports
