BUILD_PATH="../../build";
BUILD_FILE_NAME="archie";
ORIGIN_ENTRYPOINT="../../main.go";

GOOS=linux GOARCH=amd64 CGO=0 go build -o "$BUILD_PATH/$BUILD_FILE_NAME" $ORIGIN_ENTRYPOINT