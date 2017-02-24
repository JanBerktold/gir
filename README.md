## gir is a simple, offline github issues reader in your terminal.

It utilizies the new Github GraphQL API which drastically reduces the amount of requests which have to be sent over the wire to download the issues. On the downside, this requires each user to join the [github early-access program](https://github.com/prerelease/agreement) and create a personal access token which has to be exposed as the *GITHUB_TOKEN* environment variable.

### Installation
    go install github.com/JanBerktold/gir

### Usage
    > gir help
    gir is an offline github issues reader by github.com/JanBerktold
		Go to https://github.com/JanBerktold/gir for further information.

    Usage:
      gir [command]

    Available Commands:
       add         Cache an additional repository's issues
       cd          
       list        
       repos       Gives an overview of all cached repositories
       show        
       update      Updates cached issues
    
