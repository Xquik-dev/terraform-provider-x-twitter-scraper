package logging

import (
	"context"
	"net/http"

	"github.com/Xquik-dev/x-twitter-scraper-go/option"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if req != nil {
			LogRequest(ctx, req)
		}

		resp, err := next(req)

		if resp != nil {
			LogResponse(ctx, resp)
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) {
	tflog.Debug(ctx, "HTTP request", map[string]any{
		"method":   req.Method,
		"path":     req.URL.EscapedPath(),
		"protocol": req.Proto,
	})
}

func LogResponse(ctx context.Context, resp *http.Response) {
	fields := map[string]any{
		"protocol":    resp.Proto,
		"status_code": resp.StatusCode,
	}
	if resp.Request != nil {
		fields["method"] = resp.Request.Method
		fields["path"] = resp.Request.URL.EscapedPath()
	}
	tflog.Debug(ctx, "HTTP response", fields)
}
