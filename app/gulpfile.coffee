gulp 	  = require 'gulp'
jade 	  = require 'gulp-jade'
clean   = require 'gulp-clean'
coffee  = require 'gulp-coffee'
concat  = require 'gulp-concat'

paths =
  jade: ['**/*.jade', '!node_modules/**/*.jade']
  scripts: 'scripts/**/*.coffee'
  dist: '../public'

# clean public folder
gulp.task 'clean', ->
  gulp.src(paths.dist, {read: false})
    .pipe clean({force: true})

# compile jade templates
gulp.task 'jade', ->
  gulp.src(paths.jade)
  	.pipe(jade())
  	.pipe(gulp.dest(paths.dist))

# compile/concat coffeescript
gulp.task 'scripts', ->
  gulp.src(paths.scripts)
    .pipe(coffee())
    .pipe(concat('dogfort.js'))
    .pipe(gulp.dest(paths.dist + '/scripts'))

# do ALL THE THINGS
gulp.task 'build', ['clean', 'jade', 'scripts']
