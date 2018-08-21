<meta charset="UTF-8">
<head>
  <link rel="stylesheet" type="text/css" href="wiki_assets/stylesheet.css">
</head> 

![logo](wiki_assets/pics/misc/mphclub_logo.png)

# MPH Development Environment Setup

<!-- Note: table of contents create in VS Code using the 'Markdown TOC' Extension

<!-- TOC -->

- [MPH Development Environment Setup](#mph-development-environment-setup)
    - [Tools](#tools)
    - [Docker Base Images](#docker-base-images)
    - [Required Installs for MPH](#required-installs-for-mph)
    - [Recommended Installs](#recommended-installs)
        - [Docker Set Up](#docker-set-up)
        - [Postman Set Up](#postman-set-up)
        - [Running MPH](#running-mph)
        - [Misc. Info](#misc-info)
            - [Useful Docker Commands](#useful-docker-commands)


<!-- /TOC -->


## Tools
The following tools are used by the MPH:
- [Docker](https://docs.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
## Docker Base Images
The following Docker base images can used for MPH environment setup.
- [PostgreSQL](https://hub.docker.com/_/postgres/)
- [Tomcat](https://hub.docker.com/_/tomcat/)


## Required Installs for MPH
- **Main IDE**: [NetBeans](https://netbeans.org/downloads/index.html) or [Eclipse](https://www.eclipse.org/downloads/) recommended.
- **[Docker](https://docs.docker.com/engine/installation/)**
- **[SBT](http://www.scala-sbt.org/release/docs/Setup.html)**

## Recommended Installs
- **[IntelliJ IDEA Community Edition](https://www.jetbrains.com/idea/download/#section=windows)** (for Scala code)
- **[Postman](https://www.getpostman.com/)**


### Docker Set Up ###
The MPH environment runs a large number of containers and requires a minimum level of Docker resources to be allocated to function without errors. It is recommended that you provide Docker with at least 4 CPUs and 8GB of RAM.

### Postman Set Up ###
Postman can be used as an easy way to send REST calls to the MPH webapp to trigger events (i.e. data loading). This repository contains a *postman* folder which contains files which you can import to set up a collection of calls for calling Mph Club web services and the necessary environment variables. After importing these files, be sure to set your environment to *docker-local* before sending REST calls.

### Running Mph Club ###
The Docker environment can be brought up using the appropriate script for your OS: *docker-local.bat* for Windows or *docker-local.sh* for Mac. Pass a -f option to the script to rebuild (or build for the first time) the Mph Club project code; further, if it is your first time deploying the environment or you are unfamiliar with typical start times, it is recommended to remove the *-d* from the *docker-compose up* flag so that you can visually verify deployment status through the logs.

Upon successful start up, there should be 2 containers running:
- ct-mph-postgres (PostgreSQL)
- ct-mph-tomcat (Mph Club Tomcat)


If you unable to start all the containers, most likely Docker does not have enough resources and you should increase the amount of allocated RAM/CPUs. [Docker logs](https://docs.docker.com/engine/reference/commandline/logs/) or Kitematic may also be helpful for further diagnosing any issues.

After the environment is up, you may use Postman to access the Mph Club web services. Currently, this includes only data loading. First call *authenticate* to log into the system.

### Misc. Info
Below are some commands that developers may find useful for during development.

#### Useful Docker Commands ####
Bringing the environment up and down will eventually consume all available memory space for Docker and may cause issues with containers coming up or operating properly. Below are some helpful commands for freeing up space. **Note that these commands will delete persisted data in the environment.**

**Cleaning up Docker environment**
Remove all containers with:
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```

Remove dangling images with:
```
docker rmi $(docker images -f dangling=true -q)
```

**Cleaning up Docker volumes**
We can delete old Docker volumes with:
```
docker volume prune
```


