package connect

import (
	"github.com/LilyPad/GoLilyPad/packet"
	"io"
)

type RequestGetPlayers struct {
	List bool
}

func NewRequestGetPlayers() (this *RequestGetPlayers) {
	this = new(RequestGetPlayers)
	this.List = false
	return
}

func NewRequestGetPlayersList() (this *RequestGetPlayers) {
	this = new(RequestGetPlayers)
	this.List = true
	return
}

func (this *RequestGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type requestGetPlayersCodec struct {
}

func (this *requestGetPlayersCodec) Decode(reader io.Reader) (request Request, err error) {
	requestGetPlayers := new(RequestGetPlayers)
	requestGetPlayers.List, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	request = requestGetPlayers
	return
}

func (this *requestGetPlayersCodec) Encode(writer io.Writer, request Request) (err error) {
	err = packet.WriteBool(writer, request.(*RequestGetPlayers).List)
	return
}

type ResultGetPlayers struct {
	List           bool
	CurrentPlayers uint32
	MaxPlayers     uint32
	Players        []string
}

func NewResultGetPlayers(currentPlayers uint32, maxPlayers uint32) (this *ResultGetPlayers) {
	this = new(ResultGetPlayers)
	this.List = false
	this.CurrentPlayers = currentPlayers
	this.MaxPlayers = maxPlayers
	return
}

func NewResultGetPlayersList(currentPlayers uint32, maxPlayers uint32, players []string) (this *ResultGetPlayers) {
	this = new(ResultGetPlayers)
	this.List = true
	this.CurrentPlayers = currentPlayers
	this.MaxPlayers = maxPlayers
	this.Players = players
	return
}

func (this *ResultGetPlayers) Id() int {
	return REQUEST_GET_PLAYERS
}

type resultGetPlayersCodec struct {
}

func (this *resultGetPlayersCodec) Decode(reader io.Reader) (result Result, err error) {
	resultGetPlayers := new(ResultGetPlayers)
	resultGetPlayers.List, err = packet.ReadBool(reader)
	if err != nil {
		return
	}
	resultGetPlayers.CurrentPlayers, err = packet.ReadUint32(reader)
	if err != nil {
		return
	}
	resultGetPlayers.MaxPlayers, err = packet.ReadUint32(reader)
	if err != nil {
		return
	}
	if resultGetPlayers.List {
		resultGetPlayers.Players = make([]string, resultGetPlayers.CurrentPlayers)
		var i uint32
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			resultGetPlayers.Players[i], err = packet.ReadString(reader)
		}
	}
	result = resultGetPlayers
	return
}

func (this *resultGetPlayersCodec) Encode(writer io.Writer, result Result) (err error) {
	resultGetPlayers := result.(*ResultGetPlayers)
	err = packet.WriteBool(writer, resultGetPlayers.List)
	if err != nil {
		return
	}
	err = packet.WriteUint32(writer, resultGetPlayers.CurrentPlayers)
	if err != nil {
		return
	}
	err = packet.WriteUint32(writer, resultGetPlayers.MaxPlayers)
	if resultGetPlayers.List {
		var i uint32
		for i = 0; i < resultGetPlayers.CurrentPlayers; i++ {
			if err != nil {
				return
			}
			err = packet.WriteString(writer, resultGetPlayers.Players[i])
		}
	}
	return
}
