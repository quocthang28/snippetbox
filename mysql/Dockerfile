# Use an official MySQL image as a base image
FROM mysql:latest

# Set the root password for MySQL (change 'password' to your desired password)
ENV MYSQL_ROOT_PASSWORD=password

# Create a database (change 'mydatabase' to your desired database name)
ENV MYSQL_DATABASE=snippetbox

# Specify the character set and collation for the database (optional)
ENV MYSQL_CHARSET=utf8
ENV MYSQL_COLLATION=utf8_general_ci

# Expose the default MySQL port
EXPOSE 3306

# Specify a volume for MySQL data persistence
VOLUME /var/lib/mysql
