// Main module-packages
module packages

go 1.19

// go get this module from github

// Running go get link ---creates another file that  go.sum which tracks dependecndies and versions of modules.It SHOULD NEVER BE EDITED.
require github.com/donvito/hellomod v1.0.1

require github.com/Vickopiyo/figures v0.0.0-20230126073519-4c7775055d37 // indirect
