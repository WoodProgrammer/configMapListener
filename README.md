# KUBERNETES CONFIG MAP LISTENER
This module listen the file system and update configmaps.

## Usage : 
First you must build the watcher.go file after cloned project from <a href="https://github.com/WoodProgrammer/configMapListener"> github address</a>

    $ go build watcher.go
    $ ./watcher ABSOLUTE_FILE_PATH KUBERNETES_CONFIG_MAP_NAME
