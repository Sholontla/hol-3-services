SERVICES :=  rest_service publisher_service subscriber_service
.PHONY: all $(SERVICES)

all: $(SERVICES)

$(SERVICES):
	cd $@ && docker-compose up -d

