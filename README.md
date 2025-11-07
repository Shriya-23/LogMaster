🧩 LogMaster — Smart Log Organizer for Modern Developers

“One messy folder of logs sparked the idea for automation.”

💭 Why I Built This

When learning about system design, monitoring,Computer fundamentals and their working I realized that logs are the unsung heroes of every application.

They hold the story of your system,its errors, warnings, successes, and secrets.

But unless they’re organized, they’re just noise.

So I wanted to turn that noise into clarity,and that’s what LogMaster does.

“It’s not a big framework. It’s a small step towards cleaner debugging.”

🧩 What is LogMaster?

LogMaster is a CLI-based log file organizer written in Go.

It scans any folder you provide and automatically:

Categorizes logs (e.g., system_logs, security_audits, database_events, etc.)

Moves or copies them into structured subdirectories

Helps developers and DevOps engineers quickly analyze and locate critical logs

It’s a perfect utility to clean up messy environments, especially during testing, deployment, or monitoring phases.


🔹The Problem It Solves

In most projects:

Log files grow chaotically over time.

Debugging becomes a treasure hunt through mixed files.


LogMaster automates all that. It’s your personal assistant for logs,quietly keeping everything tidy while you focus on what actually matters: building and fixing things.

✨ Features

-Organizes logs into clean, structured folders

-Works locally or inside Docker containers

-Lightweight — built in Go with zero dependencies

-Supports large directories and nested folders

-Developer-friendly CLI experience

⚙️ Tech Stack

🔹Component	Technology

-Language	Go (Golang 1.22)

-Packaging	Docker

-Libraries	Standard Go libraries (os, path/filepath, strings)

-Environment	Alpine Linux (Docker Base)

🔹How to Run It Locally

🔹Option 1: Run with Go

# Clone the repository

git clone https://github.com/Shriya-23/LogMaster.git

cd LogMaster

# Build and run

go run main.go


You’ll be prompted to enter your log folder path:

Enter the path of log folder to organize: C:\Users\YourName\Documents\LogMasterTest

🔹Example Log Categories that i built locally to test them [ Not pushed to Github]

-System Logs — system_activity.log, kernel_monitor.log, os_events.log
-Security Audits — security_audit.log, firewall_events.log, intrusion_attempts.log

⚙️ Working

LogMaster automates the process of sorting and organizing log files intelligently.

-Input Folder Selection:
The user specifies a directory path containing raw log files.

-Timestamp Extraction:
LogMaster scans each log file, identifies timestamps (e.g., 2025-11-01 10:30:00), and uses them to determine when each event occurred.

 -Smart Categorization:
Based on keywords and patterns (like ERROR, INFO, WARN, DEBUG, or filenames such as security, system, database), the tool classifies each log into its respective category.

🏗️ Organized Folder Structure:

It creates structured folders such as:

/OrganizedLogs/

    /2025-11-01/
    
        /Security/
        
        /System/
        
        /Database/


Each folder contains logs sorted by date and category, making it effortless to find what you need.

-Readable Output:
Once organized, the program prints a summary of the operation — total logs processed, categories created, and storage path.

💼 About Me

Hi! I’m Shriya Sharma, A Computer Science student passionate about building practical, data-driven, and impactful tech solutions.

I love transforming ideas into simple, meaningful tools that bridge the gap between technology and real-world problems.

Have any ideas, suggestions, or feedback about this project? Feel free to reach out at shriya.sharma1923@gmail.com
I’d love to hear from you!

