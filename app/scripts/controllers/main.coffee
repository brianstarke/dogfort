app = angular.module 'dogfort.controllers'

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
      conn = new WebSocket 'ws://localhost:9000/ws/connect'
      conn.onclose = (event) ->
        console.log 'socket connection closed'

      conn.onmessage = (event) ->
        console.log event.data
    else
      alert "Your browser doesn't support WebSockets, this app will suck for you."

  $scope.logout = () ->
    delete $cookies['dogfort_token']
    delete $rootScope['authedUser']

    $rootScope.isAuthenticated = false
    $location.path '/login'

  # for highlighting active tab on navbar
  $scope.isActive = (viewLocation) -> viewLocation == $location.path()
