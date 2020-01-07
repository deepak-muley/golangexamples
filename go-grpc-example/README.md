How to install protoc?
#https://developers.google.com/protocol-buffers/docs/gotutorial

#Mac
<pre>
brew install protobuf
brew install protoc-gen-go
</pre>

#Linux
<pre>
PROTOC_ZIP=protoc-3.4.0-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

go get -u github.com/golang/protobuf/protoc-gen-go
</pre>
