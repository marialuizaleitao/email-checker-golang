# Email Checker Tool in Go

This is a simple command-line tool written in Go for checking domain-specific email configuration details, including MX, SPF, and DMARC records.

## Features

- **Domain Input**: Accepts domain names interactively via standard input.
- **Checks**: For each domain entered, the tool checks:
  - MX (Mail Exchange) records.
  - SPF (Sender Policy Framework) records.
  - DMARC (Domain-based Message Authentication, Reporting & Conformance) records.
- **Output**: Displays the results in a structured format containing:
  - Domain name
  - Presence of MX records
  - Presence of SPF records and the SPF record itself (if found)
  - Presence of DMARC records and the DMARC record itself (if found)

## How to Use

1. **Installation**: Ensure you have Go installed on your system.

2. **Clone Repository**:

   ```bash
   git clone https://github.com/marialuizaleitao/email-checker-golang.git
   cd email-checker
   
3. **Run the Tool**

   ```bash
   go run main.go

4. **Enter Domain Names:**
   - Enter domain names one per line when prompted.
   - Press Ctrl + C to exit the tool.

5. **Output Example**

   Upon entering a domain name, the tool will display:

   ```bash
      Domain, HasMX, HasSPF, SPFRecord, HasDMARC, DMARCRecord
      example.com, true, true, v=spf1 include:_spf.example.com ~all, true, v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
   ```   
   - HasMX: Indicates if MX records are found for the domain.
   - HasSPF: Indicates if SPF records (v=spf1) are found and displays the SPF record.
   - HasDMARC: Indicates if DMARC records (v=DMARC1) are found and displays the DMARC record.
  
  ## Requirements

  Go 1.16 or later
