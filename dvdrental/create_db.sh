psql -U postgres -c 'create database dvdrental'


echo "import dvd rental data..."
pg_restore -U postgres -d dvdrental /tmp/dvdrental.tar