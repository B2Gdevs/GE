# Google Extension (GE)
GE is a wrapper to the Go cli tool that helps uninstall packages installed with ```go get <package_name>```.

## Installation
```go get github.com/b2gdevs/ge```

## Usage
To use any go cli command with ```ge``` just replace ```go``` with ```ge```. 

### Example
```ge get github.com/codegangsta/gin``` could be used instead of ```go get github.com/codegangsta/gin```.  So can all of the
other go cli commands.  GE is just a wrapper to the cli.

### Feature
GE has but one job and that is to uninstall packages that have be installed with the ```get``` argument in ```ge get <package_name>``` or
```go get <package_name>```.  Lets say ```github.com/codegangsta/gin``` was installed and we want to remove all of its files src, bin, and
pkg files.  

We would just use: <br/>

```ge uninstall github.com/codegangsta/gin```


#### Caveat
```ge uninstall <package_name>``` uninstalls everything of **ONE** package.  It will not remove files/folders that do not match the package
name.  

For example:
```go get github.com/codegangsta/gin``` doesn't just install ```github.com/codegangsta/gin``` in the ```src``` folder it also installs
```github.com/codegangsta/envy```.  ```ge uninstall github.com/codegangsta/gin``` will only look for files/folders within
```github.com/codegangsta/gin```
