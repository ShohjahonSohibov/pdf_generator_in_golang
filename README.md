# PDF Generator Project

## Overview

The PDF Generator Project is designed to generate PDF documents dynamically based on input data. It provides a simple and efficient way to create PDFs, which can be used for reports, invoices, or any other type of document.

## Features

- Generate PDFs from templates
- Include text, images, and tables
- Support for custom fonts and styles
- Save generated PDFs to the filesystem or send them via email
- RESTful API for generating PDFs programmatically

## Technologies Used

- Go (Golang) for the backend
- Gin for the HTTP server
- MongoDB for storing data
- gofpdf for PDF generation
- Docker for containerization

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:

- Go (1.16 or later)
- MongoDB
- Docker (optional, for containerization)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/pdf-generator.git
   cd pdf-generator


   ```

### How to run

1. Just generating pdf with mock data:

   ```bash
   cd generate_template  
   go run main.go

   ```
2. Generate with dynamic data using HTTP:

   ```bash
   go run main.go  

   ```
