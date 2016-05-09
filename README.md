# date2unix
Convert UNIX timestamp to date string.

# Installation
```
go get -u github.com/voidint/unix2date
```

# How to use?
- System time zone
```
unix2date 1
1970/1/1 08:00:01
```

- UTC time zone
```
unix2date --utc 1
1970/1/1 00:00:01
```


- Read UNIX timestamp form stdin
```
$ echo 12345 | unix2date --utc
1970/1/1 03:25:45

or 

$ unix2date
12345
// Enter EOF by CTRL+D(Unix/Linux) or CTRL+Z(Win)
1970/1/1 11:25:45
```



