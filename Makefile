.PHONY: deploy

deploy:
	$(MAKE) -C bff build
	$(MAKE) -C deploy apply