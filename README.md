# KUBERNETES CONFIG MAP LISTENER
This module listen the file system and update configmaps.

## Usage : 
### Notice:
If you wanna run this tool on your Kubernetes Cluster . You must create your configmaps before . [It run existings ConfigMaps ]


First you must build the watcher.go file after cloned project from <a href="https://github.com/WoodProgrammer/configMapListener"> github address</a>

    $ go build watcher.go
    $ ./watcher ABSOLUTE_FILE_PATH KUBERNETES_CONFIG_MAP_NAME
