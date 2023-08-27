import psycopg2
from psycopg2 import sql
import uuid
import hashlib
import os
import random

# Util functions
def hash_password(password, salt):
    """Hash a password with salt using SHA-256."""
    pwd = password + salt
    return hashlib.sha256(pwd.encode('utf-8')).hexdigest()


def create_tables(conn, cur):
    users_table_command = (
        """
        CREATE TABLE users (
            user_id UUID PRIMARY KEY,
            username VARCHAR(255) UNIQUE NOT NULL,
            salt VARCHAR(255) NOT NULL,
            hashed_password VARCHAR(255) NOT NULL
        );
        """
    )

    products_table_command = (
        """
        CREATE TABLE products (
            product_id UUID PRIMARY KEY,
            product_name VARCHAR(255) NOT NULL,
            product_image_url VARCHAR(1024),
            product_description TEXT,
            product_cost FLOAT(10)
        );
        """
    )

    user_products_table_command = (
        """
        CREATE TABLE reviews (
            review_id SERIAL PRIMARY KEY, 
            user_id UUID,
            product_id UUID,
            review_description TEXT,
            review_date TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(user_id),
            FOREIGN KEY (product_id) REFERENCES products(product_id)
        );
        """
    )

    reaction_prompts_table_command = (
        """
        CREATE TABLE prompts (
            prompt VARCHAR(255) PRIMARY KEY,
            prompt_message TEXT NOT NULL
        );
        """
    )
    # Execute the commands
    try:
        cur.execute(users_table_command)
        cur.execute(products_table_command)
        cur.execute(user_products_table_command)
        cur.execute(reaction_prompts_table_command)
        conn.commit()
    except psycopg2.Error as e:
        conn.rollback()
        print("Error in creating tables:", e)

    print("Finished creating tables ")

def populate_tables(conn, cur):

    # POPULATE PROMPTS TABLE
    prompts_data = [
        ('thumbs_up', 'We\'re happy to hear you had a positive experience! Please leave us a quick review so we can keep doing what we do.'),
        ('thumbs_down', 'We\'re sorry to hear things didnt work out. Please leave us a quick review so we know what to improve.')
    ]

    insert_command = """
        INSERT INTO prompts (prompt, prompt_message) VALUES (%s, %s) ON CONFLICT (prompt) DO NOTHING;
    """

    try:
        # Insert the data
        cur.executemany(insert_command, prompts_data)
        conn.commit()
        print("Successfully populated prompts")
    except psycopg2.Error as e:
        conn.rollback()
        print("Error:", e)


    # POPULATE USERS TABLE
    superuser_name = "admin"
    password = "pass"  
    specific_salt = "SALTSALT10"  # This is the special salt for super user

    hashed_pwd = hash_password(password, specific_salt)
    try:
        insert_command = """
            INSERT INTO users (user_id, username, salt, hashed_password) 
            VALUES (%s, %s, %s, %s);
        """
        cur.execute(insert_command, (str(uuid.uuid4()), superuser_name, specific_salt, hashed_pwd))
        conn.commit()
        print(f"Superuser {superuser_name} inserted (if it didn't exist before)")
    except psycopg2.Error as e:
        print("Error in users :", e)
        conn.rollback()

    # POPULATE PRODUCTS TABLE
    for filename in os.listdir('assets'):
        # Ensure we're only dealing with images (based on extension)
        if filename.lower().endswith(('.png', '.jpg', '.jpeg', '.gif')):
            product_id = str(uuid.uuid4()) # Generate UUID
            name = os.path.splitext(filename)[0]  # Extract name without extension
            image_url = os.path.abspath(os.path.join('assets', filename))
            description = name
            cost = round(random.uniform(1.0, 500.0), 2)  # Generate random float between 1.0 and 500.0 for price
            
            insert_command = """
            INSERT INTO products (product_id, product_name, product_image_url, product_description, product_cost)
            VALUES (%s, %s, %s, %s, %s);
            """
            try:            
                cur.execute(insert_command, (product_id, name, image_url, description, cost))
                print(f"Inserted {name} with path {image_url} into products table.")
            except psycopg2.Error as e:
                print("Error in products :", e)
                conn.rollback()

    conn.commit()



# Update these values accordingly
db_host = "postgres-service"  # or the IP of the machine running Docker
db_port = "5432"  # Change to the exposed port if different
db_name = "postgres"
db_user = "postgres"
db_password = "password"

conn = psycopg2.connect(
    host=db_host,
    port=db_port,  # Specify the port here
    database=db_name,
    user=db_user,
    password=db_password
)
print("Connected to the database!")
# Create a new cursor
cur = conn.cursor()

create_tables(conn, cur)
populate_tables(conn, cur)

cur.close()
conn.close()
exit(0)
