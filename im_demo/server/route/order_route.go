package route

import (
	"encoding/json"
	"net"
)

type Ctx struct {
	Conn  net.Conn
	Order string      `json:"order"`
	Param interface{} `json:"param"`
}

type HandlerFn = func(ctx Ctx)

type OrderRoute struct {
	orderMap map[string]HandlerFn
}

func (r *OrderRoute) Register(order string, handlerFn HandlerFn) {
	if r.orderMap == nil {
		r.orderMap = make(map[string]HandlerFn)
	}
	r.orderMap[order] = handlerFn
}

func (r OrderRoute) GetHandlerFn(order string) HandlerFn {
	if r.orderMap == nil {
		return nil
	}
	return r.orderMap[order]
}

func (r OrderRoute) MsgHandle(conn net.Conn, msg []byte) {
	ctx := Ctx{
		Conn:  conn,
		Order: "",
		Param: nil,
	}
	err := json.Unmarshal(msg, &ctx)
	if err != nil {
		conn.Write([]byte("invalid message\n"))
		return
	}
	fn := r.orderMap[ctx.Order]
	if fn != nil {
		fn(ctx)
	} else {
		conn.Write([]byte("not find order"))
	}
}

func (c *Ctx) ReturnMsg(msg string) error {
	_, err := c.Conn.Write([]byte(msg))
	return err
}
