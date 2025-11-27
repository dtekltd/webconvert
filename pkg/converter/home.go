package converter

import (
	"github.com/gofiber/fiber/v2"
)

var version = "v1.0.1"

func homeHandler(c *fiber.Ctx) error {
	return c.Type("html").SendString(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Web Page Converter</title>
		<style>
			body { font-family: Arial, sans-serif; max-width: 800px; margin: 0 auto; padding: 20px; }
			form { display: flex; flex-direction: column; gap: 15px; }
			input, select { padding: 8px; font-size: 16px; }
			button { padding: 10px 15px; background: #007bff; color: white; border: none; cursor: pointer; }
			button:hover { background: #0056b3; }
			.field-group { display: flex; gap: 15px; }
			.field-group > div { height: 100%; }
			.copy { padding: 20px; text-align: center; font-size: 11px; color: #999; }
			.copy a { color: #999; text-decoration: none; }
		</style>
	</head>
	<body>
		<h1>Web Page Converter <small>(` + version + `)</small></h1>
		<form action="/convert" method="post">
			<label for="url">URL to convert:</label>
			<input type="url" id="url" name="url" required placeholder="https://example.com">
			
			<label for="api_key">API Key:</label>
			<input type="password" id="api_key" name="apiKey" required>
			
			<div class="field-group">
				<div>
					<label for="format">Output Format:</label>
					<select id="format" name="format" required>
						<option value="pdf">PDF</option>
						<option value="jpeg">JPEG</option>
						<option value="png">PNG</option>
					</select>
				</div>
				<div id="quality-field" style="display: none;">
					<label for="quality">Quality (1-100):</label>
					<input type="number" id="quality" name="quality" min="1" max="100" value="90">
				</div>
			</div>
			
			<div class="field-group">
				<div>
					<label for="size">Page Size:</label>
					<select id="size" name="size">
						<option value="a4">A4 (210 × 297 mm)</option>
						<option value="letter" selected>Letter (8.5 × 11 in)</option>
						<option value="legal">Legal (8.5 × 14 in)</option>
						<option value="a3">A3 (297 × 420 mm)</option>
						<option value="custom">Custom</option>
					</select>
				</div>
				<div id="custom-size-fields" style="display: none;">
					<label for="width">Width (px):</label>
					<input type="number" id="width" name="width" value="1024" style="width: 80px">
					
					<label for="height">Height (px):</label>
					<input type="number" id="height" name="height" value="768" style="width: 80px">
				</div>
			</div>

			<label for="size">Page Margin (inches)</label>
			<div class="field-group">
				<div>
					<label for="width">Top:</label>
					<input id="margin-top" name="marginTop" value="0.4" style="width: 80px">
					<label for="width">Bottom:</label>
					<input id="margin-bottom" name="marginBottom" value="0.4" style="width: 80px">
					<label for="width">Left:</label>
					<input id="margin-left" name="marginLeft" value="0.4" style="width: 80px">
					<label for="width">Right:</label>
					<input id="margin-right" name="marginRight" value="0.4" style="width: 80px">
				</div>
			</div>
			
			<label>
				<input type="checkbox" id="landscape" name="landscape">
				Landscape orientation (pdf)
			</label>
			
			<label>
				<input type="checkbox" id="full_page" name="fullpage">
				Capture full page (jpeg, png - scroll and stitch)
			</label>
			
			<label>
				<input type="checkbox" id="background" name="background" checked>
				Include background graphics
			</label>
			
			<button type="submit">Convert</button>

			<div class="copy">&copy; 2025 - <a href="mailto:peter.phan07@gmail.com">Peter Phan</a></div>
		</form>
		
		<script>
			document.getElementById('format').addEventListener('change', function() {
				document.getElementById('quality-field').style.display = 
					(this.value === 'jpeg' || this.value === 'png') ? 'block' : 'none';
			});
			
			document.getElementById('size').addEventListener('change', function() {
				document.getElementById('custom-size-fields').style.display = 
					this.value === 'custom' ? 'block' : 'none';
			});
		</script>
	</body>
	</html>
	`)
}
