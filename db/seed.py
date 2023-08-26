import psycopg2
from psycopg2 import sql

# Update these values accordingly
db_host = "postgres-service"  # or the IP of the machine running Docker
db_port = "5432"  # Change to the exposed port if different
db_name = "postgres"
db_user = "postgres"
db_password = "password"


def create_tables():
    # Connect to the PostgreSQL server
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

    # Table creation commands
    users_table_command = (
        """
        CREATE TABLE users (
            user_id UUID PRIMARY KEY,
            username VARCHAR(255) NOT NULL,
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
        print("Successfully created tables")
    except psycopg2.Error as e:
        print("Error:", e)

    # Now populate with data

    print("Finished executing script")
    # Close communication
    cur.close()
    conn.close()
    exit(0)


if __name__ == "__main__":
    create_tables()
