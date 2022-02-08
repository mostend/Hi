#!/bin/sh

APP="Hi.app"
mkdir -p $APP/Contents/{MacOS,Resources}
go build -o $APP/Contents/MacOS/Hi
cat > $APP/Contents/Info.plist << EOF
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
	  <key>CFBundlePackageType</key><string>APPL</string>
	  <key>CFBundleName</key><string>Hi</string>
	  <key>CFBundleExecutable</key><string>Hi</string>
	  <key>CFBundleVersion</key><string>1.0.0</string>
	  <key>CFBundleGetInfoString</key><string>Built</string>
	  <key>CFBundleShortVersionString</key><string>1.0.0</string>
	  <key>CFBundleIconFile</key><string>iconfile</string>
	  <key>LSMinimumSystemVersion</key><string>10.13.0</string>
	  <key>NSHighResolutionCapable</key><string>true</string>
	  <key>NSHumanReadableCopyright</key><string>Copyright Hi,World</string>
	  <key>LSUIElement</key><string>1</string>
    </dict>
</plist>
EOF
cp icon/iconfile.icns $APP/Contents/Resources/iconfile.icns
find $APP