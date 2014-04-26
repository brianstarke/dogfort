app = angular.module 'dogfort.controllers'

app.controller 'ChannelsCtrl', ($scope, Channel, toastr) ->
  modal = new $.UIkit.modal.Modal("#create")

  refreshChannels = () ->
    Channel.list()
      .success (data, status, headers, config) ->
        $scope.channels = data.channels
      .error (data, status, headers, config) ->
        toastr.error(data, 'ERROR')

  refreshChannels()

  $scope.join = (channelId) ->
    Channel.join(channelId)
      .success () ->
        toastr.success('Channel joined', 'SUCCESS')
        refreshChannels()
      .error (data) ->
        toastr.error(data, 'ERROR')

  $scope.leave = (channelId) ->
    Channel.leave(channelId)
      .success () ->
        toastr.success('Left channel', 'SUCCESS')
        refreshChannels()
      .error (data) ->
        toastr.error(data, 'ERROR')

  $scope.create = () ->
    Channel.create({
      name: $scope.newchannel.name
      description: $scope.newchannel.description
      isPrivate: $scope.newchannel.isPrivate
    }).success((data, status, headers, config) ->
      modal._hide()
      $scope.newchannel = {}
      refreshChannels()
      toastr.success('Channel created!', 'SUCCESS')
    ).error((data, status, headers, config) ->
      modal._hide()
      toastr.error(data, 'ERROR')
    )

