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

### run server

```
go get github.com/brianstarke/dogfort
cd $GOPATH/src/github.com/brianstarke/dogfort
go run main.go

OR

go install
dogfort
```

### rebuild app

The frontend app lives in the ```app/``` directory and the jade, less, and coffeescript are compiled by (gulp)[https://github.com/gulpjs/gulp] and copied to the ```public/``` folder.

You'll need node (0.10 or greater) and gulp.

```
npm install -g gulp
npm install
gulp --require coffee-script/register build
``` 