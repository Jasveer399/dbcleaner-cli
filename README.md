# DBCleaner CLI

[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A powerful and flexible command-line tool for cleaning database tables in PostgreSQL, MySQL, and MongoDB. Get rid of test data or reset your database state with ease.

## Features

- **Multiple Database Support:** Works with PostgreSQL, MySQL, and MongoDB.
- **Interactive Mode:** Interactively select which tables to clean.
- **Safe Operations:**
    - **Dry Run:** See what tables would be affected without actually performing any cleaning operations using the `--dry-run` flag.
    - **Backup:** Automatically back up your data before cleaning with the `--backup` flag.
- **Flexible Cleaning Strategies:**
    - **Truncate or Drop:** Choose to either `TRUNCATE` tables (faster, but doesn't reset auto-increment counters in all DBs) or `DROP` them entirely.
    - **Table Filtering:** Precisely control which tables to include or exclude from the cleaning process.
- **Easy Configuration:** Configure database connections and cleaning options through a simple `dbcleaner.yaml` file.

## Installation

### From source

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/Jasveer399/dbcleaner-cli.git
    cd dbcleaner-cli
    ```

2.  **Build the binary:**
    ```sh
    make build
    ```
    This will create a `dbcleaner` binary in the `bin` directory.

## Usage

The tool is straightforward to use. It revolves around two main commands: `init` and `clean`.

### 1. Initialize Configuration

First, create a configuration file in your project root.

```sh
dbcleaner init
```

This command generates a `dbcleaner.yaml` file with default settings.

### 2. Customize Configuration

Open `dbcleaner.yaml` and edit it to match your database credentials and cleaning preferences.

```yaml
database:
  driver: postgres # postgres, mysql, or mongodb
  host: localhost
  port: 5432
  user: admin
  password: adminpassword
  dbname: my_database
  sslmode: disable

cleaner:
  # A list of tables to never clean
  exclude_tables:
    - schema_migrations
    - users

  # A list of tables to exclusively clean (if not empty)
  include_tables: []

  # Set to true to TRUNCATE tables instead of dropping them
  truncate_only: false
```

### 3. Clean the Database

Run the `clean` command to start the process.

```sh
dbcleaner clean
```

The tool will connect to the database, identify the tables, and prompt you to select which ones you want to clean.

#### Command-Line Flags

You can override the configuration file settings using flags:

-   `--config`: Specify a different configuration file path.
-   `--dry-run`: Show which tables would be cleaned without executing the operation.
-   `--truncate`: Truncate tables instead of dropping them.
-   `--backup`: Perform a backup before cleaning.

**Example:**

```sh
# Perform a dry run, truncating tables specified in a custom config file
dbcleaner clean --config=./dev.yaml --dry-run --truncate
```

## Development

To contribute to the development of `dbcleaner-cli`:

1.  Clone the repository.
2.  Install dependencies: `go mod tidy`
3.  Build the project: `make build`
4.  Run tests: `go test ./...`

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.