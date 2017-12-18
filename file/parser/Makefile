test:	
	#cleaning up previous test containers 
	docker-compose -f ./docker-compose.yml down
	#starting containers
	docker-compose -f ./docker-compose.yml up 
	#cleaning up created containers
	docker-compose -f ./docker-compose.yml down
.PHONY: test