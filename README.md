revel-modz
==========

Modules, samples, and a skeleton for the Revel Framework


Get the dependencies
--------------

### Go:

Install
``` Bash
cd $HOME
sudo apt-get install gcc libc6-dev mercurial git
hg clone -u release https://code.google.com/p/go
cd go/src
./all.bash
cd $HOME
mkdir -p gocode/{src,bin,pkg}
```

Add the following to your `.profile`
``` Bash
export GOROOT=$HOME/go
export GOBIN=$GOROOT/bin
export GOPATH=$HOME/gocode
export GO_BIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GO_BIN:
```

### Revel:
``` Bash
go get -u github.com/revel/cmd/revel
```

### Grunt:

Add the following to your `.profile`
``` Bash
alias node='nodejs '
```

Run the following commands
``` Bash
sudo apt-get install nodejs npm ruby
sudo gem install sass
sudo npm install -g grunt-cli highlight.js
sudo npm install -g grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-contrib-sass grunt-contrib-less
```

### Postgres: 

Install the packages
``` Bash
sudo apt-get install postgresql postgresql-client
```

Setup the database user
``` Bash
sudo -u postgres createuser -s -P <username>
```

Installation
--------------

### Install revel-modz

``` Bash
go get -u github.com/iassic/revel-modz
```

### Setup the sample

``` Bash
cd $GOPATH/src/github.com/iassic/revel-modz/sample
bash init.sh (hit ctrl-c when prompted) [you will see a bunch of errors initially]
```

Create the sample database

``` Bash
createdb sample_dev_db
```

Add the following to your `.profile`
``` Bash
export DB_DEV_USER='<username>'
export DB_DEV_PASS='<password>'
export DB_DEV_NAME='sample_dev_db'
export DB_PROD_USER='<username>'
export DB_PROD_PASS='<password>'
export DB_PROD_NAME='sample_dev_db'
```


Usage
---------------

The following instructions will setup a new app from the revel-modz skeleton

### Create the new app

``` Bash
revel new <APP_NAME> github.com/iassic/revel-modz/skeleton
cd <APP_NAME>
bash init.sh  (hit ctrl-c when prompted) [you will see a bunch of errors initially]
cd ..
```

### Create the databases for the app

``` Bash
createdb <APP_NAME>_dev_db
createdb <APP_NAME>_prod_db
```

### Setup the environment variables

Add the following to your `.profile`
(you may change the export names if you also change them in app.conf)
[You can have several apps and DBs by altering the environment variables and respective app.conf's]
``` Bash
export DB_DEV_USER='<username>'
export DB_DEV_PASS='<password>'
export DB_DEV_NAME='<APP_NAME>_dev_db'
export DB_PROD_USER='<username>'
export DB_PROD_PASS='<password>'
export DB_PROD_NAME='<APP_NAME>_prod_db'
```

### Run your new Revel application!

```
revel run <APP_NAME>
```

View by pointing your browser at `localhost:9000`


Features
----------------

### Front-end:

- Foundation 5.1.1
- Headjs for asynchronous loading of assets
- Many JS/CSS goodies in revel-modz/modules/assets
- Templated includes for per page assets
- An `appendjs` template function for inserting JS code

### Back-end:

- JS/SASS app resources initialized in app/assets
- Hot Code watch and recompile with Grunt
- ORM with github.com/jinzhu/gorm

### Security:

- User Authentication
- CSRF protection
- `X-Frame-Options` `X-XXS-Protection` and `X-Content-Type-Options` headers

### Modules:

- assets
 -- lots of css/js goodies
- grunt
 -- manage per app assets

- user/auth
 -- registration
 -- authentication
 -- profile

- ws_comm
 -- client side
 -- server side

- plugins
 -- maillist
 -- forums
 -- user-files
 -- analytics
   --- page requests
   --- ui interaction testing

The individual modules have (will have) their own README's with more detail about each

Sample
----------------

The sample is a runnable Revel application, though you may have to do some setup

The skeleton mirrors the sample and both include all modules
