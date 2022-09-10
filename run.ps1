$URL = "http://localhost:3000"

if ( $args[0] -eq "stop" ) {
    docker-compose down
}
else {
    # Start server, frontend locally, build from source with layer caching
    docker-compose up -d --build

    # Launch browser on frontend
    Start-Process $URL
}



