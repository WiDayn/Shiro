# Shiro - A Lightweight Go Blogging System

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![Iris](https://img.shields.io/badge/Iris-5C2D91?style=for-the-badge) ![MySQL](https://img.shields.io/badge/MySQL-4479A1?style=for-the-badge&logo=mysql&logoColor=white) ![Bootstrap](https://img.shields.io/badge/Bootstrap-7952B3?style=for-the-badge&logo=bootstrap&logoColor=white)


Shiro is a lightweight blog system built with Go, leveraging the Iris framework and MySQL for data persistence through go-sql-driver. Designed to be lightweight, simple, low on resource usage, and quick to deploy, Shiro is perfect for personal blogs, technical sharing platforms, and more.

## Key Features

- **Lightweight Design**: Shiro aims to provide a light blog solution suitable for applications of all sizes.
- **Simple User Interface**: Offers a clean and simple interface to enhance reading experience.
- **Low Resource Consumption**: Optimized for low resource usage across various environments.
- **Quick Deployment**: Simplifies the installation process for quick deployment across various platforms.
- **Template System**: Uses a jet-based template system for easy customization, allowing effortless modifications and replacements of the interface style.

## Technology Stack

- **[Go](https://golang.org/)**: An efficient programming language ensuring application performance.
- **[Iris](https://www.iris-go.com/)**: A high-performance Go web framework providing a simple API and powerful features.
- **[go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)**: Stable data storage implemented with go-sql-driver/mysql.
- **[Jet Template Engine](https://github.com/CloudyKit/jet)**: A flexible template engine for easy customization and template modification.
- **[Bootstrap](https://github.com/twbs/bootstrap)**: Sleek, intuitive, and powerful front-end framework for faster and easier web development.

## Quick Start

For users who prefer not to build from source, pre-compiled versions are available for both Linux and Windows platforms. These can be directly downloaded from the [Releases](https://github.com/your-username/Shiro/releases) section of the Shiro repository.

#### Using Pre-compiled Binaries

1. Download the latest release for your platform (Linux or Windows).
2. Extract the contents of the zip file.
3. Open a command line interface in the extracted directory.
4. Run Shiro using the following command:
    - For Linux:
      ```sh
      ./shiro
      ```
    - For Windows:
      ```cmd
      shiro.exe
      ```

This will start the Shiro blogging system, and you can access it by visiting `http://localhost:8080`.

### Building from Source
### Prerequisites

- Go 1.2x
- MySQL 5.7 or higher

### Installation Steps

1. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/WiDayn/Shiro.git
   ```
2. Enter the project directory:
   ```sh
   cd Shiro
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Configure your database connection (edit `config.json` or set environment variables).

5. Start the project:
   ```sh
   go run internal/cmd
   or
   go build internal/cmd
   ```

Access `http://localhost:8080` to enjoy your Shiro blogging system!

## Contributing

Shiro welcomes contributions of any kind, whether they are feature requests, bug reports, or code contributions. Please read the contributing guide first, then submit your pull requests.

## License

Shiro is released under the MIT License. For more information, please refer to the LICENSE file.

---

Happy blogging with Shiro!