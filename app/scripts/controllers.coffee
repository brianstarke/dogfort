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

app.controller 'ChannelsCtrl', ($scope, Channel) ->
  refreshChannels = () ->
    Channel.list()
      .success (data, status, headers, config) ->
        $scope.channels = data.channels
      .error (data, status, headers, config) ->
        console.log data

  refreshChannels()

  $scope.create = () ->
    Channel.create({
      name: $scope.newchannel.name
      description: $scope.newchannel.description
      isPrivate: $scope.newchannel.isPrivate
    }).success((data, status, headers, config) ->
      modal = new $.UIkit.modal.Modal("#create")
      modal._hide()
      $scope.newchannel = {}
      refreshChannels()
    ).error((data, status, headers, config) ->
      console.log data
    )

app.controller 'LoginCtrl', ($rootScope, $scope, $cookies, $location, User) ->
  $scope.login = () ->
    User.authenticate($scope.user.username, $scope.user.password)
      .success (data, status, headers, config) ->
        $cookies.dogfort_token = data.token
        $location.path '/channels'
        $rootScope.isAuthenticated = true
      .error (data, status, headers, config) ->
        console.log data

  $scope.register = () ->
    User.create({
      email: $scope.newuser.email
      username: $scope.newuser.username
      password: $scope.newuser.password
    }).success((data, status, headers, config) ->
      $location.path '/login'
    ).error((data, status, headers, config) ->
      console.log data
    )

app.controller 'MainCtrl', ($rootScope, $scope, $cookies, $location, User) ->
  User.getAuthedUser()
    .success (data, status, headers, config) ->
      $rootScope.authedUser = data.user
      $rootScope.isAuthenticated = true
      $location.path '/channels'
    .error (data, status, headers, config) ->
      $rootScope.isAuthenticated = false
      $location.path '/login'

  $scope.logout = () ->
    delete $cookies['dogfort_token']
    delete $rootScope['authedUser']
    $rootScope.isAuthenticated = false
    $location.path '/login'

