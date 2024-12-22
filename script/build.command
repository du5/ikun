#!/bin/bash

addElement() {
    APP_PATH=$1
    /usr/libexec/PlistBuddy -c "Add :LSUIElement bool true" "$APP_PATH"
    echo "Added LSUIElement to $APP_PATH"
}

# build macOS app
fyne-cross darwin -app-id CF9336AC-2359-41E6-A65B-2E2169A2A5C2 -output ikun -icon 8c0e2877eb06f6bd84bfba168cef11df997787e7.jpg -arch arm64
fyne-cross darwin -app-id CF9336AC-2359-41E6-A65B-2E2169A2A5C2 -output ikun -icon 8c0e2877eb06f6bd84bfba168cef11df997787e7.jpg -arch amd64


addElement fyne-cross/dist/darwin-arm64/ikun.app/Contents/Info.plist
addElement fyne-cross/dist/darwin-amd64/ikun.app/Contents/Info.plist

zip -9 bin/ikun_darwin_arm64.zip fyne-cross/dist/darwin-arm64/ikun.app
zip -9 bin/ikun_darwin_amd64.zip fyne-cross/dist/darwin-amd64/ikun.app

rm -rf fyne-cross

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -o bin/ikun_windows_amd64.exe
