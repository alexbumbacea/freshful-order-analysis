download:
	go run ./cmd/download_order
denormalize:
	go run ./cmd/denormalize_orders
pivot:
	docker run --rm -v $(shell pwd):/app -w /app -it demisto/pandas:1.0.0.112452 python ./python/pivot-overall.py denormalized.csv report_overall.csv
	docker run --rm -v $(shell pwd):/app -w /app -it demisto/pandas:1.0.0.112452 python ./python/pivot-price-variation.py denormalized.csv report_price.csv