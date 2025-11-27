## Web Convert

Convert a web page to PDF, JPEG, PNG

## Key Features

1. **Multiple Output Formats**:

   - PDF (original functionality)
   - JPEG (with quality control)
   - PNG (lossless format)

2. **Enhanced UI**:

   - Dropdown to select output format
   - Quality slider for JPEG
   - Custom dimension controls
   - Full page capture option

3. **Conversion Logic**:

   - Uses `page.CaptureScreenshot()` for image capture
   - Supports viewport sizing and full-page scrolling
   - Maintains PDF generation capabilities

4. **Improved Dimension Handling**:

   - Standard paper sizes in pixels
   - Custom dimension support
   - Input validation

5. **Proper Content Types**:
   - Sets correct Content-Type headers
   - Appropriate file extensions

## How to Use

1. **Set your API key**:

   ```bash
   export API_KEY="your-secure-key-here"
   ```

2. **Run the service**:

   ```bash
   go run main.go
   ```

3. **Make conversions**:

   - Access the web interface at `http://localhost:3001`
   - Select desired output format (PDF, JPEG, or PNG)
   - Set quality (for JPEG) and dimensions
   - Submit to get your converted file

4. **API Endpoints**:

   ```bash
   # PDF conversion
   curl -X POST -H "X-API-Key: your-key" -F "url=https://example.com" -F "format=pdf" http://localhost:3001/convert -o output.pdf

   # JPEG conversion
   curl -X POST -H "X-API-Key: your-key" -F "url=https://example.com" -F "format=jpeg" -F "quality=85" http://localhost:3001/convert -o output.jpg

   # PNG conversion (full page)
   curl -X POST -H "X-API-Key: your-key" -F "url=https://example.com" -F "format=png" -F "full_page=on" http://localhost:3001/convert -o output.png
   ```

## Technical Notes

1. **Image Quality**:

   - JPEG: 1-100 quality scale (default 90)
   - PNG: Always lossless (quality parameter ignored)

2. **Viewport Sizing**:

   - Uses Chrome's device emulation to set viewport size
   - Full page capture scrolls and stitches the complete page

3. **Performance**:
   - Larger viewports and full page captures take more time
   - Timeout set to 60 seconds to accommodate complex pages
