# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Host Information
    # GET https://api.shodan.io/shodan/host/1.1.1.1

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/1.1.1.1",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Search Shodan without Results
    # GET https://api.shodan.io/shodan/host/count

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/count",
            params={
                "query": "port:22",
                "facets": "org, os",
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Search Shodan
    # GET https://api.shodan.io/shodan/host/search

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/search",
            params={
                "query": "product:nginx",
                "facets": "org, os",
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List all search facets
    # GET https://api.shodan.io/shodan/host/search/facets

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/search/facets",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List all filters that can be used when searching
    # GET https://api.shodan.io/shodan/host/search/filters

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/search/filters",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Break the search query into tokens
    # GET https://api.shodan.io/shodan/host/search/tokens

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/host/search/tokens",
            params={
                "key": "your_api_key",
                "query": "product:nginx",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List all ports that Shodan is crawling on the Internet
    # GET https://api.shodan.io/shodan/ports

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/ports",
            params={
                "query": "product:nginx",
                "facets": "org, os",
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List all protocols that can be used when performing on-demand Internet scans via Shodan
    # GET https://api.shodan.io/shodan/protocols

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/protocols",
            params={
                "query": "product:nginx",
                "facets": "org, os",
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Request Shodan to crawl an IP/ netblock Internet scans via Shodan Duplicate
    # POST https://api.shodan.io/shodan/scan

    try:
        response = requests.post(
            url="https://api.shodan.io/shodan/scan",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
                "ips": "1.1.1.1,8.8.8.8",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Crawl the Internet for a specific port and protocol using Shodan Duplicate (2)
    # POST https://api.shodan.io/shodan/scan/internet

    try:
        response = requests.post(
            url="https://api.shodan.io/shodan/scan/internet",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
                "port": "22",
                "protocol": "http",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Get list of all the created scans
    # GET https://api.shodan.io/shodan/scans

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/scans",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Get the status of a scan request
    # GET https://api.shodan.io/shodan/scan/Lxu8uAEahnlMEm9r

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/scan/Lxu8uAEahnlMEm9r",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests
import json


def send_request():
    # Create an alert to monitor a network range
    # POST https://api.shodan.io/shodan/alert

    try:
        response = requests.post(
            url="https://api.shodan.io/shodan/alert",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
            data=json.dumps({
                "name": "DNS Alert",
                "filters": {
                    "ip": [
                        "8.8.8.8",
                        "1.1.1.1"
                    ]
                },
                "expires": "0"
            })
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests
import json


def send_request():
    # Get a list of all the created alerts
    # GET https://api.shodan.io/shodan/alert/info

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/alert/info",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
            data=json.dumps({
                "name": "DNS Alert",
                "filters": {
                    "ip": [
                        "8.8.8.8",
                        "1.1.1.1"
                    ]
                },
                "expires": "0"
            })
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Get the details for a network alert
    # GET https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/info

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/info",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Delete an alert
    # DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3

    try:
        response = requests.delete(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests
import json


def send_request():
    # Edit the networks monitored in an alert
    # POST https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3

    try:
        response = requests.post(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
            data=json.dumps({
                "filters": {
                    "ip": "1.1.1.1"
                }
            })
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Get a list of available triggers
    # GET https://api.shodan.io/shodan/alert/triggers

    try:
        response = requests.get(
            url="https://api.shodan.io/shodan/alert/triggers",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Enable a trigger
    # PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware

    try:
        response = requests.put(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Disable a trigger
    # PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware

    try:
        response = requests.put(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Add to Whitelist
    # PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53

    try:
        response = requests.put(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Remove from Whitelist
    # DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53

    try:
        response = requests.delete(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Add the notifier to the alert
    # PUT https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default

    try:
        response = requests.put(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Remove the notifier from the alert
    # DELETE https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default

    try:
        response = requests.delete(
            url="https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/json",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List all user-created notifiers
    # GET https://api.shodan.io/notifier

    try:
        response = requests.get(
            url="https://api.shodan.io/notifier",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # List of available notification providers
    # GET https://api.shodan.io/notifier/provider

    try:
        response = requests.get(
            url="https://api.shodan.io/notifier/provider",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Create a new notification service for the user
    # POST https://api.shodan.io/notifier

    try:
        response = requests.post(
            url="https://api.shodan.io/notifier",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
                "provider": "email",
                "to": "test@google.de",
                "description": "\"email notification\"",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Delete a notification service
    # DELETE https://api.shodan.io/notifier/ZKJpffUnarJushBf

    try:
        response = requests.delete(
            url="https://api.shodan.io/notifier/ZKJpffUnarJushBf",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Get information about a notifier
    # GET https://api.shodan.io/notifier/ZKJpffUnarJushBf

    try:
        response = requests.get(
            url="https://api.shodan.io/notifier/ZKJpffUnarJushBf",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Edit a notifier
    # GET https://api.shodan.io/notifier/ZKJpffUnarJushBf

    try:
        response = requests.get(
            url="https://api.shodan.io/notifier/ZKJpffUnarJushBf",
            params={
                "key": "your_api_key",
            },
            headers={
                "Content-Type": "application/x-www-form-urlencoded; charset=utf-8",
            },
            data={
                "to": "test@google.de",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Domain Information
    # GET https://api.shodan.io/dns/domain/google.de

    try:
        response = requests.get(
            url="https://api.shodan.io/dns/domain/google.de",
            params={
                "key": "your_api_key",
                "history": "False",
                "page": "1",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # DNS Lookup
    # GET https://api.shodan.io/dns/resolve

    try:
        response = requests.get(
            url="https://api.shodan.io/dns/resolve",
            params={
                "key": "your_api_key",
                "hostnames": "google.com,bing.com",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # Reverse DNS Lookup
    # GET https://api.shodan.io/dns/reverse

    try:
        response = requests.get(
            url="https://api.shodan.io/dns/reverse",
            params={
                "key": "your_api_key",
                "ips": "1.1.1.1,8.8.8.8",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # HTTP Headers
    # GET https://api.shodan.io/tools/httpheaders

    try:
        response = requests.get(
            url="https://api.shodan.io/tools/httpheaders",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


# Install the Python Requests library:
# `pip install requests`

import requests


def send_request():
    # My IP Address
    # GET https://api.shodan.io/tools/myip

    try:
        response = requests.get(
            url="https://api.shodan.io/tools/myip",
            params={
                "key": "your_api_key",
            },
        )
        print('Response HTTP Status Code: {status_code}'.format(
            status_code=response.status_code))
        print('Response HTTP Response Body: {content}'.format(
            content=response.content))
    except requests.exceptions.RequestException:
        print('HTTP Request failed')


