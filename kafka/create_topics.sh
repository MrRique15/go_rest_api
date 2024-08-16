# create the commands for user to do "make create_topics"

docker exec -it kafka-kafka-1 bash -c "kafka-topics --create --topic order --bootstrap-server localhost:29092 --partitions 1 --replication-factor 1 && kafka-topics --create --topic payment --bootstrap-server localhost:29092 --partitions 1 --replication-factor 1"