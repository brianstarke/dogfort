app = angular.module 'dogfort', [
  'ngRoute'
  'ngCookies'

  'toastr'
  'angularMoment'
]

app.config [
  '$routeProvider'
  '$httpProvider'
  'toastrConfig'

  ($routeProvider, $httpProvider, toastrConfig) ->
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

    toastrConfig.positionClass = 'toast-bottom-full-width'
]
