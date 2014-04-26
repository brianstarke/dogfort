app = angular.module 'dogfort.controllers'

app.controller 'ChatCtrl', ($scope, $location, $anchorScroll, Channel, Message, User) ->
  $scope.channels = {}
  $scope.messages = []
  $scope.currentChannel = ''

  $scope.isActive = (channelId) -> channelId == $scope.currentChannel

  getChannels = () ->
    Channel.userChannels()
      .success (data, status, headers, config) ->
        $scope.currentChannel = data.channels[0].uid

        for channel in data.channels
          $scope.channels[channel.uid] = {}
          $scope.channels[channel.uid].channel = channel

        getMessages()

  $scope.changeChannel = (channelId) ->
    $scope.currentChannel = channelId

    getMessages()

  getMessages = () ->
    $scope.messages = []

    Message.forChannel($scope.currentChannel)
      .success (data, status, headers, config) ->
        for message in data
          addMessage message

  addMessage = (message) ->
    User.byId(message.userId)
      .success (data, status, headers, config) ->
        message.user = data
        $scope.messages.push message

        # just in case
        $scope.messages = $scope.messages.sort (a,b) -> a.timestamp > b.timestamp

      .error (data, status, headers, config) ->
        console.log data

  $scope.sendMessage = () ->
    Message.send($scope.message, $scope.currentChannel)
      .success (data, status, headers, config) ->
        $scope.message = ''
        getMessages()

  getChannels()
