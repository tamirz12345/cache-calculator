import argparse
import json
import requests
import time

def send_requests(server_address, requests_file, num_requests):
    with open(requests_file, 'r') as file:
        requests_data = json.load(file)

    num_requests_sent = 0
    num_requests_total = len(requests_data)

    while num_requests_sent < num_requests:
        request_index = num_requests_sent % num_requests_total
        request = requests_data[request_index]

        start_time = time.time()
        response = send_request(server_address, request)
        end_time = time.time()
        elapsed_time = end_time - start_time

        print(f'Response for request {num_requests_sent + 1}: {response.text} Time taken: {elapsed_time:.2f} seconds')

        num_requests_sent += 1

def send_request(server_address, request):
    url = server_address + request['endpoint']
    headers = {'Content-Type': 'application/json'}
    response = requests.post(url, data=json.dumps(request['data']), headers=headers)
    return response

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Send JSON requests to a server for load testing')
    parser.add_argument('--server-address', help='Server address (e.g., http://localhost:8080)')
    parser.add_argument('--requests-file', help='Path to the JSON file containing requests')
    parser.add_argument('--num-requests', type=int, help='Number of requests to send')
    args = parser.parse_args()

    send_requests(args.server_address, args.requests_file, args.num_requests)
