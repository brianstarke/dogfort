'use strict'

app = angular.module 'dogfort', []

app.controller 'ChatCtrl', ($scope) ->
  $scope.chatMessages = [
    avatarUrl: 'http://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=50'
    chatText: 'Monkeyfighting shoot mothercrusher fudge shoot.'
    username: 'someone'
    ts: '4 minutes ago'
  ]

  setInterval ->
    $scope.chatMessages.push {
      avatarUrl: 'http://www.gravatar.com/avatar/9f6fe08431ce0e906f6b2e7dd5c9a812?s=50'
      chatText: 'Clown balony melonfarmer funster airhead bloodsucker!'
      username: 'starke'
      ts: 'just now'
    }
    $scope.$digest()
  , 4000




