# shorty
A URL shortener written in GO


## Frontend Assets Pipeline with Train and Pongo

### Preparation

Get the Packages:

``` sh
go get github.com/shaoshing/train
go get -u gopkg.in/flosch/pongo2.v3
```

Build train binary:

``` sh
go build -o $GOPATH/bin/train github.com/shaoshing/train/cmd
```

Setup Ruby, install sass and coffee gems:

``` sh
sudo gem install sass
sudo gem install coffee-script
```

Setup latest JRE (for YUI compressor).

### Building in Dev

Will happen automatically. Relax...

### Building in Production

Run `train` to automatically bundle and fingerprint assets.

