gulp    = require 'gulp'
jade    = require 'gulp-jade'
clean   = require 'gulp-clean'
coffee  = require 'gulp-coffee'
less    = require 'gulp-less'
concat  = require 'gulp-concat'

paths =
  jade: ['**/*.jade', '!node_modules/**/*.jade']
  scripts: 'scripts/**/*.coffee'
  less: './less/**/*.less'
  fonts: ['./bower_components/uikit/src/fonts/*.*']
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
gulp.task 'scripts', ->
  gulp.src(paths.scripts)
    .pipe(coffee())
    .pipe(concat('dogfort.js'))
    .pipe(gulp.dest(paths.dist + '/scripts'))

# copy fonts
gulp.task 'fonts', ->
  gulp.src(paths.fonts)
    .pipe gulp.dest(paths.dist + '/fonts')

# rerun the task when a file changes
gulp.task 'watch', ->
  gulp.watch paths.scripts, ['scripts']
  gulp.watch paths.jade, ['jade']
  gulp.watch paths.less, ['less']
  gulp.watch paths.fonts, ['fonts']

# do ALL THE THINGS
gulp.task 'build', ['clean', 'jade', 'less', 'fonts', 'scripts']

gulp.task 'default', ['clean', 'jade', 'less', 'fonts', 'scripts', 'watch']
