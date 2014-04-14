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
```

### rebuild app

The frontend app lives in the ```app/``` directory and the jade, less, and coffeescript are compiled by [gulp](https://github.com/gulpjs/gulp) and copied to the ```public/``` folder.

You'll need [node](http://nodejs.org/), [gulp](https://github.com/gulpjs/gulp), and [bower](http://bower.io/).

```
cd app
npm install -g gulp
npm install -g bower
npm install
gulp --require coffee-script/register build
```

That will create the ```public/``` folder which will be served by Martini.

### dev mode

Easiest way to work on dogfort frontend is to run ```gulp --require coffee-script/register``` in the ```app``` directory which will put the ```app/``` directory in watch mode.  Updating files in that directory will automatically rebuild everything.  This way you can run ```go run main.go``` from the main directory and just keep refreshing the browser to see changes.
