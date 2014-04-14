gulp    = require 'gulp'
jade    = require 'gulp-jade'
clean   = require 'gulp-clean'
coffee  = require 'gulp-coffee'
less    = require 'gulp-less'
concat  = require 'gulp-concat'
uglify  = require 'gulp-uglify'

paths =
  jade: ['**/*.jade', '!node_modules/**/*.jade']
  coffee: ['scripts/**/*.coffee']
  scripts: [
    './bower_components/jquery/jquery.min.js'
    './bower_components/uikit/dist/js/uikit.min.js'
    './bower_components/uikit/dist/js/addons/sticky.min.js'
    './bower_components/angular/angular.min.js'
  ]
  less: './less/**/*.less'
  fonts: ['./bower_components/uikit/src/fonts/*.*']
  images: './images/*.ico'
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

# compile less styles
gulp.task 'less', ->
  gulp.src(paths.less)
    .pipe(less())
    .pipe(gulp.dest(paths.dist + '/styles'))

# compile/concat coffeescript
gulp.task 'coffee', ->
  gulp.src(paths.coffee)
    .pipe(coffee())
    .pipe(concat('dogfort.js'))
    .pipe(gulp.dest(paths.dist + '/scripts'))

gulp.task 'scripts', ->
  gulp.src(paths.scripts)
    .pipe(uglify())
    .pipe(concat('thirdparty.min.js'))
    .pipe(gulp.dest(paths.dist + '/scripts'))

# copy fonts
gulp.task 'fonts', ->
  gulp.src(paths.fonts)
    .pipe gulp.dest(paths.dist + '/fonts')

# copy images
gulp.task 'images', ->
  gulp.src(paths.images)
    .pipe gulp.dest(paths.dist + '/images')

# rerun the task when a file changes
gulp.task 'watch', ->
  gulp.watch paths.coffee, ['coffee']
  gulp.watch paths.jade, ['jade']
  gulp.watch paths.less, ['less']
  gulp.watch paths.fonts, ['fonts']

# do ALL THE THINGS
gulp.task 'build', ['clean', 'jade', 'less', 'fonts', 'scripts', 'images', 'coffee']

gulp.task 'default', ['build', 'watch']
