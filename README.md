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
hellojukay@local proc (master) $ sudo ./proc -p 757 -n
PROTOCOL       STATE               LOCAL          PORT      REMOTE               PORT
tcp            TCP_ESTABLISHED     10.91.28.69    35450     42.2.21.58            156
tcp            TCP_ESTABLISHED     10.91.28.69    35910     42.2.21.58            156
tcp            TCP_ESTABLISHED     10.91.28.69    47454     172.24.6.35           443
tcp            TCP_ESTABLISHED     10.91.28.69    35516     42.2.21.58            156
tcp            TCP_ESTABLISHED     10.91.28.69    33254     42.2.21.58            156
tcp            TCP_ESTABLISHED     10.91.28.69    34012     103.41.167.214        443
tcp            TCP_ESTABLISHED     10.91.28.69    35476     42.2.21.58            156
tcp6           TCP_LISTEN          ::             9090      ::                      0
tcp6           TCP_LISTEN          ::             7890      ::                      0
tcp6           TCP_LISTEN          ::             7891      ::                      0
tcp6           TCP_LISTEN          ::             7892      ::                      0
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     59248
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     43126
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     45594
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     45568
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     46028
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     45634
tcp6           TCP_ESTABLISHED     ::ffff:0:100:7f7890      ::ffff:0:100:7f     43372
tcp6           TCP_CLOSE           ::             7891      ::                      0
tcp6           TCP_CLOSE           ::             7892      ::                      0
```

# install
```shell
go get github.com/hellojukay/proc
```
