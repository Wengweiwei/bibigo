// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: roomNotice.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	roomNotice.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRoomNoticeBuyGuard = "/live.xuserex.v1.RoomNotice/buy_guard"
var PathRoomNoticeIsTaskFinish = "/live.xuserex.v1.RoomNotice/is_task_finish"
var PathRoomNoticeSetTaskFinish = "/live.xuserex.v1.RoomNotice/set_task_finish"

// ====================
// RoomNotice Interface
// ====================

// 房间提示 相关服务
type RoomNoticeBMServer interface {
	// 是否弹出大航海购买提示
	BuyGuard(ctx context.Context, req *RoomNoticeBuyGuardReq) (resp *RoomNoticeBuyGuardResp, err error)

	// habse 任务是否结束
	IsTaskFinish(ctx context.Context, req *RoomNoticeIsTaskFinishReq) (resp *RoomNoticeIsTaskFinishResp, err error)

	// 手动设置base 任务结束
	SetTaskFinish(ctx context.Context, req *RoomNoticeSetTaskFinishReq) (resp *RoomNoticeSetTaskFinishResp, err error)
}

var v1RoomNoticeSvc RoomNoticeBMServer

func roomNoticeBuyGuard(c *bm.Context) {
	p := new(RoomNoticeBuyGuardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomNoticeSvc.BuyGuard(c, p)
	c.JSON(resp, err)
}

func roomNoticeIsTaskFinish(c *bm.Context) {
	p := new(RoomNoticeIsTaskFinishReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomNoticeSvc.IsTaskFinish(c, p)
	c.JSON(resp, err)
}

func roomNoticeSetTaskFinish(c *bm.Context) {
	p := new(RoomNoticeSetTaskFinishReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomNoticeSvc.SetTaskFinish(c, p)
	c.JSON(resp, err)
}

// RegisterRoomNoticeBMServer Register the blademaster route
func RegisterRoomNoticeBMServer(e *bm.Engine, server RoomNoticeBMServer) {
	v1RoomNoticeSvc = server
	e.GET("/live.xuserex.v1.RoomNotice/buy_guard", roomNoticeBuyGuard)
	e.GET("/live.xuserex.v1.RoomNotice/is_task_finish", roomNoticeIsTaskFinish)
	e.GET("/live.xuserex.v1.RoomNotice/set_task_finish", roomNoticeSetTaskFinish)
}