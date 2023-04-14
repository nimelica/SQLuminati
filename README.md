## SQLuminati

Sick of staring at boring SQLite table results in the terminal? Give your eyes a treat with SQLuminati, a CLI tool that makes your data beautiful and readable with vibrant colors and neat formatting

### Pre-Requirement
- golang > 1.9

### Setup
- Go to the directory where you have your SQLite3 database file
- Clone the SQLuminati repository
- Navigate into the SQLuminati directory 
```cd SQLuminati```
- Run the Makefile
```make```

### Usage
- Start the program by running: ./draw <database-file-path> <query>

### Result
```
SQLuminati % draw ../fuel.db "SELECT * FROM user_credentials"
+----+----------+--------------------------------------------------------------------------------------------------------+
| ID | USERNAME |                                             PASSWORD HASH                                              |
+----+----------+--------------------------------------------------------------------------------------------------------+
|  1 | nim      | pbkdf2:sha256:26111$SgGRWn63xrXUzJkhW$1a4bef1b0140a9a83bj8ej09bj6079a99e2a8330c550e21e4444b9ae6d4d3b7b |
+----+----------+--------------------------------------------------------------------------------------------------------+
```


