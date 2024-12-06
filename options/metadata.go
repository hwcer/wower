package options

const (
	ServiceMetadataUID       = "uid"
	ServiceMetadataGUID      = "guid"
	ServiceMetadataServerId  = "sid"
	ServiceMetadataRequestId = "_rid"

	ServiceMessagePath   = "_msg_path"
	ServiceMessageRoom   = "_msg_room"
	ServiceMessageIgnore = "_msg_ignore"

	ServicePlayerOAuth  = "_player_oauth"
	ServicePlayerLogout = "_player_logout"

	ServicePlayerRoomJoin  = "player.room.join"  //已经加入的房间
	ServicePlayerRoomLeave = "player.room.leave" //离开房间
	ServicePlayerSelector  = "service.selector." //服务器重定向
)
