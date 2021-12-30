## Host Information
# Returns all services that have been found on the given host IP
curl "https://api.shodan.io/shodan/host/1.1.1.1?key=your_api_key"

## Search Shodan without Results
# This method behaves identical to "/shodan/host/search" with the only difference that this method does not return any host results, it only returns the total number of results that matched the query and any facet information that was requested. As a result this method does not consume query credits.
# 
# Parameters
# query: [String] Shodan search query. The provided string is used to search the database of banners in Shodan, with the additional option to provide filters inside the search query using a "filter:value" format. For example, the following search query would find Apache Web servers located in Germany: "apache country:DE".
# https://beta.shodan.io/search/filters
# 
# facets (optional): [String] A comma-separated list of properties to get summary information on. Property names can also be in the format of "property:count", where "count" is the number of facets that will be returned for a property (i.e. "country:100" to get the top 100 countries for a search query). Visit the Shodan website's Facet Analysis page for an up-to-date list of available facets:
# https://beta.shodan.io/search/facet
# 
curl "https://api.shodan.io/shodan/host/count?query=port:22&facets=org,%20os&key=your_api_key"

## Search Shodan
# Search Shodan using the same query syntax as the website and use facets to get summary information for different properties.
# 
# Parameters
# query: [String] Shodan search query. The provided string is used to search the database of banners in Shodan, with the additional option to provide filters inside the search query using a "filter:value" format. For example, the following search query would find Apache Web servers located in Germany: "apache country:DE".
# https://beta.shodan.io/search/filters
# 
# facets (optional): [String] A comma-separated list of properties to get summary information on. Property names can also be in the format of "property:count", where "count" is the number of facets that will be returned for a property (i.e. "country:100" to get the top 100 countries for a search query). Visit the Shodan website's Facet Analysis page for an up-to-date list of available facets:
# https://beta.shodan.io/search/facet
# 
curl "https://api.shodan.io/shodan/host/search?query=product:nginx&facets=org,%20os&key=your_api_key"

## List all search facets
# This method returns a list of facets that can be used to get a breakdown of the top values for a property.
# 
# 
curl "https://api.shodan.io/shodan/host/search/facets?key=your_api_key"

## List all filters that can be used when searching
# This method returns a list of search filters that can be used in the search query.
# 
# 
curl "https://api.shodan.io/shodan/host/search/filters?key=your_api_key"

## Break the search query into tokens
# This method lets you determine which filters are being used by the query string and what parameters were provided to the filters.
# 
# Parameters
# query: [String] Shodan search query. The provided string is used to search the database of banners in Shodan, with the additional option to provide filters inside the search query using a "filter:value" format. For example, the following search query would find Apache Web servers located in Germany: "apache country:DE".
# https://beta.shodan.io/search/filters
curl "https://api.shodan.io/shodan/host/search/tokens?key=your_api_key&query=product:nginx"

## List all ports that Shodan is crawling on the Internet
# This method returns a list of port numbers that the crawlers are looking for.
curl "https://api.shodan.io/shodan/ports?query=product:nginx&facets=org,%20os&key=your_api_key"

## List all protocols that can be used when performing on-demand Internet scans via Shodan
# This method returns a list of port numbers that the crawlers are looking for.
curl "https://api.shodan.io/shodan/protocols?query=product:nginx&facets=org,%20os&key=your_api_key"

## Request Shodan to crawl an IP/ netblock Internet scans via Shodan Duplicate
# Use this method to request Shodan to crawl a network.
# 
# Requirements
# This method uses API scan credits: 1 IP consumes 1 scan credit. You must have a paid API plan (either one-time payment or subscription) in order to use this method.
# 
# Parameters
# ips: [String] A comma-separated list of IPs or netblocks (in CIDR notation) that should get crawled.
# service: [Array] A list of services that should get scanned, where a service is defined as a [port, protocol].
# 
curl -X "POST" "https://api.shodan.io/shodan/scan?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Crawl the Internet for a specific port and protocol using Shodan Duplicate (2)
# Use this method to request Shodan to crawl the Internet for a specific port.
# 
# Requirements
# This method is restricted to security researchers and companies with a Shodan Enterprise Data license. To apply for access to this method as a researcher, please email jmath@shodan.io with information about your project. Access is restricted to prevent abuse.
# 
# Parameters
# port: [Integer] The port that Shodan should crawl the Internet for.
# protocol: [String] The name of the protocol that should be used to interrogate the port. See /shodan/protocols for a list of supported protocols.
# 
curl -X "POST" "https://api.shodan.io/shodan/scan/internet?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Get list of all the created scans
# Returns a listing of all the on-demand scans that are currently active on the account.
curl "https://api.shodan.io/shodan/scans?key=your_api_key"

## Get the status of a scan request
# Check the progress of a previously submitted scan request. Possible values for the status are:
# SUBMITTING
# QUEUE
# PROCESSING
# DONE
# 
curl "https://api.shodan.io/shodan/scan/Lxu8uAEahnlMEm9r?key=your_api_key"

## Create an alert to monitor a network range
# Use this method to create a network alert for a defined IP/ netblock which can be used to subscribe to changes/ events that are discovered within that range.
# 
# Parameters
# The alert is created by sending a JSON encoded object that has the structure:
# 
# {
#     "name": {name},
#     "filters": {
#         "ip": {ip},
#     },
#     "expires": {expires},
# }
# 
# name: [String] The name to describe the network alert.
# filters: [Object] An object specifying the criteria that an alert should trigger. The only supported option at the moment is the "ip" filter.
# filters.ip: [String] A list of IPs or network ranges defined using CIDR notation.
# expires (optional): [Integer] Number of seconds that the alert should be active.
curl -X "POST" "https://api.shodan.io/shodan/alert?key=your_api_key" \
     -H 'Content-Type: application/json' \
     -d $'{
  "name": "DNS Alert",
  "filters": {
    "ip": [
      "8.8.8.8",
      "1.1.1.1"
    ]
  },
  "expires": "0"
}'

## Get a list of all the created alerts
# Returns a listing of all the network alerts that are currently active on the account.
# 
curl "https://api.shodan.io/shodan/alert/info?key=your_api_key" \
     -H 'Content-Type: application/json' \
     -d $'{
  "name": "DNS Alert",
  "filters": {
    "ip": [
      "8.8.8.8",
      "1.1.1.1"
    ]
  },
  "expires": "0"
}'

## Get the details for a network alert
# Returns the information about a specific network alert.
# 
# Parameters
# id: [String] Alert ID
curl "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/info?key=your_api_key" \
     -H 'Content-Type: application/json'

## Delete an alert
# Remove the specified network alert.
# 
# 
# Parameters
# id: [String] Alert ID
curl -X "DELETE" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key" \
     -H 'Content-Type: application/json'

## Edit the networks monitored in an alert
# Use this method to edit a network alert with a new list of IPs/ networks to keep track of.
# 
# Parameters
# The alert is edited by sending a JSON encoded object that has the structure:
# 
# {
#     "filters": {
#         "ip": {ip},
#     },
# }
# 
# filters: [Object] An object specifying the criteria that an alert should trigger. The only supported option at the moment is the "ip" filter.
# filters.ip: [String] A list of IPs or network ranges defined using CIDR notation.
curl -X "POST" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3?key=your_api_key" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filters": {
    "ip": "1.1.1.1"
  }
}'

## Get a list of available triggers
# Returns a list of all the triggers that can be enabled on network alerts.
curl "https://api.shodan.io/shodan/alert/triggers?key=your_api_key" \
     -H 'Content-Type: application/json'

## Enable a trigger
# Get notifications when the specified trigger is met.
# 
# Parameters
# id: [String] Alert ID
# trigger: [String] Comma-separated list of trigger names
curl -X "PUT" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key" \
     -H 'Content-Type: application/json'

## Disable a trigger
# Stop getting notifications for the specified trigger.
# 
# Parameters
# id: [String] Alert ID
# trigger: [String] Comma-separated list of trigger names
curl -X "PUT" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware?key=your_api_key" \
     -H 'Content-Type: application/json'

## Add to Whitelist
# Ignore the specified service when it is matched for the trigger.
# 
# Parameters
# id: [String] Alert ID
# trigger: [String] Comma-separated list of trigger names
# service: [String] Service specified in the format "ip:port" (ex. "1.1.1.1:80"
curl -X "PUT" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key" \
     -H 'Content-Type: application/json'

## Remove from Whitelist
# Start getting notifications again for the specified trigger.
# 
# Parameters
# id: [String] Alert ID
# trigger: [String] Comma-separated list of trigger names
# service: [String] Service specified in the format "ip:port" (ex. "1.1.1.1:80"
curl -X "DELETE" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/trigger/malware/ignore/1.1.1.1:53?key=your_api_key" \
     -H 'Content-Type: application/json'

## Add the notifier to the alert
# Add the specified notifier to the network alert. Notifications are only sent if triggers have also been enabled. For each created user, there is a default notifier which will sent via email.
# 
# Parameters
# id: [String] Alert ID
# notifier_id: [String] Notifier ID
curl -X "PUT" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key" \
     -H 'Content-Type: application/json'

## Remove the notifier from the alert
# Remove the notification service from the alert. Notifications are only sent if triggers have also been enabled.
# 
# Parameters
# id: [String] Alert ID
# notifier_id: [String] Notifier ID
curl -X "DELETE" "https://api.shodan.io/shodan/alert/JRH1XDBUA72I9UA3/notifier/default?key=your_api_key" \
     -H 'Content-Type: application/json'

## List all user-created notifiers
# Get a list of all the notifiers that the user has created.
curl "https://api.shodan.io/notifier?key=your_api_key"

## List of available notification providers
# Get a list of all the notification providers that are available and the parameters to submit when creating them.
curl "https://api.shodan.io/notifier/provider?key=your_api_key"

## Create a new notification service for the user
# Use this method to create a new notification service endpoint that Shodan services can send notifications through.
# 
# Parameters
# The parameters depend on the type of notification service that is being created. To get a list of parameters for a provider us the /notifier/provider endpoint. The following parameters always need to be provided:
# provider: [String] Provider name as returned by /notifier/provider
# description: [String] Description of the notifier
# **args: [String] Arguments required by the provider
curl -X "POST" "https://api.shodan.io/notifier?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Delete a notification service
# Remove the notification service created for the user.
# 
# Parameters
# id: [String] Notifier ID
curl -X "DELETE" "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Get information about a notifier
# Use this method to create a new notification service endpoint that Shodan services can send notifications through.
# 
# Parameters
# id: [String] Notifier ID
# 
curl "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Edit a notifier
# Use this method to update the parameters of a notifier.
# 
# Parameters
# The parameters depend on the type of notification service that is being created. To get a list of parameters for a provider us the /notifier/provider endpoint.
# 
# id: [String] Notifier ID
# **args: [String] Arguments required by the provider
curl "https://api.shodan.io/notifier/ZKJpffUnarJushBf?key=your_api_key" \
     -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8'

## Domain Information
# Get all the subdomains and other DNS entries for the given domain. Uses 1 query credit per lookup.
# 
# Parameters
# domain: [String] Domain name to lookup; example "cnn.com"
# history (optional): [Boolean] True if historical DNS data should be included in the results (default: False)
# type (optional): [String] DNS type, possible values are: A, AAAA, CNAME, NS, SOA, MX, TXT
# page (optional): [Integer] The page number to page through results 100 at a time (default: 1)
curl "https://api.shodan.io/dns/domain/google.de?key=your_api_key&history=False&page=1"

## DNS Lookup
# Look up the IP address for the provided list of hostnames.
# 
# Parameters
# hostnames: [String] Comma-separated list of hostnames; example "google.com,bing.com"
curl "https://api.shodan.io/dns/resolve?key=your_api_key&hostnames=google.com,bing.com"

## Reverse DNS Lookup
# Look up the hostnames that have been defined for the given list of IP addresses.
# 
# Parameters
# ips: [String] Comma-separated list of IP addresses; example "74.125.227.230,204.79.197.200"
curl "https://api.shodan.io/dns/reverse?key=your_api_key&ips=1.1.1.1,8.8.8.8"

## HTTP Headers
# Shows the HTTP headers that your client sends when connecting to a webserver.
curl "https://api.shodan.io/tools/httpheaders?key=your_api_key"

## My IP Address
# Get your current IP address as seen from the Internet.
curl "https://api.shodan.io/tools/myip?key=your_api_key"
