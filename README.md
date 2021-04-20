# Azr Replace

Azure CI/CD replace cli tool

## Why the REPO

working with Azure CI/CD, replace format "#{}#" is different from others, so make this tool to generate file with args/envs  for run helm or kubectl.

## Build

install golang 1.6+, then run

```shell
go build . 

```

or download binary from release

## Run

check details by

> ./azr-replace -h


```shell



NAME:
   Azr Replace - -

USAGE:
   azr-replace [global options] command [command options] [arguments...]

DESCRIPTION:
   cli to run Azure DevOps replace locally

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --env value, -e value     env settings
   --file value, -f value    (default: ".env")
   --source value, -s value
   --output value, -o value  output to file or default - to STDOUT (default: "-")
   --log value, -l value     (default: "info") [$LOG_LEVEL]
   --start value             (default: "#{")
   --end value               (default: "}#")
   --help, -h                show help (default: false)
```

such as run:

```shell
./azr-replace -s test.txt -e v3="values from -e" -o replaced.txt
```
