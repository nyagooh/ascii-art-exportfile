
# ASCII Art Web Server - Docker Support

![image](https://blog.back4app.com/wp-content/uploads/2023/02/docker-deploy-cover-1.webp)

## Description

Ascii-art-stylize is a project aimed at enhancing the ascii--art-web project by making it more visually appealing, interactive, responsive, and intuitive, while also improving user-friendliness and feedback mechanisms.
Docker Support

We provide Docker support for easy deployment of the ASCII Art Web Stylize project. Follow these steps to build and run the project using Docker.

## Authors
- [Hezron Okwach](https://github.com/hezronokwach) 
- [Anne Maina](https://github.com/nyagooh)
- [Brian Bantu](https://github.com/Bantu-art)

## Prerequisites

  Docker installed on your machine. You can download Docker from [here](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://docs.docker.com/engine/install/ubuntu/&ved=2ahUKEwjXrar7oceHAxW2UUEAHUMWFCcQFnoECBUQAQ&usg=AOvVaw2uFia4sMgzkReEqv8xLNZy)
  Go installed on your machine .You can download Go [here](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://go.dev/doc/install&ved=2ahUKEwi6l_DJoseHAxUEQUEAHSg7JeUQFnoECBQQAQ&usg=AOvVaw0kVh2caUQCSgLfcNTr-PzY)

## Scripts for Docker

We have provided scripts to build and remove Docker images and containers for this project.

   

 - Clone the repository:

```bash
git clone https://learn.zone01kisumu.ke/git/bbantu/ascii-art-web-dockerize
cd ascii-art-web-dockerize
```

 - Build the Docker image:

You can use the provided run.sh script to build the Docker image.

```bash
./run.sh
```

 This script contains:

```bash
#!/bin/bash

docker build -t ascii-art-web-stylize .
docker run -d -p 8060:8060 --name ascii-art-web-stylize-container ascii-art-web-stylize
docker logs ascii-art-web-stylize-container
```
**docker build -t ascii-art-web-stylize** .: This command builds a Docker image named ascii-art-web-stylize from the Dockerfile in the current directory (.).

**docker run -d -p 8060:8060 --name ascii-art-web-stylize-container ascii-art-web-stylize:** This command runs the Docker container in detached mode (-d), mapping port 8060 on your local machine to port 8060 inside the container (-p 8060:8060), and names the container ascii-art-web-stylize-container.

**docker logs ascii-art-web-stylize-container:** This command displays the logs of the running Docker container, allowing you to see output and ensure that the server is running correctly.

 - Access the application:

Open a web browser and visit http://localhost:8060 to use the ASCII Art Web Stylize application.

 - Stopping and Cleaning Up

If you need to stop the Docker container and remove the Docker image, use the provided cleanup.sh script.

```bash
./cleanup.sh
```

This script contains:

```bash
#!/bin/bash
docker stop ascii-art-web-stylize-container
docker rm ascii-art-web-stylize-container
docker rmi ascii-art-web-stylize
```
**docker stop ascii-art-web-stylize-container:** This command stops the running Docker container named ascii-art-web-stylize-container.
**docker rm ascii-art-web-stylize-container:** This command removes the stopped Docker container named ascii-art-web-stylize-container.
**docker rmi ascii-art-web-stylize:** This command removes the Docker image named ascii-art-web-stylize.

## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.
