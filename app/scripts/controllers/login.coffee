app = angular.module 'dogfort.controllers'

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
