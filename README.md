# cmdarchetype
This is an example of archetype of command line tool using [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper). The command parameters can be loaded from a file, specified inline, or a mix of the two options. 

## How it works
This is the sequence to set the parameters values:

1. If a parameter value is specified in the command line, this value will be taken
2. If a config file is specified, and it contais a value for the parameter, it will be taken
3. Otherwise, if a default value for the parameter has been set, it will be taken

There is an config file as example:
```yalm
do-async: false
source-data: https://data.safe.net:443
result-receivers: 
  - receiver1
  - receiver2
  - receiver3
db-connection-params:
  user: myuser
  password: mypassword
  db-url: https://db.safe.net:443
```

## Environment used to build this archetype
* Go: go1.11.5 linux/amd64 ()
* Make: GNU Make 4.2.1

## Execute the command
As I've said above, if a config file is specified, its parameters will be loaded. In adition, we can specify some different value for a given paraemter:
```sh
❯ go run main.go -c=./example-config.yml --do-async=true --db-url=myDB
```

With the above showed config file, this is the config loaded to execute the tool:
```sh
INFO[0000] loading config from [./example-config.yml] file 
INFO[0000] loaded config: (*config.Config)(0xc000096900)({
 SourceData: (string) (len=25) "https://data.safe.net:443",
 ResultReceivers: ([]string) (len=3 cap=3) {
  (string) (len=9) "receiver1",
  (string) (len=9) "receiver2",
  (string) (len=9) "receiver3"
 },
 DBConnectionParams: (config.DBConnectionParams) {
  User: (string) (len=6) "myuser",
  Password: (string) (len=10) "mypassword",
  DBURL: (string) (len=23) "https://db.safe.net:443"
 },
 DoAsynchronously: (bool) true
})
 
INFO[0000] starting the command at 2019-03-08 21:44:47.917868658 +0100 CET m=+0.004548910 
INFO[0001] Action done sucessfully !!!, in 1.000309761s 
 ~/go/src/github.com/theskyinflames/cmdarchetype    
```

As you can see, in the config file, the parameter *db-url* value is *https://db.safe.net:443*, but executing the tool on this way, I've forced the value for this parameter to *myDB*

