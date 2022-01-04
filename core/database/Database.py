"""
THis is the database class which has the core functionality for performing database operations
like adding a new package, updating status of all packages, fetching status of a package etc.
"""
import psycopg2


class Database:

    TABLE_NAME = "ORDERS"
    DB_NAME = "testdb"
    DB_USER = "testdb"
    DB_PWD = "db1234"

    def __init__(self):
        # Creates and open a connection to the database
        self.__conn = psycopg2.connect(database=self.DB_NAME, user=self.DB_USER, password=self.DB_PWD)

    # For use with a with clause
    def __enter__(self):
        return self

    # Close the connection to the database upon exit
    def __exit__(self, exc_type, exc_val, exc_tb):
        self.__conn.close()

    # Add an entry into the table

    # Fetch an entry from the table given a tracking id

    # fetch status of all entries

    # update status of all entries

    # delete an entry given a tracking id

    # clean the table, drop the entire table

