import argparse
import json
import requests
import time
import asyncio

async def send_requests(server_address, num_requests):
    loop = asyncio.get_event_loop()
    tasks = []

    for i in range(num_requests):
        user_id = f"user_{i}"
        request = create_request(user_id)
        task = asyncio.ensure_future(send_request(server_address, request))
        tasks.append(task)

    start_time = time.time()
    responses = await asyncio.gather(*tasks)
    end_time = time.time()
    elapsed_time = end_time - start_time

    for i, response in enumerate(responses):
        print(f"Response for request {i+1}: Time taken: {elapsed_time:.2f} seconds")

    print(f"\nTotal elapsed time: {elapsed_time:.2f} seconds")

async def send_request(server_address, request):
    url = server_address + request['endpoint']
    headers = {'Content-Type': 'application/json'}
    response = await asyncio.get_event_loop().run_in_executor(None, requests.post, url, json.dumps(request['data']), headers)
    return response

def create_request(user_id):
    request = {
        'endpoint': '/multiply',
        'data': {
            'parameters': [1.23, 4.56],
            'userId': user_id,
            'operation': 'Multiply'
        }
    }
    return request

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Send JSON requests to a server for load testing')
    parser.add_argument('--server-address', help='Server address (e.g., http://localhost:8080)')
    parser.add_argument('--num-requests', type=int, help='Number of requests to send')
    args = parser.parse_args()

    loop = asyncio.get_event_loop()
    loop.run_until_complete(send_requests(args.server_address, args.num_requests))
    loop.close()