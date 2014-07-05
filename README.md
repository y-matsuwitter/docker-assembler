docker-assembler
====

docker-assembler is a tool for combine multiple Dockerfiles.

## Quickstart
### Create your project

```
$ docker-assembler create <your project name>
```

### Add your sub Dockerfiles

```
$ cd <your project dir>
$ docker-assembler add <sub dockerfile name>
```

This adds sub Dockerfiles in a directory you specified.

### Build Docker image with sub Dockerfiles

```
$ cd <your project dir>
$ docker-assembler build <image name> <sub dockerfile names>...
```

This will build docker image with dockerfiles you specified.

## Usage

```
$ docker-assembler help                                                                                                                                                                                                                                            [git master]
NAME:
   docker-assembler - Docker image builder with sub-Dockerfiles

USAGE:
   docker-assembler [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   create       Setup docker assembler project.
   add          Add new sub-Dockerfile and its directory.
   build        Build a Docker image with specified sub-Dockerfiles.
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --version, -v        print the version
   --help, -h           show help

```

## Sub Dockerfile

`docker-assembler` builds each sub dockerfile in the directory of each sub dockerfile name.
If you have your Dockerfiles with some additional files for `ADD`, you can put all files into sub directory of your project.

## Author

[y-matsuwitter](http://github.com/y-matsuwitter)
