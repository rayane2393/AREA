import discord
from discord.ext import commands
import asyncio
import socket
import json
import threading

IP = "localhost"
PORT = 55555
TOKEN = open("token.txt","r").readline()

client = discord.Client(intents=discord.Intents.default())

def start_socket_server_sync(loop):
    asyncio.run(start_socket_server(loop))

@client.event
async def on_ready():
    print('Bot has started!')
    server_thread = threading.Thread(target=start_socket_server_sync, args=(asyncio.get_running_loop(),))
    server_thread.start()

@client.event
async def on_message(message):
    if message.author == client.user:
        return
    print(f"Received message \"{message.content}\" from {message.author.id}")

async def start_socket_server(loop):
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((IP, PORT))
    server_socket.listen(1)

    while True:
        client_socket, addr = server_socket.accept()
        data = client_socket.recv(1024)
        message = json.loads(data)

        if message['command'] == 'send_message':
            server_name = message['server_name']
            channel_name = message['channel_name']
            content = message['content']
            asyncio.run_coroutine_threadsafe(send_message(server_name, channel_name, content), loop)
        elif message['command'] == 'private_message':
            user_id = message['user_id']
            content = message['content']
            asyncio.run_coroutine_threadsafe(private_message(user_id, content), loop)

def run_async_function_in_new_loop():
    new_loop = asyncio.new_event_loop()
    asyncio.set_event_loop(new_loop)
    new_loop.run_until_complete(start_socket_server())
    new_loop.close()

async def send_message(server_name, channel_name, content):
    for guild in client.guilds:
        if guild.name == server_name:
            for channel in guild.channels:
                if channel.name == channel_name and isinstance(channel, discord.TextChannel):
                    await channel.send(content)
                    return
    print(f"Could not find a channel named \"{channel_name}\" in a server named \"{server_name}\"")

async def private_message(user_id, content):
    print(f"Sending private message \"{content}\" to {user_id}")
    user = await client.fetch_user(user_id)
    await user.send(content)

def run_bot():
    print("Starting Bot ...")
    client.run(TOKEN)

run_bot()
