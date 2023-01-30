## Demo of EBPF monitoring

This demo tries to show a demo of EBPF monitoring. There is a simple golang server which calls another golang http server to get the 
message to print. The EBPF uses a Kprobe to monitor the http requests flowing through, and prints out the message to console for now.
This can easily be changed to send it to a database.


### Directory structure

- **app**: Contains the demo application we will be tracing.
- **helperapp**: Helper app which 'app' calls to get the message to print.
- **http\_trace\_kprobe**: HTTP tracer based on kprobes. This one traces the responses that a pid generates. Start the process like **./http_trace_kprobe --pid 4376**
- **http\_trace\_kprobe\_read**: HTTP tracer which traces requests.

## Building the code

Simply run `make` in each of the sub-directories to build those underlying binaries.
You will need to run these on a Linux machine with [bcc](https://github.com/iovisor/bcc/blob/master/INSTALL.md) installed.
I used a simple  virtualbox with Ubuntu 22.04,  for the purpose.
	
