import socket
import json

def send_command_to_server(command, user_id, content):
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect(('localhost', 55555))

    message = {
        "command": command,
        "user_id": user_id,
        "content": content,
    }

    client_socket.send(json.dumps(message).encode())
    client_socket.close()

if __name__ == "__main__":
    import sys
    command = sys.argv[1]
    user_id = sys.argv[2]
    content = sys.argv[3]
    send_command_to_server(command, user_id, content)