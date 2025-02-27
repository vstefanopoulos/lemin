package errors

import "errors"

var (
	ErrInvalidArguments      = errors.New("invalid arguments")
	ErrInvalidAntsInput      = errors.New("invalid ants input")
	ErrMissingStartRoom      = errors.New("missing start room")
	ErrMissingEndRoom        = errors.New("missing end room")
	ErrStartRoomDisconnected = errors.New("start room is disconnected")
	ErrEndRoomDisconnected   = errors.New("end room is disconnected")
	ErrInvalidTunnel         = errors.New("invalid tunnel")
	ErrTunnelExists          = errors.New("tunnel already exists")
	ErrInvalidRoomInput      = errors.New("invalid room input")
	ErrRoomAlreadyExists     = errors.New("room already exists")
	ErrCoord                 = errors.New("coordinates error")
	ErrStartEndNotConnected  = errors.New("start end not connected")
)
