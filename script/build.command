#!/bin/bash

rm -rf bin fyne-cross
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o bin/ikun_windows_amd64.exe

addElement() {
    APP_PATH=$1
    /usr/libexec/PlistBuddy -c "Add :LSUIElement bool true" "$APP_PATH"
    echo "Added LSUIElement to $APP_PATH"
}

# build macOS app
fyne-cross darwin -app-id CF9336AC-2359-41E6-A65B-2E2169A2A5C2 -output ikun -icon ikun.ico -arch arm64
fyne-cross darwin -app-id CF9336AC-2359-41E6-A65B-2E2169A2A5C2 -output ikun -icon ikun.ico -arch amd64


addElement fyne-cross/dist/darwin-arm64/ikun.app/Contents/Info.plist
addElement fyne-cross/dist/darwin-amd64/ikun.app/Contents/Info.plist

hdiutil create -volname ikun_darwin_arm64 -srcfolder fyne-cross/dist/darwin-arm64/ikun.app -ov -format UDZO bin/ikun_darwin_arm64.dmg
hdiutil create -volname ikun_darwin_amd64 -srcfolder fyne-cross/dist/darwin-amd64/ikun.app -ov -format UDZO bin/ikun_darwin_amd64.dmg

