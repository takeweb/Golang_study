crawler % sqlite3 data.sqlite3
SQLite version 3.32.3 2020-06-18 14:16:19
Enter ".help" for usage hints.
sqlite> CREATE TABLE md_data (
   ...> id INTEGER PRIMARY KEY AUTOINCREMENT,
   ...> title TEXT NOT NULL,
   ...> url TEXT NOT NULL,
   ...> markdown TEXT
   ...> );
sqlite> .exit

