package rbac

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"billionmail-core/api/rbac/v1"
)

func (c *ControllerV1) PermissionCheck(ctx context.Context, req *v1.PermissionCheckReq) (res *v1.PermissionCheckRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
