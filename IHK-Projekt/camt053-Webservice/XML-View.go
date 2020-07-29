package main

const XMLView = `
	<html>
		<center>
		<head>
			<div class="header-container">
			<link rel="stylesheet" type="text/css" href="/static/css/XML-View.css">
			<center>
			<title>CAMT053 - Webservice</title>
			<h1>CAMT053 - Webservice</h1>
			</center>
			</div>
		</head>
		</center>
		<body>
			<h4>XML-View</h4>
			<div class="xmlview-container">
			<h4>GRPHDR:</h4>
			<pre>{{.grphdr}}</pre>
			<br>
			<h4>STMT:</h4>
			<pre>{{.stmt}}</pre>
			</div>
			<br>
			<center>
			<form class="text-center" action="/Camt053/XML-View" method="post">
				<button type="submit" name="Zurück" formaction="/">Zurück</button>
			</form>
			</center>
		</body>
	</html>
`
