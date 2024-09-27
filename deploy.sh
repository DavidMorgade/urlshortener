# build with to run build then deploy to codigomorga.es
#
# Usage: ./deploy.sh
#
echo "Building and deploying to codigomorga.es"
go build -o urlshortener

if [ $? -ne 0 ]; then
    echo "Build failed"
    exit 1
fi

echo "Build successful"

USER=morgade
SERVER=codigomorga.es
REMOTE_DIR="/home/proyectos/urlshortener"

echo "Deleting old binary"

ssh $USER@$SERVER "rm -f $REMOTE_DIR/urlshortener"

echo "Deploying to $SERVER"
# Copy the binary to the server
scp urlshortener $USER@$SERVER:$REMOTE_DIR

echo "Deploying static files to $SERVER"

scp -r static $USER@$SERVER:$REMOTE_DIR

if [ $? -ne 0 ]; then
    echo "Deploy failed"
    exit 1
fi

echo "Deploy successful"
