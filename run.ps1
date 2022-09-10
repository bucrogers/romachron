$URL = "http://localhost:3000"

if ( $args[0] -eq "stop" ) {
    docker-compose down
}
else {
    # Start server, frontend locally
    docker-compose up -d

    # Launch browser on frontend
    Start-Process $URL
}



