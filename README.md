# LaraGo
Like Laravel for PHP <br/>

LaraGo is a Go-based framework inspired by Laravel, offering a structured approach to web application development. It includes features like routing, middleware, dependency injection, and more.  

This repository consists of two main components:  

1. **Lara Module** – The core framework providing Laravel-like functionalities in Go.  
2. **test-app** – A sample project that demonstrates how to use the Lara module.  

## Features  

### 1. Routing System  

LaraGo includes a powerful and flexible routing system similar to Laravel's. It allows developers to define routes for handling HTTP requests efficiently.  

### 2. Middleware Support  

Middleware enables request processing before they reach the controller logic, allowing functionalities such as authentication, logging, and data transformation.  

---

## **Filesystem Support**  

LaraGo supports multiple storage backends, making it easier to handle file operations across different environments. Supported filesystems include:  

- **Local Filesystem** – Allows reading from and writing to files on the server using Go's built-in filesystem operations.  
- **Amazon S3** – Enables cloud storage capabilities via the AWS SDK for Go, supporting scalable file storage.  
- **WebDAV** – Provides integration with WebDAV-based storage services, allowing remote file management.  
- **FTP** – Supports file uploads and downloads over FTP, making it compatible with traditional hosting environments.  

Developers can extend LaraGo to integrate additional storage options as needed.  

---

## **Mailing Support**  

LaraGo allows sending emails through different mail transport protocols. The framework does not include built-in mailing services but supports:  

- **SMTP** – Sending emails via standard mail servers using Go's `net/smtp` package.  
- **Gomail** – A more feature-rich third-party library for email handling, supporting attachments and HTML email formats.  

Developers can easily integrate mailing functionalities into their applications by configuring SMTP servers or third-party email APIs.  

---

## **Installation**  

To install and use LaraGo, developers can retrieve the package via Go modules.  

---

## **Running the Example Application**  

The provided test application demonstrates how to use the Lara module. Running the application launches a web server and showcases the core functionalities.  

---

## **Contributing**  

Contributions are welcome! Developers can contribute by submitting pull requests, reporting issues, or suggesting improvements.  

---

## **License**  

LaraGo is released under the [GNU General Public License v3.0](LICENSE)

---
