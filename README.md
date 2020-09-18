# proc
display linux process information

## show command line
```shell
hellojukay@local proc (master) $ sudo ./proc -p 765 -c
cmdline     /bin/offlinenotepad -db /opt/offlinenotepad/notepad.db
exe         /usr/bin/offlinenotepad
workspace   /opt/offlinenotepad
```
## show environment
```shell
hellojukay@local proc (master) $ sudo ./proc -p 765 -e
LANG                en_US.UTF-8
LC_ADDRESS          zh_CN.UTF-8
LC_MONETARY         zh_CN.UTF-8
LC_IDENTIFICATION   zh_CN.UTF-8
LC_MEASUREMENT      zh_CN.UTF-8
LC_NAME             zh_CN.UTF-8
LC_NUMERIC          zh_CN.UTF-8
LC_PAPER            zh_CN.UTF-8
LC_TELEPHONE        zh_CN.UTF-8
LC_TIME             zh_CN.UTF-8
PATH                /usr/local/sbin:/usr/local/bin:/usr/bin:/var/lib/snapd/snap/bin
INVOCATION_ID       9555009bb630491f98ba3cb78066624a
JOURNAL_STREAM      8:24929
```
## show open files
```shell
hellojukay@local proc (master) $ sudo ./proc -p 765 -f
0    /dev/null
1    socket:[24929]
2    socket:[24929]
3    /opt/offlinenotepad/notepad.db
4    anon_inode:[eventpoll]
5    pipe:[23218]
6    pipe:[23218]
7    socket:[23223]
```

## show network 
```shell
hellojukay@local proc (master) $ sudo ./proc -p 765 -n
```

# install
```shell
go get github.com/hellojukay/proc
```
