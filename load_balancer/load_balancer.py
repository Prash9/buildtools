import argparse
import logging
import socket


logging.basicConfig(level=logging.DEBUG,
                    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s")

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)

# Create handlers
# console_handler = logging.StreamHandler()
# file_handler = logging.FileHandler('app.log')

# Set the logging level for handlers
# console_handler.setLevel(logging.INFO)
# file_handler.setLevel(logging.DEBUG)

# Create formatters and add them to the handlers
# console_formatter = logging.Formatter('%(name)s - %(levelname)s - %(message)s')
# file_formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')

# console_handler.setFormatter(console_formatter)
# file_handler.setFormatter(file_formatter)

# # Add handlers to the logger
# logger.addHandler(console_handler)
# logger.addHandler(file_handler)


def start_server():
    sock = socket.socket()
    sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    sock.bind(('', args.port))
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    print("Created socket")
    server_socket.listen(5)
    
    while True:
        print("Starting to accept to connection")
        client_socket, client_address = server_socket.accept()
        print("Waiting to accept to connection")
        logging.info(f"{client_address=}; {client_socket=}")
        handle_client(client_socket)


def handle_client(client_socket):
    with client_socket:
        while True:
            # Receive data from the client
            data = client_socket.recv(1024)
            if not data:
                # If no data is received, break the loop
                break
            
            logging.info(f'Received: {data.decode("utf-8")}')
            
            # Echo the data back to the client
            client_socket.sendall({"data":"data"})
            logging.info(f'Sent: {data.decode("utf-8")}')


if __name__ == "__main__":
    
    parser = argparse.ArgumentParser()
    parser.add_argument("--port", help="Port to which run the server", type=int, default=65432)
    args = parser.parse_args()

    # add logger
    # run loop continuously
    # wait for connection
    # accept connection
    print(args)
    start_server()