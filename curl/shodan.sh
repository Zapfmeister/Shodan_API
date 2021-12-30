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
