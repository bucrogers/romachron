$URL = "http://localhost:3000"

if ( $args[0] -eq "stop" ) {
    docker-compose down
}
else {
    if ( $args[0] -ne "--build" ) {
        # Pull image if not already present
        docker-compose pull
    }

    # Start server, frontend locally, build from source with layer caching
    $env:DOCKER_TZ="America/New_York" # assume Eastern for now (see run.sh for bash implementation)
    echo $env:DOCKER_TZ
    docker-compose up -d $args[0]

    # Launch browser on frontend
    Start-Process $URL
}



