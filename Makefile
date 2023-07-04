up: # Start Solr
	docker-compose up -d
exec:
	docker-compose exec solr bash
down: # Stop Solr
	docker-compose down
ps: # Status Check Solr Container
	docker-compose ps
