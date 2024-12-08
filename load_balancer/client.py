import socket

def client(host='127.0.0.1', port=65432):
    # Create a socket object
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as client_socket:
        try:
            # Connect to the server
            client_socket.connect((host, port))
            print(f'Connected to {host}:{port}')
            
            # Send a message to the server
            message = 'Hello, Server'
            client_socket.sendall(message.encode('utf-8'))
            print(f'Sent: {message}')
            
            # Receive a response from the server
            response = client_socket.recv(1024)
            print(f'Received from server: {response.decode("utf-8")}')
        
        except ConnectionRefusedError:
            print(f'Connection refused. Make sure the server is running on {host}:{port}.')
        except Exception as e:
            print(f'An error occurred: {e}')

if __name__ == '__main__':
    # Change the host and port if necessary
    client(host='127.0.0.1', port=3000)
