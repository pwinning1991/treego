## Simple tool similar to tree written in go


### Just run the commmand and the directory you want listed
- It will default to the direcotry you are in without any agruments 
```bash
treego 
.
├──Makefile
├──README.md
├──go.mod
├──main.go
└──treego
```

```bash
treego ../advent_of_code 
advent_of_code
├──2018
│  ├──go
│  │  ├──day1
│  │  │  ├──input.txt
│  │  │  └──main.go
│  │  ├──day2
│  │  │  ├──p1
│  │  │  │  ├──input.txt
│  │  │  │  └──main.go
│  │  │  └──p2
│  │  │     ├──input.txt
│  │  │     └──main.go
│  │  ├──day3
│  │  │  ├──p1
│  │  │  │  ├──input.txt
│  │  │  │  └──main.go
│  │  │  └──p2
│  │  │     ├──input.txt
│  │  │     └──main.go
│  │  └──day4
│  │     ├──input.txt
│  │     └──main.go
│  └──python
│     ├──day1
│     │  ├──day.py
│     │  └──input.txt
│     ├──day2
│     │  ├──day2.py
│     │  └──input.txt
│     └──day3
│        ├──day3.py
│        └──input.txt
└──2019
   └──python
      ├──day1
      │  ├──day1.py
      │  ├──input.txt
      │  └──test_day1.py
      └──day2
         ├──day2.py
         └──input.txt
```
### TODO
1. add support for symlinks
2. add count of files and directories to and of output command
3. Add Support for hidden files/directories

### Releases
Found in the releases tab
