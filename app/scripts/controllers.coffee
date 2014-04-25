app = angular.module 'dogfort.controllers', [
  'ngCookies'

  'dogfort.services'
]

app.controller 'ChatCtrl', ($scope, $location, $anchorScroll, Channel, Message, User) ->
  $scope.channels = {}
  $scope.messages = []
  $scope.currentChannel = ''

  $scope.isActive = (channelId) -> channelId == $scope.currentChannel

  getChannels = () ->
    Channel.userChannels()
      .success (data, status, headers, config) ->
        $scope.currentChannel = data.channels[0].uid

        for channel in data.channels
          $scope.channels[channel.uid] = {}
          $scope.channels[channel.uid].channel = channel

        getMessages()

  $scope.changeChannel = (channelId) ->
    $scope.currentChannel = channelId

    getMessages()

  getMessages = () ->
    $scope.messages = []

    Message.forChannel($scope.currentChannel)
      .success (data, status, headers, config) ->
        for message in data
          addMessage message

  addMessage = (message) ->
    User.byId(message.userId)
      .success (data, status, headers, config) ->
        message.user = data
        $scope.messages.push message

        # just in case
        $scope.messages = $scope.messages.sort (a,b) -> a.timestamp > b.timestamp

      .error (data, status, headers, config) ->
        console.log data

  $scope.sendMessage = () ->
    Message.send($scope.message, $scope.currentChannel)
      .success (data, status, headers, config) ->
        $scope.message = ''

  getChannels()

app.controller 'ChannelsCtrl', ($scope, Channel, toastr) ->
  modal = new $.UIkit.modal.Modal("#create")

  refreshChannels = () ->
    Channel.list()
      .success (data, status, headers, config) ->
        $scope.channels = data.channels
      .error (data, status, headers, config) ->
        toastr.error(data, 'ERROR')

  refreshChannels()

  $scope.join = (channelId) ->
    Channel.join(channelId)
      .success () ->
        toastr.success('Channel joined', 'SUCCESS')
      .error (data) ->
        toastr.error(data, 'ERROR')

  $scope.leave = (channelId) ->
    Channel.leave(channelId)
      .success () ->
        toastr.success('Left channel', 'SUCCESS')
      .error (data) ->
        toastr.error(data, 'ERROR')

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
        connectToSocket()
      .error (data, status, headers, config) ->
        $rootScope.isAuthenticated = false
        $location.path '/login'

  $rootScope.setAuthedUser()

  connectToSocket = () ->
    if window["WebSocket"]
      console.log 'browser supports WebSockets'
    else
      alert "Your browser doesn't support WebSockets, this app will suck for you."

  $scope.logout = () ->
    delete $cookies['dogfort_token']
    delete $rootScope['authedUser']

    $rootScope.isAuthenticated = false
    $location.path '/login'

  # for highlighting active tab on navbar
  $scope.isActive = (viewLocation) -> viewLocation == $location.path()
