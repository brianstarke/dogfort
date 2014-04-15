app = angular.module 'dogfort', [
  'ngRoute'

  'dogfort.controllers'
  'dogfort.services'
]

app.config ['$routeProvider', ($routeProvider) ->
  $routeProvider.when '/', {
    templateUrl: '/partials/main.html'
    controller: 'ChatCtrl'
  }

  $routeProvider.otherwise {
    redirectTo: '/'
  }
]





