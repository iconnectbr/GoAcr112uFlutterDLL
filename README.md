# GoAcr112uFlutterDLL

There was no ACR112U libs for Flutter on Windows.
Just iOS/Android NFC reader libs... So I made one myself which can be used with Flutter FFI.

1)
$Env:CGO_ENABLED = 1

2)
go build -o flutterCardLib.dll -buildmode=c-shared .\main.go

//export is a must
