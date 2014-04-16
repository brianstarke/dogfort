app = angular.module 'dogfort.services', []

app.factory 'User', ['$http', ($http) ->
  new class User
    constructor: ->
      @baseUrl = '/api/v1'

    create: (newUser) ->
      $http.post @baseUrl + '/users', newUser

    authenticate: (username, password) ->
      $http.post @baseUrl + '/authenticate', {
        username: username
        password: password
      }

    verify: (token) ->
      $http.get @baseUrl + '/verify/' + token
]

