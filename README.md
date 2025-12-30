# go-syscalls
A (potentially) operating system independent system call interface for Go

This package aims to solve a single problem. 
The golang.org/x/sys/unix package provides an interface to the UNIX system calls
but the contents is different on different systems.
Functions may or not be present in the package for a particular operating system.
In particular, the functions are not available at all under Windows,
which means that a program that uses them doesn't even compile in that environment.
I'm happy to use a Windows system to write and compile software that I then run under Linux,
so this is a problem for me.

My solution 
is to have a single package with a 
version for each system that I work with.
It provides the same functions with the same signatures on all of those systems.
In the windows version, functions that can't be made to work
return a "not implemented" error when called.
There is also a function OSName which runs in all environments and
returns the name of the operating system on which it is running
("windows", "linux" and so on, matching the build tag for that system.)

An example of this in use is my go-stripe-payments website.  
Running an HTTPS server requires a certificate,
which is impractical for a local Windows PC,
An HTTP server can run anywhere
and you can do a lot of system testing using one,
so I can do a lot of system testing before
I deploy it on my Linux target.
That requires the use of the Linux system calls,
but my software still compiles uner Windows.
I just need to avoid using some of the features when I run it there.

I describe this solution as only potentially an operating system independent solution because
(a) I don't plan to implement it on all systems, just the ones I use (initially Windows and Linux)
and (b) I only plan to implement calls that I need for my own work.

If you want something more,
feel free to fork the project or use it as a guide to write your own.
Please don't create issues asking me to add functions.  The answer will be no, please create your own project.
Please don't send pull requests asking me to accept new functions that you've written.
