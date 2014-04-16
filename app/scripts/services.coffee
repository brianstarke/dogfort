app = angular.module 'dogfort.services', []

app.factory 'User', ['$http', ($http) ->
  new class User
    constructor: ->
      @baseUrl = '/api/v1/users'

    create: (newUser) ->
      $http.post @baseUrl, newUser
]

