curl -i -XPOST $DB_HOST/query --data-urlencode "q=CREATE DATABASE $DB_NAME"
curl -i -XPOST $DB_HOST/query --data-urlencode "q=CREATE USER $ADMIN_NAME WITH PASSWORD '$ADMIN_PW' WITH ALL PRIVILEGES"

curl -i -XPOST $DB_HOST/query --data-urlencode "q=CREATE USER $DASHBOARD_NAME WITH PASSWORD '$DASHBOARD_PW'"
curl -i -XPOST $DB_HOST/query --data-urlencode "q=GRANT READ ON $DB_NAME TO $DASHBOARD_NAME"

curl -i -XPOST $DB_HOST/query --data-urlencode "q=CREATE USER $SCRAPER_NAME WITH PASSWORD '$SCRAPER_PW'"
curl -i -XPOST $DB_HOST/query --data-urlencode "q=GRANT READ ON $DB_NAME TO $SCRAPER_NAME"