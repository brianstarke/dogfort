app = angular.module 'dogfort.controllers', [
  'ngCookies'

  'dogfort.services'
]

app.controller 'ChatCtrl', ($scope, $location, $anchorScroll) ->
  $scope.chatMessages = [
    avatarUrl: 'http://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=50'
    chatText: 'testing'
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
    $location.hash 'bottom'
    $anchorScroll()
  , 10000

  setInterval ->
    $scope.chatMessages.push {
      avatarUrl: 'http://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=50'
      chatText: 'testing again'
      username: 'someone'
      ts: 'just now'
    }
    $scope.$digest()
    $location.hash 'bottom'
    $anchorScroll()
  , 9000

app.controller 'ChannelsCtrl', ($scope, Channel, toastr) ->
  modal = new $.UIkit.modal.Modal("#create")

  refreshChannels = () ->
    Channel.list()
      .success (data, status, headers, config) ->
        $scope.channels = data.channels
      .error (data, status, headers, config) ->
        toastr.error(data, 'ERROR')

  refreshChannels()

  $scope.create = () ->
    Channel.create({
      name: $scope.newchannel.name
      description: $scope.newchannel.description
      isPrivate: $scope.newchannel.isPrivate
    }).success((data, status, headers, config) ->
      modal._hide()
      $scope.newchannel = {}
      refreshChannels()
      toastr.success('Channel created!', 'SUCCESS')
    ).error((data, status, headers, config) ->
      modal._hide()
      toastr.error(data, 'ERROR')
    )

app.controller 'LoginCtrl', ($rootScope, $scope, $cookies, $location, User, toastr) ->
  modal = new $.UIkit.modal.Modal("#register")

  $scope.login = () ->
    User.authenticate($scope.user.username, $scope.user.password)
      .success (data, status, headers, config) ->
        $cookies.dogfort_token = data.token
        $location.path '/channels'
        $rootScope.isAuthenticated = true
        $rootScope.setAuthedUser()
        toastr.success('Authenticated', 'success')
      .error (data, status, headers, config) ->
        toastr.error(data, 'ERROR')

  $scope.register = () ->
    User.create({
      email: $scope.newuser.email
      username: $scope.newuser.username
      password: $scope.newuser.password
    }).success((data, status, headers, config) ->
      $location.path '/login'
      modal._hide()
      toastr.success('User created successfully, login', 'success')
    ).error((data, status, headers, config) ->
      toastr.error(data, 'ERROR')
    )

app.controller 'MainCtrl', ($rootScope, $scope, $cookies, $location, User) ->
  $rootScope.setAuthedUser = () ->
    User.getAuthedUser()
      .success (data, status, headers, config) ->
        $rootScope.authedUser = data.user
        $rootScope.isAuthenticated = true
        $location.path '/chat'
      .error (data, status, headers, config) ->
        $rootScope.isAuthenticated = false
        $location.path '/login'

  $rootScope.setAuthedUser()

  $scope.logout = () ->
    delete $cookies['dogfort_token']
    delete $rootScope['authedUser']

    $rootScope.isAuthenticated = false
    $location.path '/login'

  # for highlighting active tab on navbar
  $scope.isActive = (viewLocation) -> viewLocation == $location.path()
