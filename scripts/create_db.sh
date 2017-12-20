# This script handles local database creation to make testing easier.
# Run with `make db`
set -x

postgres -V
if [ "$?" -gt "0" ]; then
  echo "MAKE: Postgres not installed! Would you like to install it? (y/n)"
  echo "MAKE: Note that you must have Homebrew installed for this to work."
  read choice
  if [ "$choice" == "y" ]; then
    echo "MAKE: Installing postgres..."
    brew install postgresql
    echo "MAKE: Process complete - please run 'make db' again."
    exit
  else
    echo "MAKE: Alright, bye!"
    exit
  fi
fi

echo "MAKE: Killing existing postgres processes..."
pkill postgres
pg_ctl -D /usr/local/var/postgres start
createuser -s postgres
createdb test_db
psql -d test_db -a -f ./scripts/test_db_setup.sql
echo "MAKE: Local database ready to go!"