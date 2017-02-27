## gir is a simple, offline github issues reader in your terminal.

It utilizies the new Github GraphQL API which drastically reduces the amount of requests which have to be sent over the wire to download the issues. On the downside, this requires each user to join the [github early-access program](https://github.com/prerelease/agreement) and create a personal access token which has to be exposed as the *GITHUB_TOKEN* environment variable.

### Installation
    go install github.com/JanBerktold/gir

### Usage
gir is an offline github issues reader by github.com/JanBerktold
        Go to https://github.com/JanBerktold/gir for further information.

    > gir help
        Usage:
            gir [command]

        Available Commands:
            add         Cache an additional repository's issues
            cr          Selects an active repository which gir list and gir show work against
            drop        Removes a repository's cached data
            list        Lists all issues in one repository
            repos       Gives an overview of all cached repositories
            show        Gives detailed information about an issue including comments
            update      Updates cached issues

        Use "gir [command] --help" for more information about a command.
  
#### Current repository
gir has a concept of an "current repository" which allows you to omit the target repository when calling the *gir list* and *gir show* commands. You can set this current repository using the *gir cr* command.

    > gir cr --help
        gir cr <owner/name> allows you to set a default repository for gir list and gir show.
        
        Example:
            > gir repos
                golang/go (1600 issues)
            > gir list
                No repository specified. You either have to specify a repo as an argument or enter a scope.
            > gir cr golang/go
            > gir list
                1/CLOSED: Make Golang awesome (bradfitz)
                2/CLOSED: Think about generics (robpike)
                3/OPEN: Gir is a hip github issues reader (JanBerktold)
            > gir show 3
                issues info here

        Usage:
            gir cr <owner/name> [flags]
 
