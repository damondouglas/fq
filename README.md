# fq

`fq` is a commandline utility for managing firestore instances.

# Install

```
$ go install github.com/damondouglas/fq
```

# Usage

```
NAME:
   fq - Manage Google Cloud firestore instance

USAGE:
   main [global options] command [command options] [arguments...]

COMMANDS:
   ls, list, l  List descendents of collection or document
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --project value, -p value  Google cloud associated project (defaults to gcloud config get-value project)
   --credentials value        Path to credentials (defaults to GOOGLE_APPLICATION_CREDENTIALS)
   --format value, -f value   Specify output format. (Allowed: NEWLINE_DELIMITED_JSON | CSV | FLAT) (default: "FLAT")
   --help, -h                 show help (default: false)

```

## ls

```
NAME:
   main ls - List descendents of collection or document

USAGE:
   main ls [command options] PATH

OPTIONS:
   --recursive, -r  List recursively (default: false)
   --help, -h       show help (default: false)

```
