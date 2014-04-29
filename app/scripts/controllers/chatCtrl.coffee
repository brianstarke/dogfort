app = angular.module 'dogfort'

class ChatCtrl extends BaseCtrl
  @register app
  @inject '$scope', '$location', 'Channel', 'Message', 'User', 'toastr', 'Penis'

  initialize: ->
    @$scope.channels          = {}
    @$scope.messages          = []
    @$scope.currentChannelId  = ''

    do @_refreshChannels

  # for highlighting channel tabs on the chat view
  isActive: (channelId) -> channelId is @$scope.currentChannelId

  changeChannel: (channelId) ->
    @$scope.currentChannelId = channelId

    do @_refreshMessages

  sendMessage: ->
    @Message.send(@$scope.message, @$scope.currentChannelId)
      .success =>
        @$scope.message = ''
        do @_refreshMessages
      .error =>
        @toastr.error data, 'ERROR'

  _refreshChannels: ->
    @Channel.userChannels()
      .success (data) =>
        @$scope.currentChannelId  = data.channels[0].uid
        @$scope.channels          = data.channels

        do @_refreshMessages
      .error (data) =>
        @toastr.error data, 'ERROR'

  _refreshMessages: ->
    @Message.forChannel(@$scope.currentChannelId)
      .success (data) =>
        for message in data
          @_addMessage message
      .error (data) =>
        @toastr.error data, 'ERROR'

  _addMessage: (message) ->
    @User.byId(message.userId)
      .success (data) =>
        message.user = data
        @$scope.messages.push message

        # just in case, make sure messages are in order
        @$scope.messages = @$scope.messages.sort (a,b) -> a.timestamp > b.timestamp
      .error (data) =>
        @toastr.error data, 'ERROR'
