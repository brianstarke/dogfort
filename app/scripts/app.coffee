app = angular.module 'dogfort', [
  'ngRoute'

  'dogfort.controllers'
  'dogfort.services'
]

app.config ['$routeProvider', '$httpProvider', ($routeProvider, $httpProvider) ->
  $httpProvider.interceptors.push 'authInterceptor'

  $routeProvider.when '/channels', {
    templateUrl: '/partials/channels.html'
    controller: 'ChatCtrl'
  }

  $routeProvider.when '/login', {
    templateUrl: '/partials/login.html'
  }

  $routeProvider.otherwise {
    redirectTo: '/channels'
  }
]
