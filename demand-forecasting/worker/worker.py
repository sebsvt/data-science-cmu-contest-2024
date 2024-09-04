import os
import json
import pandas as pd
import pika
from dotenv import load_dotenv
from eclair.usesage import demand_forecaste_next_day_from_the_last_date

# Load environment variables from .env file
load_dotenv()

# Retrieve RabbitMQ connection details from environment variables
RABBITMQ_HOST = os.getenv("RABBITMQ_HOST")
RABBITMQ_USER = os.getenv("RABBITMQ_USERNAME")
RABBITMQ_PASSWORD = os.getenv("RABBITMQ_PASSWORD")
RABBITMQ_QUEUE = "demand_forecast_tasks"


def connect_to_rabbitmq():
	"""Establish connection to RabbitMQ and declare the queue."""
	credentials = pika.PlainCredentials(RABBITMQ_USER, RABBITMQ_PASSWORD)
	connection = pika.BlockingConnection(
		pika.ConnectionParameters(host=RABBITMQ_HOST, credentials=credentials)
	)
	channel = connection.channel()
	channel.queue_declare(queue=RABBITMQ_QUEUE, durable=True)
	return channel

def do_work(data: dict):
	"""Perform demand forecasting based on received data."""
	try:
		demand_forecast = []
		df = pd.read_csv("./pure_data.csv")
		result = demand_forecaste_next_day_from_the_last_date(
				data=df,
				datetime_col="ds",
				demand_type_col="item_name",
				filter_by='เอกสาร',
				demand="qty"
			)
		demand_forecast.append(result)
		print(demand_forecast)
	except Exception as error:
		print(f"Error: {error}")
		return {"error": str(error)}

def callback(ch, method, properties, body):
	"""Callback function to process messages from the queue."""
	print(f" [x] Received {body.decode()}")
	data = json.loads(body.decode())
	do_work(data)
	print(" [x] Done")
	ch.basic_ack(delivery_tag=method.delivery_tag)

def main():
	"""Main function to start consuming messages."""
	channel = connect_to_rabbitmq()
	print(' [*] Waiting for messages. To exit press CTRL+C')
	channel.basic_qos(prefetch_count=1)
	channel.basic_consume(queue=RABBITMQ_QUEUE, on_message_callback=callback)
	channel.start_consuming()

if __name__ == "__main__":
	main()
