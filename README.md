
# ASCII Art Web Server - Docker Support

![image](https://o365.vn/wp-content/uploads/how-to-export-power-bi-excel.jpeg)

## Description

Ascii-art-web-export is a project that extends the ascii-art-web project by adding the ability to export the generated ASCII art to a file. The web server supports exporting the ASCII art in at least one file format of your choice.

## Authors
- [Hezron Okwach](https://github.com/hezronokwach) 
- [Anne Maina](https://github.com/nyagooh)
- [Brian Bantu](https://github.com/Bantu-art)

## Prerequisites

  Docker installed on your machine. You can download Docker from [here](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://docs.docker.com/engine/install/ubuntu/&ved=2ahUKEwjXrar7oceHAxW2UUEAHUMWFCcQFnoECBUQAQ&usg=AOvVaw2uFia4sMgzkReEqv8xLNZy)

  Go installed on your machine .You can download Go [here](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https://go.dev/doc/install&ved=2ahUKEwi6l_DJoseHAxUEQUEAHSg7JeUQFnoECBQQAQ&usg=AOvVaw0kVh2caUQCSgLfcNTr-PzY)

## Export Functionality
### Features

  Export ASCII Art: The web server provides an endpoint to export the generated ASCII art to a file format (e.g., .txt).

  HTTP Headers: The server includes necessary HTTP headers such as Content-Type, Content-Length, and Content-Disposition for proper file transfer.
  
  User Interface: The web page includes a button or link for users to initiate the file download.
   
  Error Handling: The server handles various HTTP status codes for different scenarios, such as 200 OK, 404 Not Found, 400 Bad Request, and 500 Internal Server Error.

## Implementation Details
The ASCII Art Web Export project is implemented using Go for the web server and follows good coding practices. Key components include:

  HTTP Endpoint: A new endpoint is created to handle file export.

  Export Handler: This handler generates the ASCII art and sends it as a downloadable file with the appropriate HTTP headers.

  Error Handling: Errors are managed appropriately throughout the application, ensuring correct HTTP status codes are returned.


## Docker Support
For those who prefer to run the application using Docker, we provide a Dockerfile and scripts for easy deployment.
## Scripts for Docker

We have provided scripts to build and remove Docker images and containers for this project.  

- Clone the repository:

`` `bash
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
