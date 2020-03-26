# swiftscan
a really fast portscanner written in go

objective:
- easy to use
- scan a few or all 65,535 ports as fast as possible
- minimal output. easy to read and redirect to other tools like nmap
- only features needed for identifying open/filtered ports, nothing else
- a single, static binary with no external dependencies
- cross platform compatible in windows, linux and osx

## install
just 'go build', or use the precompiled binaries.

## usage

scan first 1024 ports
```
swiftscan -ip 1.1.1.1
```
or by DNS name

```
swiftscan -ip client1

```

scan all 65535 ports
```
swiftscan -ip 1.1.1.1 -full
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
