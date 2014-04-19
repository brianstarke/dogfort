app = angular.module 'dogfort.services', [
  'ngCookies'
]

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

    getAuthedUser: ->
      $http.get @baseUrl + '/user'
]

app.factory 'Channel', ['$http', ($http) ->
  new class Channel
    constructor: ->
      @baseUrl = '/api/v1/channels'

    list: ->
      $http.get @baseUrl

    create: (newChannel) ->
      $http.post @baseUrl, newChannel
]

app.factory 'authInterceptor',  ['$q', '$cookies', '$location', ($q, $cookies, $location) ->
  request: (config) ->
    config.headers = config.headers or {}

    if $cookies.dogfort_token
      config.headers.Authorization = $cookies.dogfort_token
    config
  response: (response) ->
    if response.status is 401
      $location.path '/login'
      console.log 'etf'

    response or $q.when(response)
]
