#!/usr/bin/env python3.6

# copyright 2018 Philipp Hauer
# Source: https://github.com/phauer/blog-related/blob/master/smooth-local-dev-docker/local-dev/mysql-seeding/seed-mysql.py

import random
import time
import getpass

import mysql.connector
from faker import Faker
from mysql.connector import InterfaceError

faker = Faker("en")


class MySqlSeeder:
    def __init__(self):
        config = {
            "user": input("Database username: "),
            "password": getpass.getpass("Database password: "),
            "host": "mysql" if script_runs_within_container() else "localhost",
            "port": "3306",
            "database": "flexo",
        }
        while not hasattr(self, "connection"):
            try:
                self.connection = mysql.connector.connect(**config)
                self.cursor = self.connection.cursor()
            except InterfaceError:
                print("MySQL Container has not started yet. Sleep and retry...")
                time.sleep(1)

    def seed(self):
        print("Clearing old data...")
        self.drop_product_table()
        print("Start seeding...")
        self.create_product_table()
        self.insert_products()

        self.connection.commit()
        self.cursor.close()
        self.connection.close()
        print("Done")

    def create_product_table(self):
        sql = """
        CREATE TABLE products(
          id BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
          created_at DATETIME(3),
          updated_at DATETIME(3),
          deleted_at DATETIME(3),
          name LONGTEXT,
          price DOUBLE,
          description LONGTEXT
        );
        """
        self.cursor.execute(sql)

    def insert_products(self):
        for _ in range(300):
            sql = """
            INSERT INTO products (name, price, description)
            VALUES (%(name)s, %(price)s, %(description)s);
            """
            product_data = {
                "name": faker.first_name(),
                "price": round(random.uniform(1, 100), 2),
                "description": faker.sentence(nb_words=5),
            }
            self.cursor.execute(sql, product_data)

    def drop_product_table(self):
        self.cursor.execute("DROP TABLE IF EXISTS products;")


def script_runs_within_container():
    with open("/proc/1/cgroup", "r") as cgroup_file:
        return "docker" in cgroup_file.read()


MySqlSeeder().seed()
