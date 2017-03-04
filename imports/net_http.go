// this file was generated by gomacro command: import "net/http"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	pkg "net/http"
	. "reflect"
)

func init() {
	Binds["net/http"] = map[string]Value{
		"CanonicalHeaderKey":	ValueOf(pkg.CanonicalHeaderKey),
		"DefaultClient":	ValueOf(&pkg.DefaultClient).Elem(),
		"DefaultMaxHeaderBytes":	ValueOf(pkg.DefaultMaxHeaderBytes),
		"DefaultMaxIdleConnsPerHost":	ValueOf(pkg.DefaultMaxIdleConnsPerHost),
		"DefaultServeMux":	ValueOf(&pkg.DefaultServeMux).Elem(),
		"DefaultTransport":	ValueOf(&pkg.DefaultTransport).Elem(),
		"DetectContentType":	ValueOf(pkg.DetectContentType),
		"ErrAbortHandler":	ValueOf(&pkg.ErrAbortHandler).Elem(),
		"ErrBodyNotAllowed":	ValueOf(&pkg.ErrBodyNotAllowed).Elem(),
		"ErrBodyReadAfterClose":	ValueOf(&pkg.ErrBodyReadAfterClose).Elem(),
		"ErrContentLength":	ValueOf(&pkg.ErrContentLength).Elem(),
		"ErrHandlerTimeout":	ValueOf(&pkg.ErrHandlerTimeout).Elem(),
		"ErrHeaderTooLong":	ValueOf(&pkg.ErrHeaderTooLong).Elem(),
		"ErrHijacked":	ValueOf(&pkg.ErrHijacked).Elem(),
		"ErrLineTooLong":	ValueOf(&pkg.ErrLineTooLong).Elem(),
		"ErrMissingBoundary":	ValueOf(&pkg.ErrMissingBoundary).Elem(),
		"ErrMissingContentLength":	ValueOf(&pkg.ErrMissingContentLength).Elem(),
		"ErrMissingFile":	ValueOf(&pkg.ErrMissingFile).Elem(),
		"ErrNoCookie":	ValueOf(&pkg.ErrNoCookie).Elem(),
		"ErrNoLocation":	ValueOf(&pkg.ErrNoLocation).Elem(),
		"ErrNotMultipart":	ValueOf(&pkg.ErrNotMultipart).Elem(),
		"ErrNotSupported":	ValueOf(&pkg.ErrNotSupported).Elem(),
		"ErrServerClosed":	ValueOf(&pkg.ErrServerClosed).Elem(),
		"ErrShortBody":	ValueOf(&pkg.ErrShortBody).Elem(),
		"ErrSkipAltProtocol":	ValueOf(&pkg.ErrSkipAltProtocol).Elem(),
		"ErrUnexpectedTrailer":	ValueOf(&pkg.ErrUnexpectedTrailer).Elem(),
		"ErrUseLastResponse":	ValueOf(&pkg.ErrUseLastResponse).Elem(),
		"ErrWriteAfterFlush":	ValueOf(&pkg.ErrWriteAfterFlush).Elem(),
		"Error":	ValueOf(pkg.Error),
		"FileServer":	ValueOf(pkg.FileServer),
		"Get":	ValueOf(pkg.Get),
		"Handle":	ValueOf(pkg.Handle),
		"HandleFunc":	ValueOf(pkg.HandleFunc),
		"Head":	ValueOf(pkg.Head),
		"ListenAndServe":	ValueOf(pkg.ListenAndServe),
		"ListenAndServeTLS":	ValueOf(pkg.ListenAndServeTLS),
		"LocalAddrContextKey":	ValueOf(&pkg.LocalAddrContextKey).Elem(),
		"MaxBytesReader":	ValueOf(pkg.MaxBytesReader),
		"MethodConnect":	ValueOf(pkg.MethodConnect),
		"MethodDelete":	ValueOf(pkg.MethodDelete),
		"MethodGet":	ValueOf(pkg.MethodGet),
		"MethodHead":	ValueOf(pkg.MethodHead),
		"MethodOptions":	ValueOf(pkg.MethodOptions),
		"MethodPatch":	ValueOf(pkg.MethodPatch),
		"MethodPost":	ValueOf(pkg.MethodPost),
		"MethodPut":	ValueOf(pkg.MethodPut),
		"MethodTrace":	ValueOf(pkg.MethodTrace),
		"NewFileTransport":	ValueOf(pkg.NewFileTransport),
		"NewRequest":	ValueOf(pkg.NewRequest),
		"NewServeMux":	ValueOf(pkg.NewServeMux),
		"NoBody":	ValueOf(&pkg.NoBody).Elem(),
		"NotFound":	ValueOf(pkg.NotFound),
		"NotFoundHandler":	ValueOf(pkg.NotFoundHandler),
		"ParseHTTPVersion":	ValueOf(pkg.ParseHTTPVersion),
		"ParseTime":	ValueOf(pkg.ParseTime),
		"Post":	ValueOf(pkg.Post),
		"PostForm":	ValueOf(pkg.PostForm),
		"ProxyFromEnvironment":	ValueOf(pkg.ProxyFromEnvironment),
		"ProxyURL":	ValueOf(pkg.ProxyURL),
		"ReadRequest":	ValueOf(pkg.ReadRequest),
		"ReadResponse":	ValueOf(pkg.ReadResponse),
		"Redirect":	ValueOf(pkg.Redirect),
		"RedirectHandler":	ValueOf(pkg.RedirectHandler),
		"Serve":	ValueOf(pkg.Serve),
		"ServeContent":	ValueOf(pkg.ServeContent),
		"ServeFile":	ValueOf(pkg.ServeFile),
		"ServerContextKey":	ValueOf(&pkg.ServerContextKey).Elem(),
		"SetCookie":	ValueOf(pkg.SetCookie),
		"StateActive":	ValueOf(pkg.StateActive),
		"StateClosed":	ValueOf(pkg.StateClosed),
		"StateHijacked":	ValueOf(pkg.StateHijacked),
		"StateIdle":	ValueOf(pkg.StateIdle),
		"StateNew":	ValueOf(pkg.StateNew),
		"StatusAccepted":	ValueOf(pkg.StatusAccepted),
		"StatusAlreadyReported":	ValueOf(pkg.StatusAlreadyReported),
		"StatusBadGateway":	ValueOf(pkg.StatusBadGateway),
		"StatusBadRequest":	ValueOf(pkg.StatusBadRequest),
		"StatusConflict":	ValueOf(pkg.StatusConflict),
		"StatusContinue":	ValueOf(pkg.StatusContinue),
		"StatusCreated":	ValueOf(pkg.StatusCreated),
		"StatusExpectationFailed":	ValueOf(pkg.StatusExpectationFailed),
		"StatusFailedDependency":	ValueOf(pkg.StatusFailedDependency),
		"StatusForbidden":	ValueOf(pkg.StatusForbidden),
		"StatusFound":	ValueOf(pkg.StatusFound),
		"StatusGatewayTimeout":	ValueOf(pkg.StatusGatewayTimeout),
		"StatusGone":	ValueOf(pkg.StatusGone),
		"StatusHTTPVersionNotSupported":	ValueOf(pkg.StatusHTTPVersionNotSupported),
		"StatusIMUsed":	ValueOf(pkg.StatusIMUsed),
		"StatusInsufficientStorage":	ValueOf(pkg.StatusInsufficientStorage),
		"StatusInternalServerError":	ValueOf(pkg.StatusInternalServerError),
		"StatusLengthRequired":	ValueOf(pkg.StatusLengthRequired),
		"StatusLocked":	ValueOf(pkg.StatusLocked),
		"StatusLoopDetected":	ValueOf(pkg.StatusLoopDetected),
		"StatusMethodNotAllowed":	ValueOf(pkg.StatusMethodNotAllowed),
		"StatusMovedPermanently":	ValueOf(pkg.StatusMovedPermanently),
		"StatusMultiStatus":	ValueOf(pkg.StatusMultiStatus),
		"StatusMultipleChoices":	ValueOf(pkg.StatusMultipleChoices),
		"StatusNetworkAuthenticationRequired":	ValueOf(pkg.StatusNetworkAuthenticationRequired),
		"StatusNoContent":	ValueOf(pkg.StatusNoContent),
		"StatusNonAuthoritativeInfo":	ValueOf(pkg.StatusNonAuthoritativeInfo),
		"StatusNotAcceptable":	ValueOf(pkg.StatusNotAcceptable),
		"StatusNotExtended":	ValueOf(pkg.StatusNotExtended),
		"StatusNotFound":	ValueOf(pkg.StatusNotFound),
		"StatusNotImplemented":	ValueOf(pkg.StatusNotImplemented),
		"StatusNotModified":	ValueOf(pkg.StatusNotModified),
		"StatusOK":	ValueOf(pkg.StatusOK),
		"StatusPartialContent":	ValueOf(pkg.StatusPartialContent),
		"StatusPaymentRequired":	ValueOf(pkg.StatusPaymentRequired),
		"StatusPermanentRedirect":	ValueOf(pkg.StatusPermanentRedirect),
		"StatusPreconditionFailed":	ValueOf(pkg.StatusPreconditionFailed),
		"StatusPreconditionRequired":	ValueOf(pkg.StatusPreconditionRequired),
		"StatusProcessing":	ValueOf(pkg.StatusProcessing),
		"StatusProxyAuthRequired":	ValueOf(pkg.StatusProxyAuthRequired),
		"StatusRequestEntityTooLarge":	ValueOf(pkg.StatusRequestEntityTooLarge),
		"StatusRequestHeaderFieldsTooLarge":	ValueOf(pkg.StatusRequestHeaderFieldsTooLarge),
		"StatusRequestTimeout":	ValueOf(pkg.StatusRequestTimeout),
		"StatusRequestURITooLong":	ValueOf(pkg.StatusRequestURITooLong),
		"StatusRequestedRangeNotSatisfiable":	ValueOf(pkg.StatusRequestedRangeNotSatisfiable),
		"StatusResetContent":	ValueOf(pkg.StatusResetContent),
		"StatusSeeOther":	ValueOf(pkg.StatusSeeOther),
		"StatusServiceUnavailable":	ValueOf(pkg.StatusServiceUnavailable),
		"StatusSwitchingProtocols":	ValueOf(pkg.StatusSwitchingProtocols),
		"StatusTeapot":	ValueOf(pkg.StatusTeapot),
		"StatusTemporaryRedirect":	ValueOf(pkg.StatusTemporaryRedirect),
		"StatusText":	ValueOf(pkg.StatusText),
		"StatusTooManyRequests":	ValueOf(pkg.StatusTooManyRequests),
		"StatusUnauthorized":	ValueOf(pkg.StatusUnauthorized),
		"StatusUnavailableForLegalReasons":	ValueOf(pkg.StatusUnavailableForLegalReasons),
		"StatusUnprocessableEntity":	ValueOf(pkg.StatusUnprocessableEntity),
		"StatusUnsupportedMediaType":	ValueOf(pkg.StatusUnsupportedMediaType),
		"StatusUpgradeRequired":	ValueOf(pkg.StatusUpgradeRequired),
		"StatusUseProxy":	ValueOf(pkg.StatusUseProxy),
		"StatusVariantAlsoNegotiates":	ValueOf(pkg.StatusVariantAlsoNegotiates),
		"StripPrefix":	ValueOf(pkg.StripPrefix),
		"TimeFormat":	ValueOf(pkg.TimeFormat),
		"TimeoutHandler":	ValueOf(pkg.TimeoutHandler),
		"TrailerPrefix":	ValueOf(pkg.TrailerPrefix),
	}
	Types["net/http"] = map[string]Type{
		"Client":	TypeOf((*pkg.Client)(nil)).Elem(),
		"CloseNotifier":	TypeOf((*pkg.CloseNotifier)(nil)).Elem(),
		"ConnState":	TypeOf((*pkg.ConnState)(nil)).Elem(),
		"Cookie":	TypeOf((*pkg.Cookie)(nil)).Elem(),
		"CookieJar":	TypeOf((*pkg.CookieJar)(nil)).Elem(),
		"Dir":	TypeOf((*pkg.Dir)(nil)).Elem(),
		"File":	TypeOf((*pkg.File)(nil)).Elem(),
		"FileSystem":	TypeOf((*pkg.FileSystem)(nil)).Elem(),
		"Flusher":	TypeOf((*pkg.Flusher)(nil)).Elem(),
		"Handler":	TypeOf((*pkg.Handler)(nil)).Elem(),
		"HandlerFunc":	TypeOf((*pkg.HandlerFunc)(nil)).Elem(),
		"Header":	TypeOf((*pkg.Header)(nil)).Elem(),
		"Hijacker":	TypeOf((*pkg.Hijacker)(nil)).Elem(),
		"ProtocolError":	TypeOf((*pkg.ProtocolError)(nil)).Elem(),
		"PushOptions":	TypeOf((*pkg.PushOptions)(nil)).Elem(),
		"Pusher":	TypeOf((*pkg.Pusher)(nil)).Elem(),
		"Request":	TypeOf((*pkg.Request)(nil)).Elem(),
		"Response":	TypeOf((*pkg.Response)(nil)).Elem(),
		"ResponseWriter":	TypeOf((*pkg.ResponseWriter)(nil)).Elem(),
		"RoundTripper":	TypeOf((*pkg.RoundTripper)(nil)).Elem(),
		"ServeMux":	TypeOf((*pkg.ServeMux)(nil)).Elem(),
		"Server":	TypeOf((*pkg.Server)(nil)).Elem(),
		"Transport":	TypeOf((*pkg.Transport)(nil)).Elem(),
	}
}