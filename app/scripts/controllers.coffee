app = angular.module 'dogfort.controllers', [
  'ngCookies'

  'dogfort.services'
]

app.controller 'ChatCtrl', ($scope) ->
  $scope.chatMessages = [
    avatarUrl: 'http://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=50'
    chatText: 'Monkeyfighting shoot mothercrusher fudge shoot.'
    username: 'someone'
    ts: '4 minutes ago'
  ]

  setInterval ->
    $scope.chatMessages.push {
      avatarUrl: 'http://www.gravatar.com/avatar/9f6fe08431ce0e906f6b2e7dd5c9a812?s=50'
      chatText: 'bloop'
      username: 'starke'
      ts: 'just now'
    }
    $scope.$digest()
  , 10000

app.controller 'MainCtrl', ($scope, $cookies, User) ->
  User.verify($cookies.dogfort_token)
    .success (data, status, headers, config) ->
      console.log data
      $scope.authedUser = data.user
      $scope.isAuthenticated = true
    .error (data, status, headers, config) ->
      console.log data
      $scope.isAuthenticated = false

  $scope.login = () ->
    User.authenticate($scope.user.username, $scope.user.password)
      .success (data, status, headers, config) ->
        $cookies.dogfort_token = data.token
      .error (data, status, headers, config) ->
        console.log data

    $scope.isAuthenticated = true

  $scope.register = () ->
    console.log "registering new user"

    User.create({
      email: $scope.newuser.email
      username: $scope.newuser.username
      password: $scope.newuser.password
    }).success((data, status, headers, config) ->
      console.log data
    ).error((data, status, headers, config) ->
      console.log data
    )
