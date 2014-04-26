app = angular.module 'dogfort'

class ChannelsCtrl extends BaseCtrl
  @register app
  @inject '$scope', 'Channel', 'toastr'

  initialize: ->
    @$scope.channels = []

    @createModal = new $.UIkit.modal.Modal('#create')

    do @_refreshChannels

  join: (channelId) ->
    Channel.join(channelId)
      .success =>
        @toastr.success 'Channel joined', 'SUCCESS'
        do @_refreshChannels
      .error (data) =>
        @toastr.error data, 'ERROR'

  leave: (channelId) ->
    Channel.leave(channelId)
      .success =>
        @toastr.success 'Left channel', 'SUCCESS'
        do @_refreshChannels
      .error (data) =>
        @toastr.error data, 'ERROR'

  create: ->
    Channel.create(@$scope.newchannel)
      .success =>
        @createModal._hide()
        @$scope.newchannel = {}

        do @_refreshChannels
        @toastr.success 'Channel created!', 'SUCCESS'
      .error (data) =>
        @createModal._hide()
        @toastr.error data, 'ERROR'

  _refreshChannels: ->
    @Channel.list()
      .success (data) =>
        @$scope.channels = data.channels
      .error (data) =>
        @toastr.error data, 'ERROR'
