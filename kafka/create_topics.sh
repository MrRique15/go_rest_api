# create the commands for user to do "make create_topics"

docker exec -it kafka-kafka-1 bash -c "kafka-topics --create --topic order --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1"