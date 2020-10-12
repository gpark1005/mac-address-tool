# MAC Address Lookup Tool

MAC Address Lookup Tool is a dockerized CLI tool written in Go. It can be used to lookup vendor information using a given MAC address. 
The tool makes request to https://macaddress.io/ using the provided arguments.


## Obtaining an API key

Usage of this tool requires an API key. To obtain one, perform the following steps: 
1. Create an account at https://macaddress.io
2. Once you are logged in, click your login name in the upper right hand corner of the webpage.
3. Copy your API key 


## Installation

To install this tool, Docker must be installed on your local computer. Check [here](https://docs.docker.com/get-docker/) for steps on how to install docker.

Before completing the steps below make sure you have followed the instructions under **Obtaining an API key**


1. Clone the repo and `cd` into the cloned directory

2. Build the docker image: `docker build -t mac-address-tool .`
3. (Optional) To avoid having to constantly enter your API key for each request, it is convenient to set it as an environment variable: `export mlkey=<Your API key>`
4. (Optional) Store the following alias in your `~/bash_aliases` file to avoid having to run the full `docker run...` command for each request: `alias mlu='docker run --rm mac-address-tool -k=$mlkey'`


## Usage

The CLI tool accepts 3 arguments:
1. `-k=<Your API Key>` (Required) Your personal API key.
2. `-a=<MAC address>` (Required) MAC address or OUI. You can use any octet delimiters including ':', '.', or even no delimiter. At least 6 BASE16 chars should be provided. 
3. `-f=<Output format>` (Optional) Output format: There are 3 options for this argument
      - 'json' — Full MAC address information in JSON format.
      - 'xml' — Full MAC address information in XML format.
      - 'csv' — Full MAC address information in CSV format.
      - If argument is excluded, program will output vendor company name only, in text format.


If you performed steps 3 and 4 in the **Installation** section, you can query the MAC address database like so:
-  `mlu -a=<MAC address>` or `mlu -a=<MAC address> -f=json`
      
      - Example: `mlu -a=44:38:39:ff:ef:57` or `mlu -a=44:38:39:ff:ef:57 -f=json`

If you did not perform steps 3 and 4 you will need to enter the full docker command and pass your API key as an argument:
-  `docker run --rm mac-address-tool -k=<Your API key> -a=<MAC address> -f=<format>`

     - Example: `docker run --rm mac-address-tool -k=someAPIKey -a=44:38:39:ff:ef:57 -f=json`
     
Check [here](https://macaddress.io/api/documentation/making-requests) for further explanation of arguments, errors codes, and output.

> Security Note: There are inherent risk when choosing to store your API key in an environment variable, if an attacker gains access to your `bash_aliases` file, they will be able to make request using your API key.
> Passing the argument with each request instead of storing it will mitigate risk but will not eliminate it as the key can still be retrieved from shell logs



 
 ## License
 [MIT](https://choosealicense.com/licenses/mit/)