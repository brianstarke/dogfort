dogfort
=======

Share images and chat with your asshole friends.  Antisocial networking.

### .env file 

The following environment variables need to be set in a local ```.env``` file.

```
PORT=9000
HOST=localhost

MONGO_HOST=localhost
MONGO_DB=dogfort
```

### run

```
go get github.com/brianstarke/dogfort
cd $GOPATH/src/github.com/brianstarke/dogfort
go run main.go

OR

go install
dogfort
```
