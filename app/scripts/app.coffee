app = angular.module 'dogfort', [
  'ngRoute'
  'toastr'

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

  positionClass: 'toast-top-right'
]

app.config ['toastrConfig', (toastrConfig) ->
  toastrConfig.positionClass = 'toast-bottom-full-width'
]
