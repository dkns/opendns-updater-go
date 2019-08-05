```
> opendns-updater-go
NAME:
   odns-updater - Update OpenDNS ip

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --username value, -u value  Username used for authentication
   --password value, -p value  Password used for authentication
   --network value, -n value   Name of network you want to update
   --help, -h                  show help
   --version, -v               print the version
```

# How to update automatically
Download binary and do something like this in cron:

    */10 * * * * /home/foo/opendns-updater-go -u <username> -p <password> -n <network>
