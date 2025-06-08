// // We will use Docker for installing

// docker run -it -p 8080:80 ubuntu
// apt-get update
// apt-get install nginx

// http://localhost:8080

// nginx.conf file is located at /etc/nginx/nginx.conf and it contains the configuration details of how you want to configure nginx as a web ServiceWorkerRegistration

// we need to reload ngin after every change
// nginx -s reload

// since nginx is a webserver, it listens to a port. By default, it listens to port 80. We can change the port by changing the configuration file.

/*
    events {

    }
    // -> we are defining the event block. This block is used to define the number of connections that can be handled by the server.

    // -> we are creating a http block. This block is used to define the server block.
    // -> there can be multiple server blocks in the http block. Each server block can have a different configuration. multiple servers listening to different ports.
    http {
        server {
            listen 80;
            // server_name localhost;
            server_name _; // server_name is used to define the domain name of the server. _ means all the domain names, basically specifying which types of requests to handle
            location / {
                return 200 "Hello, World!";
            }
        }
    }

*/


// If not from docker, we can download nginx from the official website and install it.

// C:\Users\HP\Downloads\nginx-1.26.2\nginx-1.26.2
// open this path in cmd and run "start nginx" to start the server -> this will start nginx in the background
// http://localhost -> nginx will be accesible here, default port 80

// you will have to change the original config file only every time and reload the server to see the changes
// "nginx -s reload"

// to stop the server, run "nginx -s stop"