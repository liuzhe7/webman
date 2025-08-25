# Web Postman with Golang CORS Proxy

A lightweight web-based alternative to Postman built with a single HTML file (no frontend frameworks) and a Golang proxy service to solve CORS (Cross-Origin Resource Sharing) issues.

## Features

- **No Frontend Frameworks**: Single HTML file with vanilla JavaScript, styled with Tailwind CSS
- **Full HTTP Method Support**: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS
- **Request Customization**:
  - Custom headers
  - Query parameters
  - Multiple request body formats (JSON, Form Data, x-www-form-urlencoded, Text)
  - Authentication (Basic Auth, Bearer Token)
- **Response Visualization**:
  - Status code with color coding (success/redirect/error)
  - Response time and size metrics
  - Multiple view modes (Pretty, Raw, Preview)
  - Response headers and cookies display
- **CORS Bypass**: Golang proxy server handles cross-origin requests
- **Responsive Design**: Works on desktop and mobile devices
- **Modern UI**: Clean interface with animations and transitions

## Tech Stack

| Component    | Technology                                            |
| ------------ | ----------------------------------------------------- |
| Frontend     | HTML5, Vanilla JavaScript, Tailwind CSS, Font Awesome |
| Proxy Server | Golang (Standard Library only)                        |
| Networking   | HTTP/HTTPS, CORS Handling                             |

## Installation & Setup

### Prerequisites

- [Golang 1.16+](https://go.dev/dl/) (for running the proxy server)
- A modern web browser (Chrome, Firefox, Edge, Safari)

### Step 1: Clone or Download the Project

Get the project files (index.html and proxy.go) into a local directory:

```bash
# Create a project folder
mkdir web-postman && cd web-postman

# Download the files (or clone if using a repository)
# Example for manual download:
# - Save index.html from the project
# - Save proxy.go from the project
```

### Step 2: Run the Golang Proxy Server

1. Navigate to the project directory in your terminal
2. Start the proxy server:

```bash
go run proxy.go
```

You should see the following output:

```
2024/05/XX XX:XX:XX Proxy server starting on :8080
```

The proxy server will run on `http://localhost:8080` and is required to bypass CORS restrictions.

### Step 3: Launch the Web Postman Interface

1. Simply open the `index.html` file in your web browser (double-click the file or drag it into a browser window)
2. The interface will load immediately – no additional setup required

## Usage Guide

### Making a Request

1. **Select HTTP Method**: Click on one of the method buttons (GET, POST, PUT, etc.) – the active method will be highlighted
2. **Enter Target URL**: Type or paste the API endpoint URL into the URL input field (e.g., `https://jsonplaceholder.typicode.com/posts/1`)
3. **Customize Request (Optional)**:
   - **Headers**: Switch to the "Headers" tab to add/remove custom HTTP headers
   - **Body**: For methods that support it (POST, PUT, PATCH), switch to the "Body" tab and select a format (JSON is default)
   - **Params**: Switch to the "Params" tab to add query parameters (automatically appended to the URL)
   - **Auth**: Switch to the "Auth" tab to add Basic Auth or Bearer Token authentication
4. **Send Request**: Click the "Send Request" button (blue button with paper plane icon)

### Viewing the Response

After sending a request, the response will be displayed in the "Response" section:

1. **Status Information**:
   - Status code (color-coded: green for 2xx, yellow for 3xx, orange for 4xx, red for 5xx)
   - Response time (in milliseconds)
   - Response size (in bytes)
2. **Response Tabs**:
   - **Body**: View the response content in Pretty (formatted JSON), Raw (plain text), or Preview (simple HTML rendering) mode
   - **Headers**: View all response headers in a table format
   - **Cookies**: View any cookies returned by the server
3. **Copy Response**: Click the "Copy" button to copy the raw response to your clipboard

## Proxy Server Details

### How It Works

The Golang proxy server acts as an intermediary between your browser and the target API:

1. The web interface sends all requests to `http://localhost:8080/proxy` (your local proxy server)
2. The proxy server extracts the target URL from the `url` query parameter
3. The proxy forwards the request to the target API with the original method, headers, and body
4. The proxy receives the response from the target API, adds CORS headers to it, and sends it back to the web interface

### CORS Headers Added

The proxy server automatically adds the following headers to all responses to bypass browser CORS restrictions:

```http
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS, PATCH, HEAD
Access-Control-Allow-Headers: *
Access-Control-Expose-Headers: *
```

### Security Notes

- **Local Use Only**: This proxy is designed for local development/testing use. Do not expose it to the public internet without adding authentication and rate limiting.
- **Header Filtering**: The proxy skips forwarding sensitive headers like `Host`, `Origin`, and `Referer` to the target API to prevent conflicts.
- **URL Validation**: The proxy includes basic URL validation but always ensure you're sending requests to trusted APIs.

## Troubleshooting

### Proxy Server Not Running

If you see an error like "Failed to fetch" or "Connection refused":

1. Ensure the proxy server is running (`go run proxy.go`)
2. Verify the proxy is listening on port 8080 (check terminal output)
3. Make sure no other application is using port 8080 (change the port in `proxy.go` if needed)

### Invalid JSON Error

If you get an "Invalid JSON" error when sending a request:

1. Switch to the "Body" tab
2. Click the "Format JSON" button to fix formatting issues
3. Ensure your JSON syntax is valid (check for missing commas, quotes, or brackets)

### CORS Errors Persisting

If you still encounter CORS errors:

1. Confirm the request is being sent through the proxy (check browser dev tools > Network tab)
2. Verify the proxy server is adding the correct CORS headers (check response headers in dev tools)
3. Restart the proxy server to ensure changes take effect

## Customization

### Changing the Proxy Port

To run the proxy on a different port (e.g., 9090), modify the `ListenAndServe` line in `proxy.go`:

```go
log.Fatal(http.ListenAndServe(":9090", nil))
```

Then update the proxy URL in `index.html` (search for `proxyUrl`):

```javascript
const proxyUrl = `http://localhost:9090/proxy?url=${encodeURIComponent(
  fullUrl
)}`;
```

### Adding New Features

- **New Body Formats**: Extend the `bodyTypeBtns` event listener in `index.html` to support additional formats (e.g., XML)
- **Dark Mode**: Complete the dark mode implementation in the `themeToggle` event listener
- **Request History**: Add localStorage integration to save past requests
- **Environment Variables**: Add support for environment-specific variables (e.g., base URLs)

## License

This project is open-source and available under the MIT License. Feel free to use, modify, and distribute it as needed.
