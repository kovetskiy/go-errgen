package test

//go:generate errgen -type=basicErrMsg
type basicErrMsg struct {
	FooErr error `errmsg:"this is FOO error"`
	BarErr error `errmsg:"this is BAR error [%d, %s]" vars:"hoge int, fuga string"`
}

//go:generate errgen -type=prefixErrMsg -prefix=PREF-
type prefixErrMsg struct {
	BuzErr error `errmsg:"this is BUZ error"`
	QuxErr error `errmsg:"this is QUX error [%d, %s]" vars:"hoge int, fuga string"`
}

//go:generate errgen -type=pathSpecifiedErrMsg -out-file=foobar_errmsg_gen.go
type pathSpecifiedErrMsg struct {
	FooBarErr error `errmsg:"this is FOOBAR error"`
}

//go:generate errgen -type=obsoletedErrMsg
type obsoletedErrMsg struct {
	OneErr   error `errmsg:"this is error 1"`
	TwoErr   error `errmsg:"this is error 2" obsoleted:"true"`
	ThreeErr error `errmsg:"this is error 3"`
}
