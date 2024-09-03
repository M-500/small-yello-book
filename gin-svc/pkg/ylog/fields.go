package ylog

import "time"

func Error(err error) Field {
	return Field{Key: "error", Val: err}
}

func String(key string, val string) Field {
	return Field{
		Key: key,
		Val: val,
	}
}

func Bool(key string, val bool) Field {
	return Field{
		Key: key,
		Val: val,
	}
}

func Int32(key string, val int32) Field {
	return Field{
		Key: key,
		Val: val,
	}
}

func Int(key string, val int) Field {
	return Field{
		Key: key,
		Val: val,
	}
}

func Duration(key string, val time.Duration) Field {
	return Field{
		Key: key,
		Val: val,
	}
}
