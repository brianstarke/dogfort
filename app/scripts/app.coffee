app = angular.module 'dogfort', [
  'ngRoute'
  'toastr'
  'angularMoment'

  'dogfort.controllers'
  'dogfort.services'
]

app.config ['$routeProvider', '$httpProvider', ($routeProvider, $httpProvider) ->
  $httpProvider.interceptors.push 'authInterceptor'

  $routeProvider.when '/login', {
    templateUrl: '/partials/login.html'
    controller: 'LoginCtrl'
  }

  $routeProvider.when '/channels', {
    templateUrl: '/partials/channels.html'
    controller: 'ChannelsCtrl'
  }

  $routeProvider.when '/chat', {
    templateUrl: '/partials/chat.html'
    controller: 'ChatCtrl'
  }

  $routeProvider.otherwise {
    redirectTo: '/login'
  }
]

app.config ['toastrConfig', (toastrConfig) ->
  toastrConfig.positionClass = 'toast-bottom-full-width'
]
