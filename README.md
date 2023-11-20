# GoAcr112uFlutterDLL

There was no ACR112U libs for Flutter on Windows.
Just iOS/Androif NFC reader libs... So I made one myself with can be used with Flutter FFI.

go build -o flutterCardLib.dll -buildmode=c-shared .\main.go

//export is a must
