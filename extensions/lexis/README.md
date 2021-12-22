# LEXIS EXTENSIONs SERVICE

Cyclops Engine's Extension Service for LEXIS implemented in Go

## How to build

The building of the service is carried by a multistage docker build, we build everything in an image containing everything needed to build Golang applications and then the executable is moved to a new image that contains the bare minimum starting from the scratch image.

Within the folder build at the root of the repo there's a script call start.sh, it's invokation admits "Dev" as parameter. Running the script with the parameter will use the local version in the repository as the base for the building of the service's docker image, however doing it without providing it will make the building of the service taking everything from sources.

```
./start.sh [Dev]
```

The initial branch used when building from sources is "master", it can be changed with a few other parameters by editing the script.

Using the [Dev] optional argument will take the code present in the repo at the moment of invokation.

## How to run

Within the folder run at the root of the repo there's a docker-compose sample file and a config.toml sample file. Once configured with appropriate data to start the service just issue the following command:

```
docker-compose up -d
```


